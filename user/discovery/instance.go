package discovery

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"google.golang.org/grpc/resolver"
)

type Server struct {
	Name    string `json:"name"`    // 名称
	Addr    string `json:"addr"`    // 地址
	Version string `json:"version"` // 版本
	Weight  int    `json:"weight"`  // 权重
}

// BuildPrefix 根据服务信息生成注册路径前缀

func BuildPrefix(server Server) string {
	if server.Version == "" {
		return fmt.Sprintf("/%s/", server.Name)
	}
	return fmt.Sprintf("/%s/%s/", server.Name, server.Version)
}

// BuildRegisterPath 根据服务信息生成注册路径
func BuildRegisterPath(server Server) string {
	return fmt.Sprintf("%s%s", BuildPrefix(server), server.Addr)
}

// ParseValue 将value值反序列化到一个Server实例中
func ParseValue(value []byte) (Server, error) {
	server := Server{}
	if err := json.Unmarshal(value, &server); err != nil {
		return server, err
	}
	return server, nil
}

// SplitPath 根据注册路径分割出服务信息 (校验)
func SplitPath(path string) (Server, error) {
	server := Server{}
	strs := strings.Split(path, "/")
	if len(strs) == 0 {
		return server, errors.New("path is empty")
	}
	server.Addr = strs[len(strs)-1]
	return server, nil
}

func Exist(l []resolver.Address, addr resolver.Address) bool {
	for i := range l {
		if l[i].Addr == addr.Addr {
			return true
		}
	}
	return false
}
