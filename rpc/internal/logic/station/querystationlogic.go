package stationlogic

import (
	"context"

	"rpc/internal/svc"
	"rpc/station"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryStationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryStationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryStationLogic {
	return &QueryStationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryStationLogic) QueryStation(in *station.StationReq) (*station.StationListInfo, error) {
	// 查询数据
	list, total, err := l.svcCtx.StationModel.QueryStation(l.ctx, in)
	if err != nil {
		return nil, err
	}
	resp := &station.StationListInfo{
		BaseListInfo: &station.BaseListInfo{
			Total: total,
		},
	}
	// 封装数据
	for _, v := range list {
		resp.Data = append(resp.Data,
			&station.StationInfo{
				Id:           v.Id,
				DomainName:   v.DomainName.String,
				Ip:           v.Ip.String,
				DomainYear:   v.DomainYear.Int64,
				GoogleWeight: v.GoogleWeight.Float64,
				Type:         v.Type.String,
				Industry:     v.Industry.String,
				ArticlesNum:  v.ArticlesNum.Int64,
				Status:       v.Status.Int64,
			})
	}
	return resp, nil
}
