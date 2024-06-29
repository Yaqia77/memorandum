package main

import (
	"net"

	"github.com/yaqia77/memorandum/apps/user/internal/handler"
	"github.com/yaqia77/memorandum/apps/user/internal/repository/db/dao"
	"github.com/yaqia77/memorandum/apps/user/internal/service"
	"github.com/yaqia77/memorandum/config"
	"github.com/yaqia77/memorandum/pkg/discovery"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {
	config.InitConfig()
	dao.InitDB()

	//etcd 地址
	etcdAddress := []string{viper.GetString("etcd.address")}
	// 服务的注册
	etcdRegister := discovery.NewRegister(etcdAddress, logrus.New())
	grpcAddress := viper.GetString("server.grpcAddress")
	userNode := discovery.Server{
		Name: viper.GetString("server.domain"),
		Addr: grpcAddress,
	}
	server := grpc.NewServer()
	defer server.Stop()

	//绑定服务
	service.RegisterUserServiceServer(server, handler.NewUserService())
	lis, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		panic(err)
	}
	if _, err := etcdRegister.Register(userNode, 10); err != nil {
		panic(err)
	}
	if err = server.Serve(lis); err != nil {
		panic(err)
	}
}
