package server

import (
	"encoding/json"
	"net/http"
)


type Handler struct {
	
}




func (h *Handler) CheckApi (w http.ResponseWriter , r *http.Request) {
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "ok",
		"message": "service is up",
		"code": http.StatusOK,
	})
}

