// Code generated by goctl. DO NOT EDIT.
// Source: station.proto

package stationpostsrelation

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

	StationPostsRelation interface {
		AddStationPostsRelation(ctx context.Context, in *StationPostsRelationInfo, opts ...grpc.CallOption) (*StationPostsRelationInfo, error)
		UpdateStationPostsRelation(ctx context.Context, in *StationPostsRelationInfo, opts ...grpc.CallOption) (*StationPostsRelationInfo, error)
		DeleteStationPostsRelation(ctx context.Context, in *IDPathReq, opts ...grpc.CallOption) (*BaseDataInfo, error)
	}

	defaultStationPostsRelation struct {
		cli zrpc.Client
	}
)

func NewStationPostsRelation(cli zrpc.Client) StationPostsRelation {
	return &defaultStationPostsRelation{
		cli: cli,
	}
}

func (m *defaultStationPostsRelation) AddStationPostsRelation(ctx context.Context, in *StationPostsRelationInfo, opts ...grpc.CallOption) (*StationPostsRelationInfo, error) {
	client := station.NewStationPostsRelationClient(m.cli.Conn())
	return client.AddStationPostsRelation(ctx, in, opts...)
}

func (m *defaultStationPostsRelation) UpdateStationPostsRelation(ctx context.Context, in *StationPostsRelationInfo, opts ...grpc.CallOption) (*StationPostsRelationInfo, error) {
	client := station.NewStationPostsRelationClient(m.cli.Conn())
	return client.UpdateStationPostsRelation(ctx, in, opts...)
}

func (m *defaultStationPostsRelation) DeleteStationPostsRelation(ctx context.Context, in *IDPathReq, opts ...grpc.CallOption) (*BaseDataInfo, error) {
	client := station.NewStationPostsRelationClient(m.cli.Conn())
	return client.DeleteStationPostsRelation(ctx, in, opts...)
}
