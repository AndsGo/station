package station

import (
	"context"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 修改站点
func NewUpdateStationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStationLogic {
	return &UpdateStationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateStationLogic) UpdateStation(req *types.StationInfo) (resp *types.StationInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
