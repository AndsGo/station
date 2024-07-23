package deliveryloglogic

import (
	"context"
	"database/sql"
	"time"

	"rpc/internal/svc"
	"rpc/model"
	"rpc/station"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddDeliveryLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddDeliveryLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddDeliveryLogLogic {
	return &AddDeliveryLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 新增投递日志
func (l *AddDeliveryLogLogic) AddDeliveryLog(in *station.DeliveryLogInfo) (*station.DeliveryLogInfo, error) {
	log := &model.DeliveryLog{
		Title:        sql.NullString{String: in.Title, Valid: true},
		Author:       sql.NullInt64{Int64: int64(in.Author), Valid: true},
		Source:       sql.NullString{String: in.Source, Valid: true},
		DomainName:   sql.NullString{String: in.DomainName, Valid: true},
		DeliveryDate: sql.NullTime{Time: time.Now(), Valid: true},
		Deliverer:    sql.NullString{String: in.Deliverer, Valid: true},
		Status:       sql.NullInt64{Int64: 0, Valid: true},
		PostsId:      sql.NullInt64{Int64: int64(0), Valid: true},
		StationId:    sql.NullInt64{Int64: int64(0), Valid: true},
	}
	res, err := l.svcCtx.DeliveryLogModel.Insert(l.ctx, log)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &station.DeliveryLogInfo{
		Id:           uint64(id),
		Title:        in.Title,
		Author:       in.Author,
		Source:       in.Source,
		DomainName:   in.DomainName,
		DeliveryDate: in.DeliveryDate,
		Deliverer:    in.Deliverer,
		Status:       in.Status,
		PostsId:      in.PostsId,
		StationId:    in.PostsId,
	}, nil
}
