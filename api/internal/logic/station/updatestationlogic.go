package station

import (
	"context"
	"pkg/encryption"
	"pkg/iputil"
	"pkg/wordpressutil"
	"rpc/station"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 修改站点
func NewUpdateStationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStationLogic {
	return &UpdateStationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateStationLogic) UpdateStation(in *types.StationInfo) (resp *types.StationInfoResp, err error) {
	//校验wordpress 账号密码
	err = wordpressutil.Test(in.DomainName, in.UserName, in.PassWord)
	if err != nil {
		return nil, err
	}
	// 加密密码
	pw, err := encryption.AESEncrypt(in.PassWord, l.svcCtx.Config.Secret.AESSecret)
	if err != nil {
		return nil, err
	}

	info := &station.StationInfo{
		Id:           in.Id,
		DomainName:   in.DomainName,
		DomainYear:   in.DomainYear,
		GoogleWeight: in.GoogleWeight,
		Type:         in.Type,
		Industry:     in.Industry,
		UserName:     in.UserName,
		PassWord:     pw,
		ArticlesNum:  in.ArticlesNum,
		Ip:           in.Ip,
	}
	// 自动获取服务器ip
	if info.Ip == "" {
		ip, _ := iputil.GetDomainIP(info.DomainName)
		info.Ip = ip
	}
	// 1. 更新站点
	data, err := l.svcCtx.StationRpc.UpdateStation(l.ctx, info)
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
