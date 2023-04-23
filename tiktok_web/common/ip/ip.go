package ip

import (
	"common/log"
	"encoding/json"
	"io"
	"io/ioutil"
	"net"
	"net/http"
)

// GetInternalIP 得到内网ip
func GetInternalIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Logger.Errorf(err.Error(), "拨号失败")
	}
	defer func(conn net.Conn) {
		err = conn.Close()
		log.Logger.Errorf(err.Error(), "关闭连接失败")
	}(conn)
	localAddr := conn.LocalAddr().(*net.UDPAddr).String()
	addr, _, err := net.SplitHostPort(localAddr)
	if err != nil {
		log.Logger.Errorf(err.Error(), "获得地址失败")
	}
	return addr
}

// GetExternalIP 使用API获取外网IP
func GetExternalIP() string {
	type IP struct {
		Query string
	}
	req, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		log.Logger.Errorf(err.Error(), "拨号失败")
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			log.Logger.Errorf(err.Error(), "拨号失败")
		}
	}(req.Body)

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Logger.Errorf(err.Error(), "读取body失败")
	}

	var ip IP
	err = json.Unmarshal(body, &ip)
	if err != nil {
		log.Logger.Errorf(err.Error(), "反序列化失败")
	}

	return ip.Query
}
