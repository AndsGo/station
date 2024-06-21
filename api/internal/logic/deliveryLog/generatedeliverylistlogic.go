package deliveryLog

import (
	"context"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateDeliveryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取投放列表
func NewGenerateDeliveryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateDeliveryListLogic {
	return &GenerateDeliveryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GenerateDeliveryListLogic) GenerateDeliveryList(req *types.DeliveryInfo) (resp *types.DeliveryLogListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
