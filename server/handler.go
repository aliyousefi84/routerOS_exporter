package server

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

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
	ctx := r.Context()
	ctx , cancel := context.WithTimeout(ctx , 3 * time.Second)
	defer cancel()

	 h.Mik.GetCpu(ctx)
	
	 h.Mik.GetFreeMem(ctx)
	
	 h.Mik.GetFreeSpace(ctx)
	 
	 h.Mik.InetTrafikIn(ctx)

	 h.Mik.InetTrafikOut(ctx)

	 h.Mik.UserTrafik(ctx)
	
	promhttp.Handler().ServeHTTP(w , r)
}




