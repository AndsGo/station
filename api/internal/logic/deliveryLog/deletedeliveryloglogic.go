package deliveryLog

import (
	"context"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteDeliveryLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除DeliveryLog
func NewDeleteDeliveryLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDeliveryLogLogic {
	return &DeleteDeliveryLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteDeliveryLogLogic) DeleteDeliveryLog(req *types.IDPathReq) (resp *types.BaseDataInfo, err error) {
	// todo: add your logic here and delete this line

	return
}
