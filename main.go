package main

import (
	"fmt"

	//"github.com/aliyousefi84/routerOS_exporter/config"
	"github.com/aliyousefi84/routerOS_exporter/internal/prometheus"
	routeros "github.com/aliyousefi84/routerOS_exporter/internal/routerOS"
	"github.com/aliyousefi84/routerOS_exporter/server"
)

const (
	addr = "192.168.10.3:8728"
	user = "admin"
	pass = "Ali@1384"
)


func main () {
	prometheus.RegMetrics()
	//Env := config.InitEnv()   *** for getting environment variables from config file ***   
	Mik , err:= routeros.Initialize(addr , user , pass)
	if err != nil {
		fmt.Println(err)
	}

	apihandler := server.NewHandler(Mik)

	srv := server.Init(apihandler)

	
	srv.RunSrv("192.168.10.1:9200" , nil)

}

