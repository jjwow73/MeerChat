package util

import (
	"io/ioutil"
	"net"
	"net/http"
	"strings"
)

func GetOutboundIP() (string, string) {
	conn, _ := net.Dial("udp", "8.8.8.8:80")
	defer conn.Close()
	localAddr := conn.LocalAddr().String()
	idx := strings.LastIndex(localAddr, ":")
	return localAddr[0:idx], localAddr[idx+1:]
}

//서버에 받을 때 패킷에서 주소:포트 읽어서 저장하게 하기
func GetInboundIP() string {
	url := "https://api.ipify.org?format=text"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(ip)
}