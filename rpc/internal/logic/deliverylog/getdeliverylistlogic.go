package deliveryloglogic

import (
	"context"
	"database/sql"

	"rpc/internal/svc"
	"rpc/station"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDeliveryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDeliveryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDeliveryListLogic {
	return &GetDeliveryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取需要发布预览数据
func (l *GetDeliveryListLogic) GetDeliveryList(in *station.DeliveryInfo) (*station.DeliveryLogListInfo, error) {
	//构建日志队列
	logs := make([]*station.DeliveryLogInfo, 0)
	for _, s := range in.StationInfoList {
		for _, p := range in.PostsInfoList {
			li, _ := l.svcCtx.DeliveryLogModel.FindOneByPostsIdStationId(context.Background(), sql.NullInt64{Int64: int64(p.Id), Valid: true}, sql.NullInt64{Int64: int64(s.Id), Valid: true})
			if li != nil && li.Id > 0 {
				continue
			}
			log := &station.DeliveryLogInfo{
				Title:      p.Title,
				Author:     uint64(p.Author),
				Source:     p.Source,
				DomainName: s.DomainName,
				Deliverer:  in.Deliverer,
				PostsId:    p.Id,
				StationId:  s.Id,
				WpCateIds:  p.Categories,
			}
			logs = append(logs, log)
		}
	}
	return &station.DeliveryLogListInfo{
		Data: logs,
	}, nil
}
