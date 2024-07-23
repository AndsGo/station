package station

import (
	"context"
	"rpc/station"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryStationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询站点
func NewQueryStationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryStationLogic {
	return &QueryStationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryStationLogic) QueryStation(req *types.StationReq) (resp *types.StationListResp, err error) {
	// todo: add your logic here and delete this line
	info := &station.StationReq{
		DomainName:   req.DomainName,
		DomainYear:   req.DomainYear,
		GoogleWeight: req.GoogleWeight,
		Industry:     req.Industry,
		Ip:           req.Ip,
		Type:         req.Type,
	}
	info.PageInfo = &station.PageInfo{
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	data, err := l.svcCtx.StationRpc.QueryStation(l.ctx, info)
	if err != nil {
		return nil, err
	}
	resp = &types.StationListResp{
		BaseDataInfo: types.BaseDataInfo{
			Message: "success",
		},
	}
	resp.Data.Total = data.BaseListInfo.Total
	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data,
			types.StationInfo{
				Id:           v.Id,
				DomainName:   v.DomainName,
				Ip:           v.Ip,
				DomainYear:   v.DomainYear,
				GoogleWeight: v.GoogleWeight,
				Type:         v.Type,
				Industry:     v.Industry,
				ArticlesNum:  v.ArticlesNum,
				Status:       v.Status,
			})
	}
	return resp, nil
}
