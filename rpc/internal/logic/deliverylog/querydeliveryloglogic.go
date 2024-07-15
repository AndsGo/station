package deliveryloglogic

import (
	"context"

	"rpc/internal/svc"
	"rpc/station"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryDeliveryLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryDeliveryLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryDeliveryLogLogic {
	return &QueryDeliveryLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryDeliveryLogLogic) QueryDeliveryLog(in *station.DeliveryLogReq) (*station.DeliveryLogListInfo, error) {
	// todo: add your logic here and delete this line

	return &station.DeliveryLogListInfo{}, nil
}
