package userslogic

import (
	"context"

	"rpc/internal/svc"
	"rpc/station"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUsersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUsersLogic {
	return &AddUsersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddUsersLogic) AddUsers(in *station.UsersInfo) (*station.UsersInfo, error) {
	// todo: add your logic here and delete this line

	return &station.UsersInfo{}, nil
}
