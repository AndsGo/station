package posts

import (
	"context"
	"rpc/station"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePostsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除Posts
func NewDeletePostsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePostsLogic {
	return &DeletePostsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletePostsLogic) DeletePosts(req *types.IDPathReq) (resp *types.BaseDataInfo, err error) {
	info, err := l.svcCtx.PostsRpc.DeletePosts(l.ctx, &station.IDPathReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &types.BaseDataInfo{
		Message: info.Message,
	}, nil
}
