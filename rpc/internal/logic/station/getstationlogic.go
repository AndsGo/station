package stationlogic

import (
	"context"

	"rpc/internal/svc"
	"rpc/station"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStationLogic {
	return &GetStationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetStation 获取站点信息
func (l *GetStationLogic) GetStation(in *station.IDPathReq) (*station.StationInfo, error) {
	data, err := l.svcCtx.StationModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &station.StationInfo{
		Id:           data.Id,
		DomainName:   data.DomainName.String,
		DomainYear:   data.DomainYear.Int64,
		GoogleWeight: data.GoogleWeight.Float64,
		Type:         data.Type.String,
		Industry:     data.Industry.String,
		UserName:     data.UserName.String,
		PassWord:     data.PassWord.String,
		ArticlesNum:  data.ArticlesNum.Int64,
		Ip:           data.Ip.String,
	}, nil
}
