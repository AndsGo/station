package postslogic

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

func (l *QueryStationLogic) QueryStation(in *station.IDPathReq) (*station.PostsStationResp, error) {
	// todo: add your logic here and delete this line

	return &station.PostsStationResp{}, nil
}
