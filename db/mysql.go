package db

import (
	_ "github.com/go-sql-driver/mysql"
	"go-tiktok/conf"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

func InitMq(mq *conf.MysqlConf) (*xorm.Engine, error) {
	dataSourceName := mq.User + ":" + mq.Password + "@tcp(" + mq.Host + ":" + mq.Port + ")/" + mq.Dbname + "?charset=utf8mb4&parseTime=true"
	engine, err := xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		conf.PaincErr(err, "db：创建mysql引擎错误")
	}

	// 如果没有这一行，可能会出现 `ID`被映射成了`i_d`，从而导致报错`i_d`列不存在
	engine.SetMapper(names.GonicMapper{})

	if err := engine.Ping(); err != nil {
		conf.PaincErr(err, "db：mysql连接失败")
	}

	return engine, nil
}
