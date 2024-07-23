package posts

import (
	"context"
	"rpc/station"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询Posts详情
func NewGetPostsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostsLogic {
	return &GetPostsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPostsLogic) GetPosts(req *types.IDPathReq) (resp *types.PostsInfoResp, err error) {
	info, err := l.svcCtx.PostsRpc.GetPosts(l.ctx, &station.IDPathReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return &types.PostsInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Message: types.Success,
		},
		Data: types.PostsInfo{
			Id:         info.Id,
			Title:      info.Title,
			Author:     info.Author,
			Source:     info.Source,
			Categories: info.Categories,
			Content:    info.Content,
			CreateTime: info.CreateTime,
			ThrownNum:  info.ThrownNum,
		},
	}, nil
}
