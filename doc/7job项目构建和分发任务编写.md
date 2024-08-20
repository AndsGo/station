本次目标我们主要讲解引入job框架和定时任务编写。

## 1.项目初始化

项目结构

```
|-job.go	--main入口
|-etc
	|-station_job.yaml	--配置文件
|-internal
	|-config
		|-config.go	--配置文件实体
	|-logic
		|-这里面各种逻辑实现
	|-svc
		|-servicecontext.go	--服务上下文配置
	|-xxljob	--xxljob配置相关，也可以替换成其他框架实现
		|-xxljob.go
        |-jobhandler -- 自定义的任务处理器
```

整个项目的大致结构如上，其中`xxljob`可以换成其他框架实现。

## 2.引入xxl框架

首先我们得先安装下[xxl-job可执行代码包](https://github.com/AndsGo/xxl-job-sqllite/releases/download/2.3.1/xxl-job-admin-sqllite-2.3.1.jar)，这个包是sqllite版可以无依赖mysql运行。

下载下来后运行

```shell
java -jar xxl-job-admin-sqllite-2.3.1.jar
```

启动完成后我们就可以在http://127.0.0.1:8080/xxl-job-admin/ 访问`xxl-job` ,登录账号密码 **admin/12345678**



查看Example https://github.com/xxl-job/xxl-job-executor-go

通过下面的代码就是引入xxl-job了。将下面得代码复制到

```go
package main

import (
	xxl "github.com/xxl-job/xxl-job-executor-go"
	"github.com/xxl-job/xxl-job-executor-go/example/task"
)

const Port = "9999"

func main() {
	//初始化执行器
	exec := xxl.NewExecutor(
        xxl.ServerAddr("http://127.0.0.1:8080/xxl-job-admin"),
		xxl.AccessToken(""),            //请求令牌(默认为空)
		xxl.ExecutorIp("127.0.0.1"),    //可自动获取
		xxl.ExecutorPort(Port),         //默认9999（此处要与gin服务启动port必需一至）
		xxl.RegistryKey("golang-jobs"), //执行器名称
	)
	exec.Init()
	defer exec.Stop()
	//注册任务handler
	exec.RegTask("task.test", task.Test)
	exec.RegTask("task.test2", task.Test2)
	exec.RegTask("task.panic", task.Panic)
    /启动xxl
    exec.Run()
}
```

我们可以建上面得代码简单运行进行测试。保证xxl可以调用到。

简单看下task.Test方法，后面我们编写的方法需要定义成一样的结构。`func Test(cxt context.Context, param *xxl.RunReq) (msg string)`

```go
func Test(cxt context.Context, param *xxl.RunReq) (msg string) {
	fmt.Println("test one task" + param.ExecutorHandler + " param：" + param.ExecutorParams + " log_id:" + xxl.Int64ToStr(param.LogID))
	return "test done"
}
```

## 3.编写框架基础代码

`etc/station_job.yaml`

```yaml
Mode: console
ServiceName: station-job
Path: ../logs
XxlJob:
  Addresses: http://127.0.0.1/xxl-job-admin
  AccessToken: 
  Appname: station-job
  Port: 9989
  Ip: 127.0.0.1
```

`internal/config/config.go`

```go 
package config

import (
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	StationRpcConf zrpc.RpcClientConf
	XxlJob         struct {
		Addresses   string
		AccessToken string
		Appname     string
		Ip          string
		Port        int
	}
}
```

`internal/svc/servicecontext.go`

```go
package svc

import (
	"job/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}

// 全局svc
var sc *ServiceContext

func NewSvc(c config.Config) {
	sc = NewServiceContext(c)
}
func GetSvc() *ServiceContext {
	return sc
}
```

`internal/xxljob/xxljob.go`

```go
package xxljob

import (
	"job/internal/config"
	"strconv"

	xxl "github.com/xxl-job/xxl-job-executor-go"
	"github.com/xxl-job/xxl-job-executor-go/example/task"
)

func Init(conf config.Config) {
	exec := xxl.NewExecutor(
		xxl.ServerAddr(conf.XxlJob.Addresses),
		xxl.AccessToken(conf.XxlJob.AccessToken),         //请求令牌(默认为空)
		xxl.ExecutorIp(conf.XxlJob.Ip),                   //可自动获取
		xxl.ExecutorPort(strconv.Itoa(conf.XxlJob.Port)), //默认9999（非必填）
		xxl.RegistryKey(conf.XxlJob.Appname),             //执行器名称
	)
	exec.Init()
	defer exec.Stop()
	exec.RegTask("demoJobHandler", task.Test)
	// todo 注册你自己的任务

	exec.Run()
}
```

后续我们的业务任务需要配置在 `// todo 注册你自己的任务` 这里。

`job.go`

```go
package main

import (
	"flag"
	"fmt"
	"job/internal/config"
	"job/internal/svc"
	"job/internal/xxljob"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
)

var c config.Config
var configFile = flag.String("f", "etc/station_job.yaml", "the config file")

func main() {
	//获取配置
	flag.Parse()
	var logc logx.LogConf
	conf.MustLoad(*configFile, &logc)
	// 设置log
	logx.MustSetup(logc)
	conf.MustLoad(*configFile, &c)
	//初始化svc
	svc.NewSvc(c)
	//初始化 xxl
	xxljob.Init(c)
	fmt.Printf("admin job server at %s...\n", c.XxlJob.Ip)
}
```

完成上面的步骤我们一个完整的job项目就搭建起来了，接下来就是业务层代码的编写。

这时我们可以启动job.go进行简单的测试。

## 4.调用rpc

job我们需要调用rpc接口，因此我们需要像api一样引入rpc。

修改station_job.yaml 增加

```yaml
StationRpcConf:
  Target: dns:///127.0.0.1:8080
  Timeout: 10000 #rpc 超时时间
```

在config.go增加rpc的配置

```
StationRpcConf zrpc.RpcClientConf
```

在servicecontext.go中初始化rpc client

```go
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

初始化完成后我们就是可以调用rpc了

## 5.编写定时任务

我们这里以文章分发任务为例。简单介绍下我们的分发逻辑。我们新增一些发布文章到`wordpress`站点的任务，一条记录表示上传一个文章到一个站点。我们可以一次性批量新增发布任务，由于发布时需要调用到很多其他资源。所以短时间内不能完成，因此我们需要有个定时任务进行一部处理。

大致流程如下

![image-20240801142432395](images\image-20240801142432395.png)

下面我们简单编写一个简单的`job`代码。

在`logic`目录下创建一个`delivery`目录，这里面编写分发相关的逻辑

创建 `postsdeliverylogic.go`

```go
package delivery

import (
	"context"

	"job/internal/svc"

	"github.com/xxl-job/xxl-job-executor-go"
	"github.com/zeromicro/go-zero/core/logx"
)

type PostsDeliveryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostsDeliveryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostsDeliveryLogic {
	return &PostsDeliveryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostsDeliveryLogic) Execute(param *xxl.RunReq) error {
	// todo: add your logic here and delete this line
	// 1.获取待处理的队列数据
	// 2.循环处理待处理的数据
	// 2.1.获取文章内容
	// 2.2.组装发送内容
	// 2.3.发送到wordpress
	// 3.3.更新结果到队列表
	l.Logger.Info("PostsDeliveryLogic run end...")
	return nil
}
```

在`xxljob/jobhandler`下创建一个

`delivery.go`

```go
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

```

再将`PostsDelivery`注册到`xxljob`中

`xxljob.go`

```go
package xxljob

import (
	"job/internal/config"
	"job/internal/xxljob/jobhandler"
	"strconv"

	xxl "github.com/xxl-job/xxl-job-executor-go"
	"github.com/xxl-job/xxl-job-executor-go/example/task"
)

func Init(conf config.Config) {
	exec := xxl.NewExecutor(
		xxl.ServerAddr(conf.XxlJob.Addresses),
		xxl.AccessToken(conf.XxlJob.AccessToken),         //请求令牌(默认为空)
		xxl.ExecutorIp(conf.XxlJob.Ip),                   //可自动获取
		xxl.ExecutorPort(strconv.Itoa(conf.XxlJob.Port)), //默认9999（非必填）
		xxl.RegistryKey(conf.XxlJob.Appname),             //执行器名称
	)
	exec.Init()
	defer exec.Stop()
	exec.RegTask("demoJobHandler", task.Test)
	// 注册你自己的任务
	exec.RegTask("PostsDelivery", jobhandler.PostsDelivery)
	exec.Run()
}
```

## 6.测试

进入 http://127.0.0.1:8080/xxl-job-admin/  

1. 配置执行器 station-job
2. 配置任务 PostsDelivery
3. 运行进行测试





