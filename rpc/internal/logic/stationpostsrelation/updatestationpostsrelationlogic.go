package stationpostsrelationlogic

import (
	"context"

	"rpc/internal/svc"
	"rpc/station"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStationPostsRelationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateStationPostsRelationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStationPostsRelationLogic {
	return &UpdateStationPostsRelationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateStationPostsRelationLogic) UpdateStationPostsRelation(in *station.StationPostsRelationInfo) (*station.StationPostsRelationInfo, error) {
	// todo: add your logic here and delete this line

	return &station.StationPostsRelationInfo{}, nil
}
