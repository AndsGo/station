package deliveryloglogic

import (
	"context"

	"rpc/internal/svc"
	"rpc/station"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteDeliveryLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteDeliveryLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDeliveryLogLogic {
	return &DeleteDeliveryLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteDeliveryLogLogic) DeleteDeliveryLog(in *station.IDPathReq) (*station.BaseDataInfo, error) {
	err := l.svcCtx.DeliveryLogModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &station.BaseDataInfo{}, nil
}
