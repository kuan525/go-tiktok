package conf

import (
	"common/logs"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"io"
	"os"
	"time"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

var (
	Cfg    *TiktokConfig
	Mqcli  *xorm.Engine
	Logger *logrus.Logger
)

type Address struct {
	Host string `yaml:"host"` // 开放地址
	Port string `yaml:"port"` // 开放端口
}

type MysqlConf struct {
	Host     string `yaml:"host"`     // 数据库mysql地址
	Port     string `yaml:"port"`     // 数据库开放端口
	User     string `yaml:"user"`     // 数据库用户名
	Password string `yaml:"password"` // 数据库密码
	Dbname   string `yaml:"dbname"`   // 数据库名
}

type LogConfig struct {
	Debug        bool   `yaml:"debug"`        // 是否debug模式,标准输出
	Level        string `yaml:"level"`        // 日志级别 "panic", "fatal", "error", "warning", "info", "debug", "trace"
	SavePath     string `yaml:"savePath"`     // 保存路径
	Suffix       string `yaml:"suffix"`       // 日志后缀 %Y:年,%m:月,%d:日,%H:时,%M:分,%s:秒
	MaxAge       int64  `yaml:"maxAge"`       // 设置文件清理前的最长保存时间 单位秒(s) 默认保存 30天
	RotationTime int64  `yaml:"rotationTime"` // 设置日志分割的时间，隔多久分割一次 单位秒(s) 默认保存 24小时
	TimeFormat   string `yaml:"timeFormat"`   // 时间格式化，默认格式化毫秒 2006-01-02 15:04:05.000000
}

// TiktokConfig 项目所有初始资料
type TiktokConfig struct {
	HttpAddr  Address   `yaml:"httpAddr"`  // http服务地址
	MysqlConf MysqlConf `yaml:"mysqlconf"` // mysql配置
	Log       LogConfig `yaml:"log"`       // 日志配置
}

// InitConfig 初始化配置文件
func InitConfig(path string) error {
	Cfg = new(TiktokConfig) // 初始化一个全局空间
	file, err := os.Open(path)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			logs.HandlePaincErr(err, "conf:关闭Config资源失败")
		}
	}(file)
	if err != nil {
		logs.HandlePaincErr(err, "conf:设置配置文件错误")
	}
	body, err := io.ReadAll(file)
	if err != nil {
		logs.HandlePaincErr(err, "conf:打开配置文件错误")
	}
	if err := yaml.Unmarshal(body, Cfg); err != nil {
		logs.HandlePaincErr(err, "conf:序列化配置文件错误")
	}

	return nil
}

// InitLogger 初始化日志
func InitLogger(c *LogConfig) error {
	savePath := c.SavePath

	writer, _ := rotatelogs.New(
		fmt.Sprintf("%s.%s", savePath, c.Suffix),
		rotatelogs.WithLinkName(savePath),
		rotatelogs.WithMaxAge(time.Duration(c.MaxAge)*time.Second),
		rotatelogs.WithRotationTime(time.Duration(c.RotationTime)*time.Second),
	)

	level, err := logrus.ParseLevel(c.Level)
	if err != nil {
		logs.HandlePaincErr(err, "conf:日志等级配置失败")
	}

	Logger = logrus.New()
	Logger.SetLevel(level)

	if c.Debug {
		Logger.SetOutput(io.MultiWriter(os.Stdout, writer)) // 两者
	} else {
		Logger.SetOutput(writer)
	}

	Logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000000",
	})

	return nil
}

func InitMq(mq *MysqlConf) (*xorm.Engine, error) {
	//dataSourceName := mq.User + ":" + mq.Password + "@tcp(mq.db)/" + mq.Dbname + "?charset=utf8mb4&parseTime=true"
	dataSourceName := mq.User + ":" + mq.Password + "@tcp(" + mq.Host + ":" + mq.Port + ")/" + mq.Dbname + "?charset=utf8mb4&parseTime=true"
	engine, err := xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		logs.HandlePaincErr(err, "conf：创建mysql引擎错误")
	}

	// 如果没有这一行，可能会出现 `ID`被映射成了`i_d`，从而导致报错`i_d`列不存在
	engine.SetMapper(names.GonicMapper{})

	if err := engine.Ping(); err != nil {
		logs.HandlePaincErr(err, "conf：mysql连接失败")
	}

	return engine, nil
}
