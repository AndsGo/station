package main

import (
	"flag"
	"fmt"

	"rpc/internal/config"
	deliverylogServer "rpc/internal/server/deliverylog"
	postsServer "rpc/internal/server/posts"
	stationServer "rpc/internal/server/station"
	stationPostsRelationServer "rpc/internal/server/stationpostsrelation"
	usersServer "rpc/internal/server/users"
	"rpc/internal/svc"
	"rpc/station"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/station.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		station.RegisterDeliveryLogServer(grpcServer, deliverylogServer.NewDeliveryLogServer(ctx))
		station.RegisterStationServer(grpcServer, stationServer.NewStationServer(ctx))
		station.RegisterPostsServer(grpcServer, postsServer.NewPostsServer(ctx))
		station.RegisterUsersServer(grpcServer, usersServer.NewUsersServer(ctx))
		station.RegisterStationPostsRelationServer(grpcServer, stationPostsRelationServer.NewStationPostsRelationServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
