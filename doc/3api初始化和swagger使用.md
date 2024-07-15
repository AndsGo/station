[代码分支](https://github.com/AndsGo/station/tree/3api%E4%BB%A3%E7%A0%81%E7%94%9F%E6%88%90%E5%92%8Cswagger%E4%BD%BF%E7%94%A8)

# api初始化和swagger使用

目标实现api编写和swagger使用

## 0.本次需要使用到的脚手架命令

生成 http server 代码

```shell
goctl api go -api all.api -dir ..
```

生成swagger文档

```shell
goctl api plugin -plugin goctl-swagger="swagger -filename station.json -host 127.0.0.1:8888" -api all.api -dir .

```

运行 swagger 

```shell
swagger serve --no-open -F=swagger --port 36666 station.json
```

## 1.编写api

api详细[文档](https://go-zero.dev/docs/tutorials)

**base.api** api公共types

```go
syntax = "v1"

// The basic response with data | 基础带数据信息
type BaseDataInfo {
    Code int    `json:"code"`           // Error code | 错误代码
    Message  string `json:"message"`    // Message | 提示信息
    Data string `json:"data,omitempty"` // Data | 数据
}
// The basic response with data | 基础带数据信息
type BaseListInfo {
    Total uint64 `json:"total"`          // The total number of data | 数据总数
    Data string `json:"data,omitempty"`  // Data | 数据
}
// The basic response without data | 基础不带数据信息
type BaseMsgResp {
    Code int    `json:"code"`          // Error code | 错误代码
    Message  string `json:"message"`   // Message | 提示信息
}
// The page request parameters | 列表请求参数
type PageInfo {
    Page   uint64    `form:"page" validate:"required,number,gt=0"`             // Page number | 第几页
    PageSize  uint64    `form:"pageSize" validate:"required,number,lt=100000"` // Page size | 单页数据行数
}
// Basic ID request | 基础ID参数请求
type IDReq {
    Id  uint64 `json:"id" validate:"number"` // ID Required: true
}
// Basic IDs request | 基础ID数组参数请求
type IDsReq {
    Ids  []uint64 `json:"ids"` // IDs Required: true
}
// Basic ID request | 基础ID地址参数请求
type IDPathReq {
    Id  uint64 `path:"id"` // ID  Required: true
}
```

**all.api** 用于聚合api

```go
import "base.api"
import "./station/station.api"
import "./station/posts.api"
import "./station/delivery_log.api"
```

**station.api**   站点api

```go
//站点信息
type  (
	// Station 
	StationInfo {
		Id           uint64  `json:"id,optional"`            // 主键
		DomainName   string  `json:"domainName"`   // 域名
		Ip           string  `json:"ip,optional"`   // 域名
		DomainYear   int64   `json:"domainYear"`   // 域名年份
		GoogleWeight float64 `json:"googleWeight"` // 谷歌权重
		Type         string  `json:"type"`          // 网站类型
		Industry     string  `json:"industry"`      // 网站行业
		ArticlesNum  int64   `json:"articlesNum"`     // 文章数量
		UserName     string  `json:"userName,optional"`     // 账号名
		PassWord     string  `json:"passWord,optional"`     // 密码
	}
	// Station 页面查询
	StationReq {
		PageInfo
		DomainName   string  `form:"domainName,optional"`   // 域名
		Ip           string  `form:"ip,optional"`          // ip
		DomainYear   int64   `form:"domainYear,optional"`   // 域名年份
		GoogleWeight string  `form:"googleWeight,optional"` // 谷歌权重
		Type         string  `form:"type,optional"`          // 网站类型
		Industry     string  `form:"industry,optional"`      // 网站行业
	}
	// The response data of Station list | Station 列表数据
	StationListInfo {
		BaseListInfo
		Data []StationInfo `json:"data"` // StationInfo list data | StationInfo列表数据
	}
	//Station 列表返回体
	StationListResp {
		BaseDataInfo
		Data StationListInfo `json:"data"` // Station list data | Station列表数据
	}
	// Station 普通返回体
	StationInfoResp {
        BaseDataInfo
        Data StationInfo `json:"data"` // Station information | Station数据
    }
	// Station Posts 返回体  
	StationPostsInfo {
       Id         uint64 `json:"id"`
	   Title      string `json:"title"`       // Title
    }
	// 返回体
	StationPostsResp{
		BaseDataInfo
		Data []StationPostsInfo `json:"data"`  // Station Posts 返回体
	}
)  
@server (
	group:      station
	prefix: 	/station
	timeout:    10s
)
service Station {
	@doc "新增站点"
	@handler addStation
	post /api/station (StationInfo) returns (StationInfoResp)
	@doc "修改站点"
	@handler updateStation
	put /api/station (StationInfo) returns (StationInfoResp)
	@doc "删除站点"
	@handler deleteStation
	delete /api/station/:id (IDPathReq) returns (BaseDataInfo)
	@doc "查询站点"
	@handler queryStation
	get /api/station (StationReq) returns (StationListResp)
    @doc "获取关联的文章"
	@handler queryPosts
	get /api/station/posts/:id (IDPathReq) returns (StationPostsResp)
}
```

**posts.api ** 文章api

```go
//博客信息
type (
	// Posts
	PostsInfo {
		Id         uint64 `json:"id,optional"`
		Title      string `json:"title"` // 标题
		Source     string `json:"source"` // 来源
		Author     int64 `json:"author"` // 作者
		ThrownNum  int64  `json:"thrownNum"` // 投放数量
		Categories  string  `json:"categories"` // 分类
		CreateTime string `json:"createTime,optional"` //时间
		Content    string `json:"content"` //详情
	}
	// Posts 页面查询
	PostsReq {
		PageInfo
		Title      string `form:"title,optional"` // 标题
		Source     string `form:"source,optional"` // 来源
		Categories     string `form:"categories,optional"` // 分类
		Author     int64 `form:"author,optional"` // 作者
		CreateTime int64  `form:"createTime,optional"` // 时间
	}
	// The response data of Posts list | Posts 列表数据
	PostsListInfo {
		BaseListInfo
		Data []PostsInfo `json:"data"` // PostsInfo list data | PostsInfo列表数据
	}
	//Posts 列表返回体
	PostsListResp {
		BaseDataInfo
		Data PostsListInfo `json:"data"` // Posts list data | Posts列表数据
	}
	// Posts 普通返回体
	PostsInfoResp {
		BaseDataInfo
		Data PostsInfo `json:"data"` // Posts information | Posts数据
	}
)
@server (
	group:      posts
	prefix:     /station
)
service Station {
	@doc "新增Posts"
	@handler addPosts
	post /api/posts (PostsInfo) returns (PostsInfoResp)

	@doc "修改Posts"
	@handler updatePosts
	put /api/posts (PostsInfo) returns (PostsInfoResp)

	@doc "删除Posts"
	@handler deletePosts
	delete /api/posts/:id (IDPathReq) returns (BaseDataInfo)

    @doc "查询Posts"
	@handler queryPosts
	get /api/posts (PostsReq) returns (PostsListResp)

	@doc "查询Posts详情"
	@handler getPosts
	get /api/posts/:id (IDPathReq) returns (PostsInfoResp)
}
```

**delivery_log.api** 文章分发日志api

```go
//博客信息
type (
	// DeliveryLog
	DeliveryLogInfo {
		Id           uint64 `json:"id,optional"`
		Title        string `json:"title"` // 标题
		Source       string `json:"source,optional"` // 来源
		DomainName   string `json:"domainName"` // 域名
		DeliveryDate string `json:"deliveryDate,optional"` // 投放日期
		Deliverer    string `json:"deliverer,optional"` // 投放人
		Status       int64  `json:"status,optional"` // 投放状态
		Author       uint64 `json:"author,optional"` // 作者
		WpCateIds    string `json:"wpCateIds,optional"` // wp分类
		StationId    uint64 `json:"stationId,optional"` // 站点id
		PostsId      uint64 `json:"postsId,optional"` // 文章id
	}
	// 投放对象
	DeliveryInfo {
		StationInfoList []StationInfo `json:"stationInfoList"` // 站点id
		PostsInfoList   []PostsInfo   `json:"postsInfoList"` // 文章id
	}
	// DeliveryLog 页面查询
	DeliveryLogReq {
		PageInfo
		Title        string `form:"title,optional"` // 标题
		Source       string `form:"source,optional"` // 来源
		DomainName   string `form:"domainName,optional"` // 域名
		DeliveryDate int64  `form:"deliveryDate,optional"` // 投放日期
		Deliverer    string `form:"deliverer,optional"` // 投放人
		Status       int64  `form:"status,optional"` // 投放状态
		Author       uint64 `form:"author,optional"` // 作者
	}
	// The response data of DeliveryLog list | DeliveryLog 列表数据
	DeliveryLogListInfo {
		BaseListInfo
		Data []DeliveryLogInfo `json:"data"`	// DeliveryLogInfo list data | DeliveryLogInfo列表数据
	}
	// The response data of DeliveryLog list | DeliveryLog 列表数据
	DeliveryListInfo {
		Data []DeliveryLogInfo `json:"data"`	// DeliveryInfo list data | DeliveryInfo列表数据
	}
	//DeliveryLog 列表返回体
	DeliveryLogListResp {
		BaseDataInfo
		Data DeliveryLogListInfo `json:"data"` // DeliveryLog list data | DeliveryLog列表数据
	}
	// DeliveryLog 普通返回体
	DeliveryLogInfoResp {
		BaseDataInfo
		Data DeliveryLogInfo `json:"data"`  // DeliveryLog information | DeliveryLog数据
	}
)

@server (
	group:      deliveryLog
	prefix:     /station
)
service Station {
	@doc "投放" 
	@handler addDeliveryLog
	post /api/deliverylog (DeliveryListInfo) returns (BaseDataInfo)

	@doc "修改DeliveryLog"
	@handler updateDeliveryLog
	put /api/deliverylog (DeliveryLogInfo) returns (DeliveryLogInfoResp)

	@doc "删除DeliveryLog"
	@handler deleteDeliveryLog
	delete /api/deliverylog/:id (IDPathReq) returns (BaseDataInfo)

	@doc "查询DeliveryLog"
	@handler queryDeliveryLog
	get /api/deliverylog (DeliveryLogReq) returns (DeliveryLogListResp)

	@doc "获取投放列表"
	@handler generateDeliveryList
	post /api/deliverylog/list (DeliveryInfo) returns (DeliveryLogListResp)
}
```

## 2.生成代码

进入到api项目的desc目录，运行生成api代码脚手架命令

```go
cd ./desc
goctl api go -api all.api -dir ..
```

## 3.swagger使用

还是在desc目录下，生成swaagger json

```go
goctl api plugin -plugin goctl-swagger="swagger -filename station.json -host 127.0.0.1:8000" -api all.api -dir .
```

运行swagger

```go
swagger serve --no-open -F=swagger --port 36666 station.json
```

## 4.允许跨域配置

在入口文件 station.go中修改

```
//rest.MustNewServer(c.RestConf)
//修改为下面的代码
rest.MustNewServer(c.RestConf, rest.WithCors())
```

## 5.测试api

运行rpc和api后。

启动完swagger，我们可以在swagger上面进行api测试了。

访问 http://127.0.0.1:36666/docs 页面，可以直接在页面上进行测试。