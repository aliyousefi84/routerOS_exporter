package server

import (
	"encoding/json"
	"net/http"

	routeros "github.com/aliyousefi84/routerOS_exporter/internal/routerOS"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)


type Handler struct {
	Mik *routeros.MikSvc	
}

func  NewHandler (Mik *routeros.MikSvc) *Handler{
	return &Handler{
		Mik: Mik,
	} 
}


func (h *Handler) CheckApi (w http.ResponseWriter , r *http.Request) {
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "ok",
		"message": "service is up",
		"code": http.StatusOK,
	})
}

func (h *Handler) PromCheckMetrics (w http.ResponseWriter , r *http.Request) {
	h.Mik.GetCpu()
	promhttp.Handler().ServeHTTP(w , r)
}



