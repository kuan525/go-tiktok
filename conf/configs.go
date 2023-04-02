package conf

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"io"
	"os"
	"xorm.io/xorm"
)

var (
	Cfg   *TiktokConfig
	Mqcli *xorm.Engine
)

type Address struct {
	Host string `yaml:"host"` // 开放地址
	Port int    `yaml:"port"` // 开放端口
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

// PaincErr painc错误处理
func PaincErr(err error, msg string) {
	err = fmt.Errorf("%s: %s", err.Error(), msg)
	panic(err)
}

// LogErr 打印错误信息
func LogErr(err error, msg string) {
	err = fmt.Errorf("%s: %s", err.Error(), msg)
	log.Errorf("%s: %s", err.Error(), msg)
}

// InitConfig 初始化配置文件
func InitConfig(path string) error {
	Cfg = new(TiktokConfig) // 初始化一个全局空间
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		PaincErr(err, "conf:设置配置文件错误")
	}
	body, err := io.ReadAll(file)
	if err != nil {
		PaincErr(err, "conf:打开配置文件错误")
	}
	if err := yaml.Unmarshal(body, Cfg); err != nil {
		PaincErr(err, "conf:序列化配置文件错误")
	}

	return nil
}
