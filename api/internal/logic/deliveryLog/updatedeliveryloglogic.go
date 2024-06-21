package deliveryLog

import (
	"context"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDeliveryLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 修改DeliveryLog
func NewUpdateDeliveryLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDeliveryLogLogic {
	return &UpdateDeliveryLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateDeliveryLogLogic) UpdateDeliveryLog(req *types.DeliveryLogInfo) (resp *types.DeliveryLogInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
