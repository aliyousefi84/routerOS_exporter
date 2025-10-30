package main

import (
	"fmt"

	routeros "github.com/aliyousefi84/routerOS_exporter/internal/routerOS"
	"github.com/aliyousefi84/routerOS_exporter/server"
)

func main () {
	// pass Mik to Prometheus layer
	if Mik , err:= routeros.Initialize(); err != nil {
		fmt.Println(err)
	}
	
	run := server.Init()
	
	run.RunSrv("localhost:9200" , nil)	

}