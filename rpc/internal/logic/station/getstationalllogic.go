package stationlogic

import (
	"context"

	"rpc/internal/svc"
	"rpc/station"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStationAllLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStationAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStationAllLogic {
	return &GetStationAllLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetStationAllLogic) GetStationAll(in *station.IDPathReq) (*station.StationListInfo, error) {
	// todo: add your logic here and delete this line

	return &station.StationListInfo{}, nil
}
