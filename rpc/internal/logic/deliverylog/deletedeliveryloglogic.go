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
	// todo: add your logic here and delete this line

	return &station.BaseDataInfo{}, nil
}
