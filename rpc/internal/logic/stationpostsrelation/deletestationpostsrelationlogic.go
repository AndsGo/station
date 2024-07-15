package stationpostsrelationlogic

import (
	"context"

	"rpc/internal/svc"
	"rpc/station"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteStationPostsRelationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteStationPostsRelationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteStationPostsRelationLogic {
	return &DeleteStationPostsRelationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteStationPostsRelationLogic) DeleteStationPostsRelation(in *station.IDPathReq) (*station.BaseDataInfo, error) {
	// todo: add your logic here and delete this line

	return &station.BaseDataInfo{}, nil
}
