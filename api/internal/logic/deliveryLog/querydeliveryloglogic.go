package deliveryLog

import (
	"context"
	"rpc/station"
	"time"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryDeliveryLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询DeliveryLog
func NewQueryDeliveryLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryDeliveryLogLogic {
	return &QueryDeliveryLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryDeliveryLogLogic) QueryDeliveryLog(req *types.DeliveryLogReq) (resp *types.DeliveryLogListResp, err error) {
	in := &station.DeliveryLogReq{
		PageInfo: &station.PageInfo{
			Page:     req.Page,
			PageSize: req.PageSize,
		},
		Author:     int64(req.Author),
		Deliverer:  req.Deliverer,
		DomainName: req.DomainName,
		Source:     req.Source,
		Status:     req.Status,
		Title:      req.Title,
	}
	if req.DeliveryDate > 0 {
		in.DeliveryDate = time.UnixMilli(req.DeliveryDate).Format("2006-01-02")
	}
	list, err := l.svcCtx.DeliverLogRpc.QueryDeliveryLog(l.ctx, in)
	if err != nil {
		return nil, err
	}
	resp = &types.DeliveryLogListResp{
		BaseDataInfo: types.BaseDataInfo{
			Message: types.Success,
		},
		Data: types.DeliveryLogListInfo{
			BaseListInfo: types.BaseListInfo{
				Total: list.BaseListInfo.Total,
			},
		},
	}
	for _, item := range list.Data {
		resp.Data.Data = append(resp.Data.Data, types.DeliveryLogInfo{
			Id:           item.Id,
			DomainName:   item.DomainName,
			Title:        item.Title,
			Deliverer:    item.Deliverer,
			DeliveryDate: item.DeliveryDate,
			Source:       item.Source,
			Status:       item.Status,
			PostsId:      item.PostsId,
			StationId:    item.StationId,
			Author:       item.Author,
			WpCateIds:    item.WpCateIds,
		})
	}
	return resp, nil
}
