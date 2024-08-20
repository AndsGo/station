package jobhandler

import (
	"context"
	"job/internal/logic/delivery"
	"job/internal/svc"

	xxl "github.com/xxl-job/xxl-job-executor-go"
	"github.com/zeromicro/go-zero/core/logx"
)

// 文章投放
func PostsDelivery(ctx context.Context, param *xxl.RunReq) (msg string) {
	logger := logx.WithContext(ctx)
	logger.Info("PostsDelivery start:" + param.ExecutorHandler + " param:" + param.ExecutorParams + " log_id:" + xxl.Int64ToStr(param.LogID))
	delivery.NewPostsDeliveryLogic(ctx, svc.GetSvc()).Execute(param)
	logger.Info("PostsDelivery end:" + param.ExecutorHandler + " param:" + param.ExecutorParams + " log_id:" + xxl.Int64ToStr(param.LogID))
	return "PostsDelivery done"
}
