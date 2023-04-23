package conf

import (
	"common/log"
	"runtime"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

func InitMq(mq *MysqlConf) {
	var dataSourceName string
	if runtime.GOOS == "darwin" && runtime.GOARCH == "arm64" { // 本机
		dataSourceName = "root:mysql_112525@tcp(localhost:3306)/tiktok?charset=utf8mb4&parseTime=true"
	} else {
		dataSourceName = mq.User + ":" + mq.Password + "@tcp(" + mq.Host + ":" + mq.Port + ")/" + mq.Dbname + "?charset=utf8mb4&parseTime=true"
	}
	engine, err := xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		log.Logger.Panicf(err.Error(), "conf：创建mysql引擎错误")
	}

	// 如果没有这一行，可能会出现 `ID`被映射成了`i_d`，从而导致报错`i_d`列不存在
	engine.SetMapper(names.GonicMapper{})

	if err := engine.Ping(); err != nil {
		log.Logger.Panicf(err.Error(), "conf：mysql连接失败")
	}

	Mqcli = engine
}
