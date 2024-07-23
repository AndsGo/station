package deliveryLog

import (
	"context"
	"rpc/station"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDeliveryLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 修改DeliveryLog
func NewUpdateDeliveryLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDeliveryLogLogic {
	return &UpdateDeliveryLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateDeliveryLogLogic) UpdateDeliveryLog(req *types.DeliveryLogInfo) (resp *types.DeliveryLogInfoResp, err error) {
	info := &station.DeliveryLogInfo{
		Id:           req.Id,
		Author:       req.Author,
		Deliverer:    req.Deliverer,
		DeliveryDate: req.DeliveryDate,
		DomainName:   req.DomainName,
		PostsId:      req.PostsId,
		Source:       req.Source,
		Status:       req.Status,
		Title:        req.Title,
		WpCateIds:    req.WpCateIds,
	}
	data, err := l.svcCtx.DeliverLogRpc.UpdateDeliveryLog(l.ctx, info)
	if err != nil {
		return nil, err
	}
	res := &types.DeliveryLogInfo{
		Id:           data.Id,
		Author:       data.Author,
		Deliverer:    data.Deliverer,
		DeliveryDate: data.DeliveryDate,
		DomainName:   data.DomainName,
		PostsId:      data.PostsId,
		Source:       data.Source,
		Status:       data.Status,
		Title:        data.Title,
		WpCateIds:    data.WpCateIds,
	}
	return &types.DeliveryLogInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Message: types.Success,
		},
		Data: *res,
	}, nil
}
