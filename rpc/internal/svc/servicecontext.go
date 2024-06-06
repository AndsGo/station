package svc

import (
	"rpc/internal/config"
	"rpc/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	// 数据 redis
	StationUsersModel model.StationUsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	cnn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:            c,
		StationUsersModel: model.NewStationUsersModel(cnn),
	}
}
