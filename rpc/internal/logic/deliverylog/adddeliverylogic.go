package deliveryloglogic

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"rpc/internal/svc"
	"rpc/model"
	"rpc/station"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddDeliveryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddDeliveryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddDeliveryLogic {
	return &AddDeliveryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 批量新增投递日志
func (l *AddDeliveryLogic) AddDelivery(in *station.DeliveryLogListInfo) (*station.BaseDataInfo, error) {
	//构建日志队列
	logs := make([]*model.DeliveryLog, 0)
	for _, s := range in.Data {
		// 查询是否有重复
		li, _ := l.svcCtx.DeliveryLogModel.FindOneByPostsIdStationId(context.Background(), sql.NullInt64{Int64: int64(s.Id), Valid: true}, sql.NullInt64{Int64: int64(s.Id), Valid: true})
		if li != nil && li.Id > 0 {
			continue
		}
		log := &model.DeliveryLog{
			Title:        sql.NullString{String: s.Title, Valid: true},
			Author:       sql.NullInt64{Int64: int64(s.Author), Valid: true},
			Source:       sql.NullString{String: s.Source, Valid: true},
			DomainName:   sql.NullString{String: s.DomainName, Valid: true},
			DeliveryDate: sql.NullTime{Time: time.Now(), Valid: true},
			Deliverer:    sql.NullString{String: s.Deliverer, Valid: true},
			Status:       sql.NullInt64{Int64: 0, Valid: true},
			PostsId:      sql.NullInt64{Int64: int64(s.PostsId), Valid: true},
			StationId:    sql.NullInt64{Int64: int64(s.StationId), Valid: true},
			WpCateIds:    sql.NullString{String: s.WpCateIds, Valid: true},
			Result:       sql.NullString{String: "", Valid: true},
		}
		logs = append(logs, log)
	}
	//批量入库
	if len(logs) > 0 {
		err := l.svcCtx.DeliveryLogModel.InsertBulk(context.Background(), logs, l.bulkResut)
		if err != nil {
			return nil, err
		}
	}
	return &station.BaseDataInfo{Data: fmt.Sprint(len(logs))}, nil
}

// 批量执行结果
func (l *AddDeliveryLogic) bulkResut(res sql.Result, err error) {
	if err != nil {
		l.Logger.Error("bulkResut fail", err.Error())
	}
	i, err := res.RowsAffected()
	if err != nil {
		l.Logger.Error("bulkResut RowsAffected fail", err.Error())
	}
	l.Logger.Info("bulkResut:", i)
}
