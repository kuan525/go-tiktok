package log

import (
	"common/conf"
	"common/file"
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"runtime"
	"time"
)

var Logger *logrus.Logger

// InitLogger 初始化日志
func InitLogger(c *conf.LogConfig) {
	var err error
	formatter := &logrus.TextFormatter{
		ForceColors:               true,
		DisableColors:             false,
		ForceQuote:                true,
		DisableQuote:              false,
		EnvironmentOverrideColors: false,
		DisableTimestamp:          false,
		FullTimestamp:             true,
		TimestampFormat:           "2006-01-02 15:04:05.000",
		DisableSorting:            false,
		SortingFunc:               nil,
		DisableLevelTruncation:    false,
		PadLevelText:              true,
		QuoteEmptyFields:          true,
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "@timestamp",
			logrus.FieldKeyLevel: "@level",
			logrus.FieldKeyMsg:   "@message",
			logrus.FieldKeyFunc:  "@caller",
		},
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			// 处理文件名
			fileName := path.Base(frame.File)
			return frame.Function, fileName
		},
	}
	Logger = logrus.New()
	Logger.SetFormatter(formatter) // 日志格式
	Logger.SetReportCaller(true)   // 输出调用者信息，比如main.main等，以及行号

	savePath := c.SavePath
	if !file.Exists(savePath) {
		// 权限，自己可读写执行，别人可读执行
		err = os.Mkdir(savePath, 0755)
		if err != nil {
			panic("log:创建日志目录失败")
		}
	}

	// 创建一个用于将日志输出到文件的 rotatelogs.Writer
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

	Logger.SetLevel(level)

	if c.Debug {
		Logger.SetOutput(io.MultiWriter(os.Stdout, writer)) // 两者
	} else {
		Logger.SetOutput(writer)
	}

	Logger.Infof("log init success! file path: %s, level: %s", c.SavePath, c.Level)
}
