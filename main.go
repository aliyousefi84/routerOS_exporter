package main

import (
	"fmt"

	//"github.com/aliyousefi84/routerOS_exporter/config"
	"github.com/aliyousefi84/routerOS_exporter/config"
	"github.com/aliyousefi84/routerOS_exporter/internal/prometheus"
	routeros "github.com/aliyousefi84/routerOS_exporter/internal/routerOS"
	"github.com/aliyousefi84/routerOS_exporter/server"
)

func main() {
	prometheus.RegMetrics()

	Env := config.InitEnv()
	Mik, err := routeros.Initialize(Env.RoterAddr, Env.RouterUser, Env.RouterPass)
	if err != nil {
		fmt.Println(err)
	}

	apihandler := server.NewHandler(Mik)

	srv := server.Init(apihandler)

	srv.RunSrv(Env.SrvAddr, nil)

}
