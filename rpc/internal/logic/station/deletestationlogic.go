package stationlogic

import (
	"context"

	"rpc/internal/svc"
	"rpc/station"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteStationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteStationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteStationLogic {
	return &DeleteStationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteStationLogic) DeleteStation(in *station.IDPathReq) (*station.BaseDataInfo, error) {
	// todo: add your logic here and delete this line

	return &station.BaseDataInfo{}, nil
}
