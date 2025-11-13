package prometheus

import (
	prometheus "github.com/prometheus/client_golang/prometheus"
)

type Collect struct{}

func NewCollector() *Collect {
	return &Collect{}
}

var (
	cpu_gauge_load = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "mik_cpu_load",
		Help: "routerOS cpu load",
	})

	mik_mem_free = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "mik_mem_free",
		Help: "routerOS free memory",
	})

	mik_hard_free = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "mik_hard_free",
		Help: "routerOS free space",
	})

	mik_in_trafik = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "mik_in_trafik",
		Help: "mikrotik input trafik for interfaces",
	},
		[]string{"interface_name", "direction"},
	)

	mik_out_trafik = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "mik_out_trafik",
		Help: "mikrotik out trafik from interfaces",
	},
		[]string{"interface_name", "direction"},
	)

	mik_user_trafik = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "mik_user_trafik",
		Help: "mikrotik user trafik",
	},
		[]string{"src_address"},
	)
)

func (c *Collect) SetRouterCpuLoad(value float64) {
	cpu_gauge_load.Set(value)

}

func (c *Collect) SetRouterFreeMem(value float64) {
	mik_mem_free.Set(value)
}

func (c *Collect) SetHardFreeSpace(value float64) {
	mik_hard_free.Set(value)
}

func (c *Collect) GetTrafikIn(name, direction string, rx float64) {
	mik_in_trafik.WithLabelValues(name, direction).Add(rx)
}

func (c *Collect) GetTrafikOut(name, direction string, tx float64) {
	mik_out_trafik.WithLabelValues(name, direction).Add(tx)
}

func (c *Collect) GetUserTrafik(source string, trafik float64) {
	mik_user_trafik.WithLabelValues(source).Add(trafik)
}

func RegMetrics() {
	prometheus.MustRegister(
		cpu_gauge_load,
		mik_mem_free,
		mik_hard_free,
		mik_in_trafik,
		mik_out_trafik,
		mik_user_trafik,
	)
}
