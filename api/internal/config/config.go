package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	RpcConf zrpc.RpcClientConf
	Secret  struct {
		AESSecret string
	}
}
