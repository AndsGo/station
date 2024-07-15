package stationlogic

import (
	"context"

	"rpc/internal/svc"
	"rpc/station"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStationLogic {
	return &GetStationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetStationLogic) GetStation(in *station.IDPathReq) (*station.StationInfo, error) {
	// todo: add your logic here and delete this line

	return &station.StationInfo{}, nil
}
