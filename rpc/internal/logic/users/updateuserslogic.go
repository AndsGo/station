package userslogic

import (
	"context"

	"rpc/internal/svc"
	"rpc/station"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUsersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUsersLogic {
	return &UpdateUsersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUsersLogic) UpdateUsers(in *station.UsersInfo) (*station.UsersInfo, error) {
	// todo: add your logic here and delete this line

	return &station.UsersInfo{}, nil
}
