package stationpostsrelationlogic

import (
	"context"

	"rpc/internal/svc"
	"rpc/station"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddStationPostsRelationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddStationPostsRelationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddStationPostsRelationLogic {
	return &AddStationPostsRelationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddStationPostsRelationLogic) AddStationPostsRelation(in *station.StationPostsRelationInfo) (*station.StationPostsRelationInfo, error) {
	// todo: add your logic here and delete this line

	return &station.StationPostsRelationInfo{}, nil
}
