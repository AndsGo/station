package station

import (
	"context"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddStationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 新增站点
func NewAddStationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddStationLogic {
	return &AddStationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddStationLogic) AddStation(req *types.StationInfo) (resp *types.StationInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
