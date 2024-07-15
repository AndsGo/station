package stationlogic

import (
	"context"

	"rpc/internal/svc"
	"rpc/station"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddStationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddStationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddStationLogic {
	return &AddStationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddStationLogic) AddStation(in *station.StationInfo) (*station.StationInfo, error) {
	// todo: add your logic here and delete this line

	return &station.StationInfo{}, nil
}
