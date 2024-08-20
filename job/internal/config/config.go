package config

import (
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	StationRpcConf zrpc.RpcClientConf
	XxlJob         struct {
		Addresses   string
		AccessToken string
		Appname     string
		Ip          string
		Port        int
	}
}
