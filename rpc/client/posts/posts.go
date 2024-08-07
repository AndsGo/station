// Code generated by goctl. DO NOT EDIT.
// Source: station.proto

package posts

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

	Posts interface {
		AddPosts(ctx context.Context, in *PostsInfo, opts ...grpc.CallOption) (*PostsInfo, error)
		UpdatePosts(ctx context.Context, in *PostsInfo, opts ...grpc.CallOption) (*PostsInfo, error)
		DeletePosts(ctx context.Context, in *IDPathReq, opts ...grpc.CallOption) (*BaseDataInfo, error)
		QueryPosts(ctx context.Context, in *PostsReq, opts ...grpc.CallOption) (*PostsListInfo, error)
		GetPosts(ctx context.Context, in *IDPathReq, opts ...grpc.CallOption) (*PostsInfo, error)
		QueryStation(ctx context.Context, in *IDPathReq, opts ...grpc.CallOption) (*PostsStationResp, error)
	}

	defaultPosts struct {
		cli zrpc.Client
	}
)

func NewPosts(cli zrpc.Client) Posts {
	return &defaultPosts{
		cli: cli,
	}
}

func (m *defaultPosts) AddPosts(ctx context.Context, in *PostsInfo, opts ...grpc.CallOption) (*PostsInfo, error) {
	client := station.NewPostsClient(m.cli.Conn())
	return client.AddPosts(ctx, in, opts...)
}

func (m *defaultPosts) UpdatePosts(ctx context.Context, in *PostsInfo, opts ...grpc.CallOption) (*PostsInfo, error) {
	client := station.NewPostsClient(m.cli.Conn())
	return client.UpdatePosts(ctx, in, opts...)
}

func (m *defaultPosts) DeletePosts(ctx context.Context, in *IDPathReq, opts ...grpc.CallOption) (*BaseDataInfo, error) {
	client := station.NewPostsClient(m.cli.Conn())
	return client.DeletePosts(ctx, in, opts...)
}

func (m *defaultPosts) QueryPosts(ctx context.Context, in *PostsReq, opts ...grpc.CallOption) (*PostsListInfo, error) {
	client := station.NewPostsClient(m.cli.Conn())
	return client.QueryPosts(ctx, in, opts...)
}

func (m *defaultPosts) GetPosts(ctx context.Context, in *IDPathReq, opts ...grpc.CallOption) (*PostsInfo, error) {
	client := station.NewPostsClient(m.cli.Conn())
	return client.GetPosts(ctx, in, opts...)
}

func (m *defaultPosts) QueryStation(ctx context.Context, in *IDPathReq, opts ...grpc.CallOption) (*PostsStationResp, error) {
	client := station.NewPostsClient(m.cli.Conn())
	return client.QueryStation(ctx, in, opts...)
}
