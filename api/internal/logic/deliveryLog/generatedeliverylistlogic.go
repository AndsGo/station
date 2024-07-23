package deliveryLog

import (
	"context"
	"rpc/station"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateDeliveryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取投放列表
func NewGenerateDeliveryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateDeliveryListLogic {
	return &GenerateDeliveryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GenerateDeliveryListLogic) GenerateDeliveryList(req *types.DeliveryInfo) (resp *types.DeliveryLogListResp, err error) {
	// 1.构造station.DeliveryInfo
	var info station.DeliveryInfo
	for _, item := range req.PostsInfoList {
		p := &station.PostsInfo{
			Id:         item.Id,
			Title:      item.Title,
			Source:     item.Source,
			Categories: item.Categories,
			Author:     item.Author,
		}
		info.PostsInfoList = append(info.PostsInfoList, p)
	}
	for _, item := range req.StationInfoList {
		p := &station.StationInfo{
			Id:         item.Id,
			DomainName: item.DomainName,
		}
		info.StationInfoList = append(info.StationInfoList, p)
	}
	// 2.调用rpc
	res, err := l.svcCtx.DeliverLogRpc.GetDeliveryList(l.ctx, &info)
	if err != nil {
		return nil, err
	}
	// 3.构造返回值
	resp = &types.DeliveryLogListResp{
		BaseDataInfo: types.BaseDataInfo{
			Message: types.Success,
		},
	}
	for _, item := range res.Data {
		p := &types.DeliveryLogInfo{
			Id:           item.Id,
			DomainName:   item.DomainName,
			Title:        item.Title,
			Deliverer:    item.Deliverer,
			PostsId:      item.PostsId,
			Status:       item.Status,
			Source:       item.Source,
			WpCateIds:    item.WpCateIds,
			Author:       item.Author,
			DeliveryDate: item.DeliveryDate,
		}
		resp.Data.Data = append(resp.Data.Data, *p)
	}
	resp.Data.Total = uint64(len(res.Data))
	return resp, nil
}
