package main

import (
	"flag"
	"fmt"
	"runtime"
	"web_user/conf"
	"web_user/internal/router"

	"github.com/kataras/iris/v12"
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

	if err = conf.InitConfig(strPath); err != nil {
		conf.Logger.Infof(err.Error(), "main:初始化配置文件失败")
		panic(err)
	}
	if err = conf.InitLogger(&conf.Cfg.Log); err != nil {
		conf.Logger.Infof(err.Error(), "main:初始化日志失败")
		panic(err)
	}
	if conf.Mqcli, err = conf.InitMq(&conf.Cfg.MysqlConf); err != nil {
		conf.Logger.Infof(err.Error(), "main:初始化数据库失败")
		panic(err)
	}

	app := newApp()
	addr := fmt.Sprintf("%s:%s", conf.Cfg.HttpAddr.Host, conf.Cfg.HttpAddr.Port)
	// X-Forwarded-For 用于标识原始客户端的 IP 地址，代理服务器等通常会将真实ip放在这个里面
	err = app.Run(
		iris.Addr(addr),
		iris.WithRemoteAddrHeader("X-Forwarded-For"),
		// 允许多次消费Body
		iris.WithoutBodyConsumptionOnUnmarshal)
	if err != nil {
		conf.Logger.Infof(err.Error(), "main:iris启动失败")
		panic(err)
	}
}
