package userslogic

import (
	"context"

	"rpc/internal/svc"
	"rpc/station"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUsersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUsersLogic {
	return &DeleteUsersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUsersLogic) DeleteUsers(in *station.IDPathReq) (*station.BaseDataInfo, error) {
	// todo: add your logic here and delete this line

	return &station.BaseDataInfo{}, nil
}
