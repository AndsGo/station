package svc

import (
	"rpc/internal/config"
	"rpc/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	// 数据 redis
	DeliveryLogModel          model.DeliveryLogModel
	StationModel              model.StationModel
	PostsModel                model.PostsModel
	StationPostsTextModel     model.StationPostsTextModel
	StationPostsRelationModel model.StationPostsRelationModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	cnn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:                    c,
		DeliveryLogModel:          model.NewDeliveryLogModel(cnn),
		StationModel:              model.NewStationModel(cnn),
		PostsModel:                model.NewPostsModel(cnn),
		StationPostsTextModel:     model.NewStationPostsTextModel(cnn),
		StationPostsRelationModel: model.NewStationPostsRelationModel(cnn),
	}
}
