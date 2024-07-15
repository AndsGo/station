package deliveryloglogic

import (
	"context"

	"rpc/internal/svc"
	"rpc/station"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddDeliveryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddDeliveryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddDeliveryLogic {
	return &AddDeliveryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddDeliveryLogic) AddDelivery(in *station.DeliveryLogListInfo) (*station.BaseDataInfo, error) {
	// todo: add your logic here and delete this line
	// l.svcCtx.DeliveryLogModel.Insert(l.ctx, in)
	return &station.BaseDataInfo{}, nil
}
