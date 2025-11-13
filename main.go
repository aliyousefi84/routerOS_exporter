package main

import (
	"fmt"
	"os"

	"github.com/aliyousefi84/routerOS_exporter/config"
	"github.com/aliyousefi84/routerOS_exporter/internal/prometheus"
	routeros "github.com/aliyousefi84/routerOS_exporter/internal/routerOS"
	"github.com/aliyousefi84/routerOS_exporter/server"
)

func main() {
	prometheus.RegMetrics()

	Env := config.NewConfig()
	Mik, err := routeros.Initialize(Env.RoterAddr, Env.RouterUser, Env.RouterPass)
	if err != nil {
		fmt.Printf("error to initialize routerOS: %v\n", err)
		os.Exit(1)
	}

	apihandler := server.NewHandler(Mik)

	srv := server.Init(apihandler)

	srv.RunSrv(Env.SrvAddr, nil)

}
