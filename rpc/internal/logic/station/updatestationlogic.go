package stationlogic

import (
	"context"
	"database/sql"

	"rpc/internal/svc"
	"rpc/model"
	"rpc/station"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateStationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStationLogic {
	return &UpdateStationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateStationLogic) UpdateStation(in *station.StationInfo) (*station.StationInfo, error) {
	info := &model.Station{
		Id:           in.Id,
		DomainName:   sql.NullString{String: in.DomainName, Valid: true},
		Ip:           sql.NullString{String: in.Ip, Valid: true},
		DomainYear:   sql.NullInt64{Int64: in.DomainYear, Valid: true},
		GoogleWeight: sql.NullFloat64{Float64: in.GoogleWeight, Valid: true},
		Type:         sql.NullString{String: in.Type, Valid: true},
		Industry:     sql.NullString{String: in.Industry, Valid: true},
		ArticlesNum:  sql.NullInt64{Int64: in.ArticlesNum, Valid: true},
		UserName:     sql.NullString{String: in.UserName, Valid: true},
		PassWord:     sql.NullString{String: in.PassWord, Valid: true},
	}
	err := l.svcCtx.StationModel.Update(l.ctx, info)
	if err != nil {
		return nil, err
	}
	result := &station.StationInfo{
		Id:           info.Id,
		DomainName:   info.DomainName.String,
		Ip:           info.Ip.String,
		DomainYear:   info.DomainYear.Int64,
		GoogleWeight: info.GoogleWeight.Float64,
		Type:         info.Type.String,
		Industry:     info.Industry.String,
		ArticlesNum:  info.ArticlesNum.Int64,
		UserName:     info.UserName.String,
	}
	return result, nil
}
