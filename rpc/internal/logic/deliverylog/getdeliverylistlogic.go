package deliveryloglogic

import (
	"context"

	"rpc/internal/svc"
	"rpc/station"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDeliveryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDeliveryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDeliveryListLogic {
	return &GetDeliveryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDeliveryListLogic) GetDeliveryList(in *station.DeliveryInfo) (*station.DeliveryLogListInfo, error) {
	// todo: add your logic here and delete this line

	return &station.DeliveryLogListInfo{}, nil
}
