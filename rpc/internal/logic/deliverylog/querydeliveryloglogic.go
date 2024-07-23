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

// 分页查询
func (l *QueryDeliveryLogLogic) QueryDeliveryLog(in *station.DeliveryLogReq) (*station.DeliveryLogListInfo, error) {
	// 查询数据库
	list, total, err := l.svcCtx.DeliveryLogModel.QueryDeliveryLog(l.ctx, in)
	if err != nil {
		return nil, err
	}
	// 组装数据
	resp := &station.DeliveryLogListInfo{
		BaseListInfo: &station.BaseListInfo{
			Total: total,
		},
	}
	for _, v := range list {
		resp.Data = append(resp.Data,
			&station.DeliveryLogInfo{
				Id:           v.Id,
				Title:        v.Title.String,
				Source:       v.Source.String,
				DomainName:   v.DomainName.String,
				DeliveryDate: v.DeliveryDate.Time.UnixMilli(),
				Deliverer:    v.Deliverer.String,
				Status:       v.Status.Int64,
				Author:       uint64(v.Author.Int64),
				PostsId:      uint64(v.PostsId.Int64),
				StationId:    uint64(v.StationId.Int64),
				WpCateIds:    v.WpCateIds.String,
			})
	}
	return resp, nil
}
