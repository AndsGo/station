package svc

import (
	"api/internal/config"
	"rpc/client/greet"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	//rpc  类似bean管理
	GreetRpc greet.Greet
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		GreetRpc: greet.NewGreet(zrpc.MustNewClient(c.RpcConf)),
	}
}
