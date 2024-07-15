package deliveryloglogic

import (
	"context"

	"rpc/internal/svc"
	"rpc/station"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddDeliveryLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddDeliveryLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddDeliveryLogLogic {
	return &AddDeliveryLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddDeliveryLogLogic) AddDeliveryLog(in *station.DeliveryLogInfo) (*station.DeliveryLogInfo, error) {
	// todo: add your logic here and delete this line

	return &station.DeliveryLogInfo{}, nil
}
