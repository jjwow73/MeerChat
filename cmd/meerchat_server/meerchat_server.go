package main

import (
	"github.com/wkd3475/MeerChat/conf"
	"github.com/wkd3475/MeerChat/pkg/meerchat_server"
	"strconv"
)

func main() {
	port, _ := strconv.Atoi(conf.MainServerPort)
	meerchat_server.MakeServer(port)
}
