package svc

import (
	"job/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}

// 全局svc
var sc *ServiceContext

func NewSvc(c config.Config) {
	sc = NewServiceContext(c)
}
func GetSvc() *ServiceContext {
	return sc
}
