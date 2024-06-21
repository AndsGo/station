package posts

import (
	"context"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePostsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 修改Posts
func NewUpdatePostsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePostsLogic {
	return &UpdatePostsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePostsLogic) UpdatePosts(req *types.PostsInfo) (resp *types.PostsInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
