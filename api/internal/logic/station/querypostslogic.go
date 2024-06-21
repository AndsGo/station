package station

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

// 获取关联的文章
func NewQueryPostsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryPostsLogic {
	return &QueryPostsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryPostsLogic) QueryPosts(req *types.IDPathReq) (resp *types.StationPostsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
