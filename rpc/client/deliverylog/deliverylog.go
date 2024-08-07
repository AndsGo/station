// Code generated by goctl. DO NOT EDIT.
// Source: station.proto

package deliverylog

import (
	"context"

	"rpc/station"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	BaseDataInfo             = station.BaseDataInfo
	BaseListInfo             = station.BaseListInfo
	DeliveryInfo             = station.DeliveryInfo
	DeliveryLogInfo          = station.DeliveryLogInfo
	DeliveryLogListInfo      = station.DeliveryLogListInfo
	DeliveryLogReq           = station.DeliveryLogReq
	IDPathReq                = station.IDPathReq
	IDStatusReq              = station.IDStatusReq
	PageInfo                 = station.PageInfo
	PostsInfo                = station.PostsInfo
	PostsListInfo            = station.PostsListInfo
	PostsReq                 = station.PostsReq
	PostsStationInfo         = station.PostsStationInfo
	PostsStationResp         = station.PostsStationResp
	StationInfo              = station.StationInfo
	StationInfoResp          = station.StationInfoResp
	StationListInfo          = station.StationListInfo
	StationListResp          = station.StationListResp
	StationPostsInfo         = station.StationPostsInfo
	StationPostsRelationInfo = station.StationPostsRelationInfo
	StationPostsResp         = station.StationPostsResp
	StationReq               = station.StationReq
	UUIDReq                  = station.UUIDReq
	UsersInfo                = station.UsersInfo
	UsersListInfo            = station.UsersListInfo
	UsersReq                 = station.UsersReq

	DeliveryLog interface {
		AddDelivery(ctx context.Context, in *DeliveryLogListInfo, opts ...grpc.CallOption) (*BaseDataInfo, error)
		GetDeliveryList(ctx context.Context, in *DeliveryInfo, opts ...grpc.CallOption) (*DeliveryLogListInfo, error)
		AddDeliveryLog(ctx context.Context, in *DeliveryLogInfo, opts ...grpc.CallOption) (*DeliveryLogInfo, error)
		UpdateDeliveryLog(ctx context.Context, in *DeliveryLogInfo, opts ...grpc.CallOption) (*DeliveryLogInfo, error)
		UpdateStatus(ctx context.Context, in *IDStatusReq, opts ...grpc.CallOption) (*BaseDataInfo, error)
		DeleteDeliveryLog(ctx context.Context, in *IDPathReq, opts ...grpc.CallOption) (*BaseDataInfo, error)
		QueryDeliveryLog(ctx context.Context, in *DeliveryLogReq, opts ...grpc.CallOption) (*DeliveryLogListInfo, error)
	}

	defaultDeliveryLog struct {
		cli zrpc.Client
	}
)

func NewDeliveryLog(cli zrpc.Client) DeliveryLog {
	return &defaultDeliveryLog{
		cli: cli,
	}
}

func (m *defaultDeliveryLog) AddDelivery(ctx context.Context, in *DeliveryLogListInfo, opts ...grpc.CallOption) (*BaseDataInfo, error) {
	client := station.NewDeliveryLogClient(m.cli.Conn())
	return client.AddDelivery(ctx, in, opts...)
}

func (m *defaultDeliveryLog) GetDeliveryList(ctx context.Context, in *DeliveryInfo, opts ...grpc.CallOption) (*DeliveryLogListInfo, error) {
	client := station.NewDeliveryLogClient(m.cli.Conn())
	return client.GetDeliveryList(ctx, in, opts...)
}

func (m *defaultDeliveryLog) AddDeliveryLog(ctx context.Context, in *DeliveryLogInfo, opts ...grpc.CallOption) (*DeliveryLogInfo, error) {
	client := station.NewDeliveryLogClient(m.cli.Conn())
	return client.AddDeliveryLog(ctx, in, opts...)
}

func (m *defaultDeliveryLog) UpdateDeliveryLog(ctx context.Context, in *DeliveryLogInfo, opts ...grpc.CallOption) (*DeliveryLogInfo, error) {
	client := station.NewDeliveryLogClient(m.cli.Conn())
	return client.UpdateDeliveryLog(ctx, in, opts...)
}

func (m *defaultDeliveryLog) UpdateStatus(ctx context.Context, in *IDStatusReq, opts ...grpc.CallOption) (*BaseDataInfo, error) {
	client := station.NewDeliveryLogClient(m.cli.Conn())
	return client.UpdateStatus(ctx, in, opts...)
}

func (m *defaultDeliveryLog) DeleteDeliveryLog(ctx context.Context, in *IDPathReq, opts ...grpc.CallOption) (*BaseDataInfo, error) {
	client := station.NewDeliveryLogClient(m.cli.Conn())
	return client.DeleteDeliveryLog(ctx, in, opts...)
}

func (m *defaultDeliveryLog) QueryDeliveryLog(ctx context.Context, in *DeliveryLogReq, opts ...grpc.CallOption) (*DeliveryLogListInfo, error) {
	client := station.NewDeliveryLogClient(m.cli.Conn())
	return client.QueryDeliveryLog(ctx, in, opts...)
}
