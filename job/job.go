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
