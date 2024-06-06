package base

import (
	"context"
	"rpc/example/proto/greet"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HelloLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHelloLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HelloLogic {
	return &HelloLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HelloLogic) Hello() (resp *types.BaseResponse, err error) {
	// todo: add your logic here and delete this line
	// 业务代码
	res, err := l.svcCtx.GreetRpc.SayHello(l.ctx, &greet.SayHelloReq{})
	if err != nil {
		return nil, err
	}
	return &types.BaseResponse{
		Code:    0,
		Data:    "hello",
		Message: res.Message,
	}, nil
}
