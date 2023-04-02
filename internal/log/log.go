package log

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
	"go-tiktok/conf"
	"io"
	"os"
	"time"
)

var Logger *log.Logger

// InitLogger 初始化日志
func InitLogger(c *conf.LogConfig) error {
	savePath := c.SavePath

	writer, _ := rotatelogs.New(
		fmt.Sprintf("%s.%s", savePath, c.Suffix),
		rotatelogs.WithLinkName(savePath),
		rotatelogs.WithMaxAge(time.Duration(c.MaxAge)*time.Second),
		rotatelogs.WithRotationTime(time.Duration(c.RotationTime)*time.Second),
	)

	level, err := log.ParseLevel(c.Level)
	if err != nil {
		conf.PaincErr(err, "conf:日志等级配置失败")
	}

	Logger = log.New()
	Logger.SetLevel(level)

	if c.Debug {
		Logger.SetOutput(io.MultiWriter(os.Stdout, writer)) // 两者
	} else {
		Logger.SetOutput(writer)
	}

	Logger.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000000",
	})

	return nil
}
