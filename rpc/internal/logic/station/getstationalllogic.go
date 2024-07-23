package stationlogic

import (
	"context"

	"rpc/internal/svc"
	"rpc/station"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStationAllLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStationAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStationAllLogic {
	return &GetStationAllLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取所有站点信息
func (l *GetStationAllLogic) GetStationAll(in *station.IDPathReq) (*station.StationListInfo, error) {
	list, err := l.svcCtx.StationModel.FindAll(l.ctx)
	if err != nil {
		return nil, err
	}
	resp := &station.StationListInfo{}
	for _, v := range list {
		resp.Data = append(resp.Data, &station.StationInfo{
			Id:           v.Id,
			DomainName:   v.DomainName.String,
			Ip:           v.Ip.String,
			DomainYear:   v.DomainYear.Int64,
			GoogleWeight: v.GoogleWeight.Float64,
			Type:         v.Type.String,
			Industry:     v.Industry.String,
			ArticlesNum:  v.ArticlesNum.Int64,
			PassWord:     v.PassWord.String,
			UserName:     v.UserName.String,
		})
	}
	return resp, nil
}
