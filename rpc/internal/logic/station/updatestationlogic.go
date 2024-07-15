package stationlogic

import (
	"context"

	"rpc/internal/svc"
	"rpc/station"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateStationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStationLogic {
	return &UpdateStationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateStationLogic) UpdateStation(in *station.StationInfo) (*station.StationInfo, error) {
	// todo: add your logic here and delete this line

	return &station.StationInfo{}, nil
}
