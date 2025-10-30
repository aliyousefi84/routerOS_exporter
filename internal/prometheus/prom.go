package prometheus

import (
	prometheus "github.com/prometheus/client_golang/prometheus"
	//"github.com/prometheus/client_golang/prometheus/promhttp"
) 

type Collect struct {}


func NewCollector () *Collect{
	return &Collect{}
}


var (
	cpu_gauge_load = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "mik_cpu_load",
		Help: "routerOS cpu load",
	})
)

func (c *Collect) SetRouterCpuLoad (value float64) {
	cpu_gauge_load.Set(value)
	
}


func  RegMetrics () {
	prometheus.MustRegister(cpu_gauge_load)
}




