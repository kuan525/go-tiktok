package logs

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

// HandlePaincErr painc错误处理
func HandlePaincErr(err error, msg string) {
	err = fmt.Errorf("%s: %s", err.Error(), msg)
	logrus.Errorf("%s: %s", err.Error(), msg)
	panic(err)
}

// HandleLogsErr 打印错误信息
func HandleLogsErr(err error, msg string) {
	err = fmt.Errorf("%s: %s", err.Error(), msg)
	logrus.Errorf("%s: %s", err.Error(), msg)
}
