package userslogic

import (
	"context"

	"rpc/internal/svc"
	"rpc/station"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryUsersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryUsersLogic {
	return &QueryUsersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryUsersLogic) QueryUsers(in *station.UsersReq) (*station.UsersListInfo, error) {
	// todo: add your logic here and delete this line

	return &station.UsersListInfo{}, nil
}
