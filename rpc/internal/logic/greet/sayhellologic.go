package greetlogic

import (
	"context"

	"rpc/example/proto/greet"
	"rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SayHelloLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSayHelloLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SayHelloLogic {
	return &SayHelloLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 定义一个 SayHello 一元 rpc 方法，请求体和响应体必填。
func (l *SayHelloLogic) SayHello(in *greet.SayHelloReq) (*greet.SayHelloResp, error) {
	// todo: add your logic here and delete this line
	data, err := l.svcCtx.StationUsersModel.FindOne(l.ctx, 2)
	if err != nil {
		return nil, err
	}
	return &greet.SayHelloResp{
		Message: data.Slug.String,
	}, nil
}
