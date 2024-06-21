package deliveryLog

import (
	"context"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddDeliveryLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 投放
func NewAddDeliveryLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddDeliveryLogLogic {
	return &AddDeliveryLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddDeliveryLogLogic) AddDeliveryLog(req *types.DeliveryListInfo) (resp *types.BaseDataInfo, err error) {
	// todo: add your logic here and delete this line

	return
}
