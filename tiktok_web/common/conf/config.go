package conf

import (
	"common/ip"
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
	"os"
)

// InitConfig 初始化配置文件
func InitConfig(path string) {
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

	// 获取内网或者外网ip
	if Cfg.HttpAddr.NetEnv == "internal" {
		Cfg.HttpAddr.Host = ip.GetInternalIP()
	} else {
		Cfg.HttpAddr.Host = ip.GetExternalIP()
	}
}
