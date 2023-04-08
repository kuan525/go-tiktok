package main

import (
	"common/logs"
	"flag"
	"fmt"
	"github.com/kataras/iris/v12"
	"web_feed/conf"
	"web_feed/internal/router"
)

func newApp() *iris.Application {
	app := iris.New()
	router.InitRouters(app)
	return app
}

func main() {
	var strPath string
	flag.StringVar(&strPath, "conf_path", "./conf/configs.yaml", "--conf_path")
	flag.Parse()

	var err error
	if err = conf.InitConfig(strPath); err != nil {
		logs.HandlePaincErr(err, "main:初始化配置文件失败")
	}

	if err = conf.InitLogger(&conf.Cfg.Log); err != nil {
		logs.HandlePaincErr(err, "main:初始化日志失败")
	}

	if conf.Mqcli, err = conf.InitMq(&conf.Cfg.MysqlConf); err != nil {
		logs.HandlePaincErr(err, "main:初始化数据库失败")
	}

	app := newApp()
	addr := fmt.Sprintf("%s:%s", conf.Cfg.HttpAddr.Host, conf.Cfg.HttpAddr.Port)
	// X-Forwarded-For 用于标识原始客户端的 IP 地址
	err = app.Run(iris.Addr(addr), iris.WithRemoteAddrHeader("X-Forwarded-For"))
	if err != nil {
		logs.HandlePaincErr(err, "main:iris启动失败")
	}
}
