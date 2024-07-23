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
	err := l.svcCtx.StationModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &station.BaseDataInfo{}, nil
}
