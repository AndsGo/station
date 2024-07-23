package deliveryloglogic

import (
	"context"

	"rpc/internal/svc"
	"rpc/station"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDeliveryLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateDeliveryLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDeliveryLogLogic {
	return &UpdateDeliveryLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateDeliveryLogLogic) UpdateDeliveryLog(in *station.DeliveryLogInfo) (*station.DeliveryLogInfo, error) {

	return &station.DeliveryLogInfo{}, nil
}
