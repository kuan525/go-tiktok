package ip

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
)

// GetIp 获取ip
func GetIp(s string) string {
	if s == "internal" {
		return GetInternalIP()
	} else {
		return GetExternalIP()
	}
}

// GetInternalIP 得到内网ip
func GetInternalIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		panic(fmt.Errorf(err.Error(), "拨号失败"))
	}
	defer func(conn net.Conn) {
		err = conn.Close()
		if err != nil {
			panic(fmt.Errorf(err.Error(), "关闭连接失败"))
		}
	}(conn)
	localAddr := conn.LocalAddr().(*net.UDPAddr).String()
	addr, _, err := net.SplitHostPort(localAddr)
	if err != nil {
		panic(fmt.Errorf(err.Error(), "关获得地址失败"))
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
		panic(fmt.Errorf(err.Error(), "拨号失败"))
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			panic(fmt.Errorf(err.Error(), "拨号失败"))
		}
	}(req.Body)

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(fmt.Errorf(err.Error(), "读取body失败"))
	}

	var ip IP
	err = json.Unmarshal(body, &ip)
	if err != nil {
		panic(fmt.Errorf(err.Error(), "反序列化失败"))
	}

	return ip.Query
}
