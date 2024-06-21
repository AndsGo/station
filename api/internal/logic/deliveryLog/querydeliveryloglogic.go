package deliveryLog

import (
	"context"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryDeliveryLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询DeliveryLog
func NewQueryDeliveryLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryDeliveryLogLogic {
	return &QueryDeliveryLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryDeliveryLogLogic) QueryDeliveryLog(req *types.DeliveryLogReq) (resp *types.DeliveryLogListResp, err error) {
	// todo: add your logic here and delete this line

	return &types.DeliveryLogListResp{
		BaseDataInfo: types.BaseDataInfo{
			Code:    0,
			Message: "hello",
			Data:    "world",
		},
	}, nil
}
