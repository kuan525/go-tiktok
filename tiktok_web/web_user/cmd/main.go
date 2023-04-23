package main

import (
	"common/conf"
	"common/initialize"
	"common/log"
	"flag"
	"fmt"
	"github.com/kataras/iris/v12"
	"runtime"
	"web_user/internal/router"
)

func newApp() *iris.Application {
	app := iris.New()
	router.InitRouters(app)
	return app
}

func main() {
	var strPath string
	if runtime.GOOS == "darwin" {
		flag.StringVar(&strPath, "conf_path", "./conf/configs.yaml", "--conf_path")
	} else {
		flag.StringVar(&strPath, "conf_path", "./configs.yaml", "--conf_path")
	}

	flag.Parse()
	var err error

	initialize.InitConfig(strPath)
	log.InitLogger(&conf.Cfg.Log)
	initialize.InitMq(&conf.Cfg.MysqlConf)

	app := newApp()
	addr := fmt.Sprintf("%s:%s", conf.Cfg.HttpAddr.Host, conf.Cfg.HttpAddr.Port)
	// X-Forwarded-For 用于标识原始客户端的 IP 地址，代理服务器等通常会将真实ip放在这个里面
	err = app.Run(
		iris.Addr(addr),
		iris.WithRemoteAddrHeader("X-Forwarded-For"),
		// 允许多次消费Body
		iris.WithoutBodyConsumptionOnUnmarshal)

	if err != nil {
		log.Logger.Errorf(err.Error(), "main:iris启动失败")
		panic(err)
	}
}
