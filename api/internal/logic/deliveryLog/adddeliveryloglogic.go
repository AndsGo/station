package deliveryLog

import (
	"context"
	"rpc/station"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddDeliveryLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 投放
func NewAddDeliveryLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddDeliveryLogLogic {
	return &AddDeliveryLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddDeliveryLogLogic) AddDeliveryLog(req *types.DeliveryListInfo) (resp *types.BaseDataInfo, err error) {
	var info station.DeliveryLogListInfo
	//for循环,将 req.Data 转换为 info.Data
	for _, v := range req.Data {
		log := &station.DeliveryLogInfo{
			Author:       v.Author,
			Deliverer:    v.Deliverer,
			DeliveryDate: v.DeliveryDate,
			DomainName:   v.DomainName,
			PostsId:      v.PostsId,
			Source:       v.Source,
			StationId:    v.StationId,
			Status:       v.Status,
			Title:        v.Title,
			WpCateIds:    v.WpCateIds,
		}
		info.Data = append(info.Data, log)
	}
	res, err := l.svcCtx.DeliverLogRpc.AddDelivery(l.ctx, &info)
	if err != nil {
		return nil, err
	}
	return &types.BaseDataInfo{
		Message: res.Message,
		Data:    res.Data,
	}, nil
}
