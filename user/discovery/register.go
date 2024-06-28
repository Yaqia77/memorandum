package discovery

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"go.etcd.io/etcd/clientv3"
	"golang.org/x/net/context"
)

type Register struct {
	EtcdAddrs   []string
	DialTimeout int
	closeCh     chan struct{}
	leasesID    clientv3.LeaseID
	keepAliveCh <-chan *clientv3.LeaseKeepAliveResponse
	srvInfo     Server
	srvTTL      int64
	cli         *clientv3.Client
	logger      *logrus.Logger
	lease       *clientv3.Lease
}

// NewRegister 基于etcd创建一个register实例
func NewRegister(etcdAddrs []string, logger *logrus.Logger) *Register {
	return &Register{
		EtcdAddrs:   etcdAddrs,
		DialTimeout: 3,
		logger:      logger,
	}
}

func (r *Register) Register(srvInfo Server, ttl int64) (chan<- struct{}, error) {
	var err error
	if strings.Split(srvInfo.Addr, ":")[0] == "" {
		return nil, errors.New("invalid address")
	}
	//初始化
	if r.cli, err = clientv3.New(clientv3.Config{
		Endpoints:   r.EtcdAddrs,
		DialTimeout: time.Duration(r.DialTimeout) * time.Second,
	}); err != nil {
		return nil, err
	}

	r.srvInfo = srvInfo
	r.srvTTL = ttl
	if err = r.register(); err != nil {
		return nil, err
	}
	r.closeCh = make(chan struct{})
	go r.keepAlive()
	return r.closeCh, nil
}

func (r *Register) register() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(r.DialTimeout)*time.Second)
	defer cancel()
	// 获取租约
	leaseResp, err := r.cli.Create(ctx, r.srvTTL)

	if err != nil {
		return err
	}
	r.leasesID = clientv3.LeaseID(leaseResp.ID)
	if r.keepAliveCh, err = r.cli.KeepAlive(context.Background(), r.leasesID); err != nil {
		return err
	}
	data, err := json.Marshal(r.srvInfo)
	if err != nil {
		return err
	}
	_, err = r.cli.Put(context.Background(), BuildRegisterPath(r.srvInfo), string(data), clientv3.WithLease(r.leasesID))
	return err
}

func (r *Register) keepAlive() {
	ticker := time.NewTicker(time.Duration(r.srvTTL) * time.Second)
	for {
		select {
		case <-r.closeCh:
			if err := r.unregister(); err != nil {
				fmt.Println("unregister error:", err)
			}
			if _, err := r.cli.Revoke(context.Background(), r.leasesID); err != nil {
				fmt.Println("revoke error:", err)
			}
		case res := <-r.keepAliveCh:
			if res == nil {
				if err := r.register(); err != nil {
					fmt.Println("register error:", err)
				}
			}
		case <-ticker.C:
			if r.keepAliveCh == nil {
				if err := r.register(); err != nil {
					fmt.Println("register error:", err)
				}
			}
		}
	}
}

func (r *Register) unregister() error {
	_, err := r.cli.Delete(context.Background(), BuildRegisterPath(r.srvInfo))

	return err
}
