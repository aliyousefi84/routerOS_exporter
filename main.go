package main

import (
	"fmt"

	"github.com/aliyousefi84/routerOS_exporter/internal/prometheus"
	routeros "github.com/aliyousefi84/routerOS_exporter/internal/routerOS"
	"github.com/aliyousefi84/routerOS_exporter/server"
)

func main () {
	prometheus.RegMetrics()
	// pass Mik to Prometheus layer
	Mik , err:= routeros.Initialize()
	if err != nil {
		fmt.Println(err)
	}

	apihandler := server.NewHandler(Mik)

	srv := server.Init(apihandler)

	
	srv.RunSrv("192.168.10.1:9200" , nil)

}

