// Code generated by goctl. DO NOT EDIT.
// Source: station.proto

package users

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

	Users interface {
		AddUsers(ctx context.Context, in *UsersInfo, opts ...grpc.CallOption) (*UsersInfo, error)
		UpdateUsers(ctx context.Context, in *UsersInfo, opts ...grpc.CallOption) (*UsersInfo, error)
		DeleteUsers(ctx context.Context, in *IDPathReq, opts ...grpc.CallOption) (*BaseDataInfo, error)
		QueryUsers(ctx context.Context, in *UsersReq, opts ...grpc.CallOption) (*UsersListInfo, error)
		GetUsers(ctx context.Context, in *IDPathReq, opts ...grpc.CallOption) (*UsersInfo, error)
	}

	defaultUsers struct {
		cli zrpc.Client
	}
)

func NewUsers(cli zrpc.Client) Users {
	return &defaultUsers{
		cli: cli,
	}
}

func (m *defaultUsers) AddUsers(ctx context.Context, in *UsersInfo, opts ...grpc.CallOption) (*UsersInfo, error) {
	client := station.NewUsersClient(m.cli.Conn())
	return client.AddUsers(ctx, in, opts...)
}

func (m *defaultUsers) UpdateUsers(ctx context.Context, in *UsersInfo, opts ...grpc.CallOption) (*UsersInfo, error) {
	client := station.NewUsersClient(m.cli.Conn())
	return client.UpdateUsers(ctx, in, opts...)
}

func (m *defaultUsers) DeleteUsers(ctx context.Context, in *IDPathReq, opts ...grpc.CallOption) (*BaseDataInfo, error) {
	client := station.NewUsersClient(m.cli.Conn())
	return client.DeleteUsers(ctx, in, opts...)
}

func (m *defaultUsers) QueryUsers(ctx context.Context, in *UsersReq, opts ...grpc.CallOption) (*UsersListInfo, error) {
	client := station.NewUsersClient(m.cli.Conn())
	return client.QueryUsers(ctx, in, opts...)
}

func (m *defaultUsers) GetUsers(ctx context.Context, in *IDPathReq, opts ...grpc.CallOption) (*UsersInfo, error) {
	client := station.NewUsersClient(m.cli.Conn())
	return client.GetUsers(ctx, in, opts...)
}
