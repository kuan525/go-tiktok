package conf

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"time"

	_ "github.com/go-sql-driver/mysql"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
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
			panic(fmt.Errorf(err.Error(), "conf:关闭Config资源失败"))
		}
	}(file)
	if err != nil {
		panic(fmt.Errorf(err.Error(), "conf:设置配置文件错误"))
	}
	body, err := io.ReadAll(file)
	if err != nil {
		panic(fmt.Errorf(err.Error(), "conf:打开配置文件错误"))
	}
	if err := yaml.Unmarshal(body, Cfg); err != nil {
		panic(fmt.Errorf(err.Error(), "conf:序列化配置文件错误"))
	}

	return nil
}

// InitLogger 初始化日志
func InitLogger(c *LogConfig) error {
	savePath := c.SavePath

	// 创建一个用于将日志输出到文件的 rotatelogs.Writer
	/*
			关于软连接的问题：
			1. 第一个参数是我当前创建的文件名
			2. 第二个参数是只想我当前文件名的软链接，目录以及名称
			3. 当软链接名存在的时候，会删除之前那个软链接，并新绑定上（和当前的日志文件）
			gpt：当日志文件存在时，WithLinkName 方法只会创建软链接，而不会删除任何文件。
		         软链接会覆盖同名的已有软链接，所以如果之前已经有同名的软链接，它将被替换为新的软链接。
		 		 如果没有同名的软链接，将会创建一个新的软链接。
	*/
	writer, err := rotatelogs.New(
		// 日志文件的路径和文件名后缀
		fmt.Sprintf("%s%s.%s", savePath, time.Now().Format("2006-01"), c.Suffix),
		// 用于在创建软链接时指定软链接的名称
		rotatelogs.WithLinkName(savePath+"/current_log"),
		// 日志文件的最大保存时间，time.Duration是一个时间段类型，单位是ns，后面也是ns
		rotatelogs.WithMaxAge(time.Duration(c.MaxAge)*time.Second),
		// 日志轮转的时间
		rotatelogs.WithRotationTime(time.Duration(c.RotationTime)*time.Second),
	)
	defer writer.Close()
	if err != nil {
		panic(fmt.Errorf(err.Error(), "日志文件创建出错"))
	}

	// info：用于记录程序运行的一些重要信息，例如请求的处理时间、用户行为等等。
	level, err := logrus.ParseLevel(c.Level)
	if err != nil {
		panic(fmt.Errorf("%s:%s", err.Error(), "conf:日志等级配置失败"))
	}

	Logger = logrus.New()
	Logger.SetLevel(level)

	if c.Debug {
		Logger.SetOutput(io.MultiWriter(os.Stdout, writer)) // 两者
	} else {
		Logger.SetOutput(writer)
	}

	// 将日志的格式设置为 JSON 格式，并指定时间戳的格式
	Logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000000",
	})

	return nil
}

func InitMq(mq *MysqlConf) (*xorm.Engine, error) {
	var dataSourceName string
	if runtime.GOOS == "darwin" && runtime.GOARCH == "arm64" { // 本机
		dataSourceName = "root:mysql_112525@tcp(localhost:3306)/tiktok?charset=utf8mb4&parseTime=true"
	} else {
		dataSourceName = mq.User + ":" + mq.Password + "@tcp(" + mq.Host + ":" + mq.Port + ")/" + mq.Dbname + "?charset=utf8mb4&parseTime=true"
	}
	engine, err := xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		Logger.Infof(err.Error(), "conf：创建mysql引擎错误")
	}

	// 如果没有这一行，可能会出现 `ID`被映射成了`i_d`，从而导致报错`i_d`列不存在
	engine.SetMapper(names.GonicMapper{})

	if err := engine.Ping(); err != nil {
		Logger.Infof(err.Error(), "conf：mysql连接失败")
	}
	//engine.Sync2(models.User{})
	return engine, nil
}
