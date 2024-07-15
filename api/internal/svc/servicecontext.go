package svc

import (
	"api/internal/config"
	"rpc/client/deliverylog"
	"rpc/client/posts"
	"rpc/client/station"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	//rpc  类似bean管理
	StationRpc    station.Station
	PostsRpc      posts.Posts
	DeliverLogRpc deliverylog.DeliveryLog
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		StationRpc:    station.NewStation(zrpc.MustNewClient(c.RpcConf)),
		PostsRpc:      posts.NewPosts(zrpc.MustNewClient(c.RpcConf)),
		DeliverLogRpc: deliverylog.NewDeliveryLog(zrpc.MustNewClient(c.RpcConf)),
	}
}
