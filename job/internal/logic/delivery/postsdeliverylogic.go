package delivery

import (
	"context"

	"job/internal/svc"

	"github.com/xxl-job/xxl-job-executor-go"
	"github.com/zeromicro/go-zero/core/logx"
)

type PostsDeliveryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostsDeliveryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostsDeliveryLogic {
	return &PostsDeliveryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostsDeliveryLogic) Execute(param *xxl.RunReq) error {
	// todo: add your logic here and delete this line
	// 1.获取待处理的队列数据
	// 2.循环处理待处理的数据
	// 2.1.获取文章内容
	// 2.2.组装发送内容
	// 2.3.发送到wordpress
	// 3.3.更新结果到队列表
	l.Logger.Info("PostsDeliveryLogic run end...")
	return nil
}
