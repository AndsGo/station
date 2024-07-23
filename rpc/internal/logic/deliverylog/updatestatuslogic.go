package deliveryloglogic

import (
	"context"

	"rpc/internal/svc"
	"rpc/station"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStatusLogic {
	return &UpdateStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新日志状态
func (l *UpdateStatusLogic) UpdateStatus(in *station.IDStatusReq) (*station.BaseDataInfo, error) {
	err := l.svcCtx.DeliveryLogModel.UpdateStatus(l.ctx, in.Id, int(in.Status), in.Result)
	if err != nil {
		return nil, err
	}
	return &station.BaseDataInfo{}, err
}
