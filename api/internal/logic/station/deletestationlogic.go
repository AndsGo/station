package station

import (
	"context"
	"rpc/station"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteStationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除站点
func NewDeleteStationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteStationLogic {
	return &DeleteStationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteStationLogic) DeleteStation(req *types.IDPathReq) (resp *types.BaseDataInfo, err error) {
	info, err := l.svcCtx.StationRpc.DeleteStation(l.ctx, &station.IDPathReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return &types.BaseDataInfo{
		Message: info.Message,
	}, nil
}
