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
