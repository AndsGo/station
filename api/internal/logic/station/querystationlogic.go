package station

import (
	"context"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryStationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询站点
func NewQueryStationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryStationLogic {
	return &QueryStationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryStationLogic) QueryStation(req *types.StationReq) (resp *types.StationListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
