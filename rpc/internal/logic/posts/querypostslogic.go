package postslogic

import (
	"context"

	"rpc/internal/svc"
	"rpc/station"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryPostsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryPostsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryPostsLogic {
	return &QueryPostsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分页查询
func (l *QueryPostsLogic) QueryPosts(in *station.PostsReq) (*station.PostsListInfo, error) {
	// 1.转换擦查询体条件
	// 2.查询数据看
	list, total, err := l.svcCtx.PostsModel.QueryPosts(l.ctx, in)
	if err != nil {
		return nil, err
	}
	// 3.转换返回数据
	resp := &station.PostsListInfo{
		BaseListInfo: &station.BaseListInfo{
			Total: total,
		},
	}
	if total > 0 {
		for _, item := range list {
			resp.Data = append(resp.Data, &station.PostsInfo{
				Id:         item.Id,
				Title:      item.Title.String,
				Source:     item.Source.String,
				Author:     item.Author.Int64,
				ThrownNum:  item.ThrownNum.Int64,
				CreateTime: item.CreateAt.UnixMilli(),
				Categories: item.Categories.String,
			})
		}
	}
	return resp, nil
}
