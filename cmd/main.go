package main

import (
	"flag"
	"fmt"
	"github.com/kataras/iris/v12"
	"go-tiktok/conf"
	"go-tiktok/db"
	"go-tiktok/internal/log"
	"go-tiktok/internal/router"
)

func newApp() *iris.Application {
	app := iris.New()
	//application.PreSettion(app)

	router.InitRouters(app)
	return app
}

func main() {
	var strPath string
	flag.StringVar(&strPath, "conf_path", "./conf/configs.yaml", "--conf_path")
	flag.Parse()

	var err error
	if err = conf.InitConfig(strPath); err != nil {
		conf.PaincErr(err, "main:初始化配置文件失败")
	}

	if err = log.InitLogger(&conf.Cfg.Log); err != nil {
		conf.PaincErr(err, "main:初始化日志失败")
	}

	if conf.Mqcli, err = db.InitMq(&conf.Cfg.MysqlConf); err != nil {
		conf.PaincErr(err, "main:初始化数据库失败")
	}

	app := newApp()
	addr := fmt.Sprintf("%s:%s", conf.Cfg.HttpAddr.Port, conf.Cfg.HttpAddr.Port)
	// X-Forwarded-For 用于标识原始客户端的 IP 地址
	err = app.Run(iris.Addr(addr), iris.WithRemoteAddrHeader("X-Forwarded-For"))
	if err != nil {
		conf.PaincErr(err, "main:iris启动失败")
	}
}
