package deliveryLog

import (
	"context"
	"rpc/station"

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
	info, err := l.svcCtx.DeliverLogRpc.DeleteDeliveryLog(l.ctx, &station.IDPathReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return &types.BaseDataInfo{
		Message: info.Message,
	}, nil
}
