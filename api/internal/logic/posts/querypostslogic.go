package posts

import (
	"context"
	"rpc/client/station"
	"time"

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
	// 1.转换查询struct
	rpcRep := &station.PostsReq{
		PageInfo: &station.PageInfo{
			Page:     req.Page,
			PageSize: req.PageSize,
		},
		Title:      req.Title,
		Source:     req.Source,
		Author:     req.Author,
		Categories: req.Categories,
	}
	if req.CreateTime > 0 {
		rpcRep.CreateTime = time.UnixMilli(req.CreateTime).Format("2006-01-02")
	}
	// 2调用rpc接口查询
	postsListInfo, err := l.svcCtx.PostsRpc.QueryPosts(l.ctx, rpcRep)
	if err != nil {
		return nil, err
	}
	// 3.转换结果
	resp = &types.PostsListResp{
		BaseDataInfo: types.BaseDataInfo{
			Message: types.Success,
		},
		Data: types.PostsListInfo{
			BaseListInfo: types.BaseListInfo{
				Total: postsListInfo.BaseListInfo.Total,
			},
		},
	}
	if postsListInfo.BaseListInfo.Total > 0 {
		for _, item := range postsListInfo.Data {
			resp.Data.Data = append(resp.Data.Data, types.PostsInfo{
				Id:         item.Id,
				Title:      item.Title,
				Source:     item.Source,
				Author:     item.Author,
				ThrownNum:  item.ThrownNum,
				Categories: item.Categories,
				CreateTime: item.CreateTime,
				Content:    item.Content,
			})
		}
	}
	return resp, nil
}
