package station

import (
	"context"
	"rpc/station"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddStationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 新增站点
func NewAddStationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddStationLogic {
	return &AddStationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddStationLogic) AddStation(in *types.StationInfo) (resp *types.StationInfoResp, err error) {
	// todo 校验账号密码
	// todo 加密密码
	info := &station.StationInfo{
		DomainName:   in.DomainName,
		DomainYear:   in.DomainYear,
		GoogleWeight: in.GoogleWeight,
		Type:         in.Type,
		Industry:     in.Industry,
		UserName:     in.UserName,
		PassWord:     in.PassWord,
		ArticlesNum:  in.ArticlesNum,
		Ip:           in.Ip,
	}
	// todo 自动获取服务器ip

	// 1. 新增站点
	data, err := l.svcCtx.StationRpc.AddStation(l.ctx, info)
	if err != nil {
		return nil, err
	}
	// 2. 返回站点信息
	return &types.StationInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code:    0,
			Message: "success",
		},
		Data: types.StationInfo{
			Id:           data.Id,
			DomainName:   data.DomainName,
			DomainYear:   data.DomainYear,
			GoogleWeight: data.GoogleWeight,
			Type:         data.Type,
			Industry:     data.Industry,
			UserName:     data.UserName,
			PassWord:     data.PassWord,
			ArticlesNum:  data.ArticlesNum,
			Ip:           data.Ip,
			Status:       data.Status,
		},
	}, nil
}
