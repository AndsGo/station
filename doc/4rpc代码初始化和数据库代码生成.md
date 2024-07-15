[代码分支](https://github.com/AndsGo/station/tree/4rpc%E4%BB%A3%E7%A0%81%E5%88%9D%E5%A7%8B%E5%8C%96%E5%92%8C%E6%95%B0%E6%8D%AE%E5%BA%93%E4%BB%A3%E7%A0%81%E7%94%9F%E6%88%90)

rpc代码初始化和数据库代码生成

## 0.本次需要使用到的脚手架命令

生成 grpc server 代码

```shell
goctl rpc protoc station.proto --go_out=. --go-grpc_out=. --zrpc_out=. -m
```

生成 mysql model

加上 cache --cache 可以增加缓存加快查询

```shell
goctl model mysql ddl --src station.sql --dir .
```

## 1.编写.proto

protobuf[文档](https://protobuf.dev/getting-started/gotutorial/)

创建station.proto文件，写入如下内容

```protobuf
syntax = "proto3";

option go_package = "./station";

package station;

// The basic response with data
message BaseDataInfo {
    int32 code = 1;
    string message = 2;
    string data = 3;
}
// The basic response with data
message BaseListInfo {
    uint64 total = 1;
    string data = 2;
}
// The page request parameters
message PageInfo {
    uint64 page = 1;
    uint64 pageSize = 2;
}
// ID Status request
message IDStatusReq {
    uint64 id = 1;
    int64 status = 2;
    string result = 3; // 执行结果
}


// Basic ID request
message IDPathReq {
    uint64 id = 1;
}

// Basic UUID request
message UUIDReq {
    string id = 1;
}

// Station information
message StationInfo {
    uint64 id = 1;
    string domain_name = 2;
    int64 domain_year = 3;
    double google_weight = 4;
    string type = 5;
    string industry = 6;
    string user_name = 7;
    string pass_word = 8;
    int64 articles_num = 9;
    string ip = 10;
}

// Station page query
message StationReq {
    PageInfo page_info = 1;
    string domain_name = 2;
    int64 domain_year = 3;
    string google_weight = 4;
    string type = 5;
    string industry = 6;
    string ip = 7;
}

// The response data of Station list
message StationListInfo {
    BaseListInfo base_list_info = 1;
    repeated StationInfo data = 2;
}

// Station list response
message StationListResp {
    BaseDataInfo base_data_info = 1;
    StationListInfo data = 2;
}

// Station normal response
message StationInfoResp {
    BaseDataInfo base_data_info = 1;
    StationInfo data = 2;
}

// Posts information
message PostsInfo {
    uint64 id = 1;
    string title = 2;
    string source = 3;
    int64 author = 4;
    int64 thrown_num = 5;
    string create_time = 6;
    string content = 7;
    string media = 8;
    string categories = 9;
}

// Posts page query
message PostsReq {
    PageInfo page_info = 1;
    string title = 2;
    string source = 3;
    int64 author = 4;
    string create_time = 5;
    string categories = 6;
}

// The response data of Posts list
message PostsListInfo {
    BaseListInfo base_list_info = 1;
    repeated PostsInfo data = 2;
}

// DeliveryLog information
message DeliveryLogInfo {
    uint64 id = 1;
    string title = 2;
    string source = 3;
    string domain_name = 4;
    string delivery_date = 5;
    string deliverer = 6;
    int64 status = 7;
    uint64 author = 8;
    uint64 posts_id =9;
	uint64 station_id =10;
    string wp_cate_ids =11;
}

message DeliveryInfo {
	repeated StationInfo station_info_list =1;
	repeated PostsInfo posts_info_list =2;
    string deliverer = 3;
}

// DeliveryLog page query
message DeliveryLogReq {
    PageInfo page_info = 1;
    string title = 2;
    string source = 3;
    string domain_name = 4;
    string delivery_date = 5;
    string deliverer = 6;
    int64 status = 7;
    int64 author = 8;
}

// The response data of DeliveryLog list
message DeliveryLogListInfo {
    BaseListInfo base_list_info = 1;
    repeated DeliveryLogInfo data = 2;
}

// StationPostsRelation information
message StationPostsRelationInfo{
    uint64 id = 1;
    uint64 station_id =2;
    uint64 posts_id =3;
    int64 wp_id =4;
}

message StationPostsInfo{
    uint64 id = 1;
    string title = 2;   
}

// Station Posts  返回体
message	StationPostsResp {
    repeated StationPostsInfo data = 1;
}

message PostsStationInfo {
    uint64 id = 1;
    string domain_name = 2;
}
// Posts Station 返回体
message	PostsStationResp {
    repeated PostsStationInfo data = 1;
}

// Users
message UsersInfo {
  uint64 id = 1; // 可选
  string name = 2; // 姓名，页面显示
  string email = 3; // 邮箱
  string slug = 4; // 用户标识符
  string username = 5; // 登录名称
  string password = 6; // 登录密码，实际应用中应避免直接传输明文密码
}
// Users 页面查询
message UsersReq {
  string name = 1; // 姓名
  string email = 2; // 邮箱
  string slug = 3; // 用户标识符
  string create_time = 4; // 创建时间
  PageInfo page_info = 5; // 分页信息，假设 PageInfo 已定义
}
// The response data of Users list | Users 列表数据
message UsersListInfo {
    BaseListInfo base_list_info = 1;
    repeated UsersInfo data =2;
}

service Users {
    rpc addUsers(UsersInfo) returns (UsersInfo);
    rpc updateUsers(UsersInfo) returns (UsersInfo);
    rpc deleteUsers(IDPathReq) returns (BaseDataInfo);
    rpc queryUsers(UsersReq) returns (UsersListInfo);
    rpc getUsers(IDPathReq) returns (UsersInfo);
}

service Station {
    rpc addStation(StationInfo) returns (StationInfo);
    rpc updateStation(StationInfo) returns (StationInfo);
    rpc deleteStation(IDPathReq) returns (BaseDataInfo);
    rpc queryStation(StationReq) returns (StationListInfo);
    rpc getStationAll(IDPathReq) returns (StationListInfo);
    rpc getStation(IDPathReq) returns (StationInfo);
    rpc queryPosts(IDPathReq) returns (StationPostsResp);
    
}

service Posts {
    rpc addPosts(PostsInfo) returns (PostsInfo);
    rpc updatePosts(PostsInfo) returns (PostsInfo);
    rpc deletePosts(IDPathReq) returns (BaseDataInfo);
    rpc queryPosts(PostsReq) returns (PostsListInfo);
    rpc getPosts(IDPathReq) returns (PostsInfo);
    rpc queryStation(IDPathReq) returns (PostsStationResp);
}

service DeliveryLog {
    rpc addDelivery(DeliveryLogListInfo) returns (BaseDataInfo);
    rpc getDeliveryList(DeliveryInfo) returns (DeliveryLogListInfo);
    rpc addDeliveryLog(DeliveryLogInfo) returns (DeliveryLogInfo);
    rpc updateDeliveryLog(DeliveryLogInfo) returns (DeliveryLogInfo);
    rpc updateStatus(IDStatusReq) returns (BaseDataInfo);
    rpc deleteDeliveryLog(IDPathReq) returns (BaseDataInfo);
    rpc queryDeliveryLog(DeliveryLogReq) returns (DeliveryLogListInfo);
}

service StationPostsRelation {
	rpc addStationPostsRelation(StationPostsRelationInfo) returns (StationPostsRelationInfo);
	rpc updateStationPostsRelation(StationPostsRelationInfo) returns (StationPostsRelationInfo);
	rpc deleteStationPostsRelation(IDPathReq) returns (BaseDataInfo);
}
```

## 2.生成rpc代码

进入rpc项目 ，执行

```
goctl rpc protoc station.proto --go_out=. --go-grpc_out=. --zrpc_out=. -m
```

## 3.生成数据库CRUD代码

我们这里使用 sql方式生成

创建目录

```shell
mkdir /model
cd model
```

创建sql文件

station.sql

```sql
CREATE TABLE `station` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `domain_name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '域名',
  `ip` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'ip',
  `domain_year` int(11) DEFAULT NULL COMMENT '域名年份',
  `google_weight` decimal(10,0) DEFAULT NULL COMMENT '谷歌权重',
  `type` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '网站类型',
  `industry` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '网站行业',
  `articles_num` int(11) DEFAULT NULL COMMENT '文章数量',
  `user_name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '账号名',
  `pass_word` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '密码',
  `create_at` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
  `update_at` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='站点信息';

CREATE TABLE `posts` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '标题',
  `source` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '来源',
  `author` int(11) DEFAULT NULL COMMENT '作者',
  `thrown_num` int(11) DEFAULT NULL COMMENT '投放数量',
  `categories` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '分类',
  `creater` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '创建人',
  `create_at` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
  `modifier` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '修改人',
  `update_at` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='站点博客内容';

CREATE TABLE `station_posts_relation` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `station_id` int(10) unsigned NOT NULL COMMENT '站点id',
  `posts_id` int(10) unsigned NOT NULL COMMENT '内容id',
  `wp_id` int(10) unsigned DEFAULT '0' COMMENT 'wp id',
  PRIMARY KEY (`id`),
  KEY `uk_posts_station` (`posts_id`,`station_id`),
  KEY `idx_station_id` (`station_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='站点与post关系表';

CREATE TABLE `delivery_log` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `title` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '标题',
  `source` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '来源',
  `domain_name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '域名',
  `delivery_date` date DEFAULT NULL COMMENT '投放日期',
  `author` int(11) DEFAULT NULL COMMENT '作者',
  `deliverer` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '投放人',
  `status` tinyint(2) DEFAULT NULL COMMENT '投放状态 0待投放 1已投放 2 投放失败',
  `result` text COLLATE utf8mb4_unicode_ci COMMENT '投放执行结果',
  `create_time` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
  `update_time` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '修改时间',
  `posts_id` int(11) DEFAULT NULL COMMENT '文章id',
  `station_id` int(11) DEFAULT NULL COMMENT '站点id',
  `wp_cate_ids` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'wp分类',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_posts_station` (`posts_id`,`station_id`),
  KEY `idx_status` (`status`) USING BTREE COMMENT '用来获取为0的数据初始化'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='投放日志';

CREATE TABLE `station_posts_text` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `posts_id` int(11) DEFAULT NULL COMMENT ' posts 主键',
  `content` mediumtext COLLATE utf8mb4_unicode_ci COMMENT 'html内容',
  `create_time` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
  `update_time` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_posts_id` (`posts_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='站点博客内容大字段';

```

生成数据库代码

```shell
cd /model
goctl model mysql ddl --src station.sql --dir .
```

需要注意到一点，goctl model会对生成的文件进行覆盖，所有后续对表有调整需要手动修改。

## 4.测试项目

配置数据库, 在配置文件中增加DB配置

```yaml
Name: station.rpc
ListenOn: 0.0.0.0:8080
DB:
  DataSource: root:123456@tcp(127.0.0.1:3306)/station?charset=utf8mb4&parseTime=True&loc=Local
```

config.go中增加

```go
type Config struct {
	zrpc.RpcServerConf
	DB         struct {
		DataSource string
	}
}
```

将model 注入svc,修改 servicecontext.go文件

```go
package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"rpc/station/internal/config"
	"rpc/station/model"
)

type ServiceContext struct {
	Config                   config.Config
	DeliveryLogModel          model.DeliveryLogModel
	StationModel              model.StationModel
	PostsModel                model.PostsModel
	StationPostsTextModel     model.StationPostsTextModel
	StationPostsRelationModel model.StationPostsRelationModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	dbCon := sqlx.NewMysql(c.DB.DataSource)
	svc := &ServiceContext{
		Config:                   c,
		DeliveryLogModel:          model.NewDeliveryLogModel(cnn),
		StationModel:              model.NewStationModel(cnn),
		PostsModel:                model.NewPostsModel(cnn),
		StationPostsTextModel:     model.NewStationPostsTextModel(cnn),
		StationPostsRelationModel: model.NewStationPostsRelationModel(cnn),
	}
	return svc
}
```

// 注册服务 station.go

```go
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

```

进入rpc主目录

```shell
go run station.go
```

### api连接rpc

修改 api的 station.yaml

```yaml
Name: Station
Host: 0.0.0.0
Port: 8888
Timeout: 10000 #api 超时时间
StationRpcConf:
  Target: dns:///127.0.0.1:8080
  Timeout: 10000 #rpc 超时时间
```

将rpc 注入svc,修改 servicecontext.go文件

```go
package svc

import (
	"api/internal/config"
	"rpc/client/deliverylog"
	"rpc/client/posts"
	"rpc/client/station"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	//rpc  类似bean管理
	StationRpc    station.Station
	PostsRpc      posts.Posts
	DeliverLogRpc deliverylog.DeliveryLog
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		StationRpc:    station.NewStation(zrpc.MustNewClient(c.RpcConf)),
		PostsRpc:      posts.NewPosts(zrpc.MustNewClient(c.RpcConf)),
		DeliverLogRpc: deliverylog.NewDeliveryLog(zrpc.MustNewClient(c.RpcConf)),
	}
}

```

### 编写业务代码

修改 api\internal\logic\posts\addpostslogic.go

```go
func (l *AddPostsLogic) AddPosts(req *types.PostsInfo) (resp *types.PostsInfoResp, err error) {
	info, err := l.svcCtx.PostsRpc.AddPosts(l.ctx, &station.PostsInfo{
		Id:         req.Id,
		Content:    req.Content,
		Title:      req.Title,
		Source:     req.Source,
		Author:     req.Author,
		ThrownNum:  req.ThrownNum,
		Categories: req.Categories,
		CreateTime: req.CreateTime,
	})
	if err != nil {
		return nil, err
	}
	return &types.PostsInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code:    200,
			Message: "success",
		},
		Data: types.PostsInfo{
			Id:         info.Id,
			Content:    info.Content,
			Title:      info.Title,
			Source:     info.Source,
			Author:     info.Author,
			ThrownNum:  info.ThrownNum,
			Categories: info.Categories,
			CreateTime: info.CreateTime,
		},
	}, nil
}
```

### 测试数据库连接

修改 rpc\internal\logic\posts\addpostslogic.go

```go
func (l *AddPostsLogic) AddPosts(in *station.PostsInfo) (*station.PostsInfo, error) {
	info := &model.Posts{
		Title:      sql.NullString{String: in.Title, Valid: true},
		Source:     sql.NullString{String: in.Source, Valid: true},
		Author:     sql.NullInt64{Int64: in.Author, Valid: true},
		ThrownNum:  sql.NullInt64{Int64: in.ThrownNum, Valid: true},
		Categories: sql.NullString{String: in.Categories, Valid: true},
	}
	r, err := l.svcCtx.PostsModel.Insert(l.ctx, info)
	if err != nil {
		return nil, err
	}
	id, err := r.LastInsertId()
	if err != nil {
		return nil, err
	}
	result := &station.PostsInfo{
		Id:         uint64(id),
		Title:      info.Title.String,
		Source:     info.Source.String,
		Author:     info.Author.Int64,
		ThrownNum:  info.ThrownNum.Int64,
		Categories: info.Categories.String,
	}
	return result, nil
}
```

### 启动swagger进行测试

进入api的desc目录

```go
swagger serve --no-open -F=swagger --port 36666 station.json
```

启动完成后即可在swagger页面进行测试