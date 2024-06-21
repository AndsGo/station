package posts

import (
	"context"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryPostsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询Posts
func NewQueryPostsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryPostsLogic {
	return &QueryPostsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryPostsLogic) QueryPosts(req *types.PostsReq) (resp *types.PostsListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
