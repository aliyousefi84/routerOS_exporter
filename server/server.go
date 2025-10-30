package server

import (
	"fmt"
	"net/http"

)

func Init () *Handler {
	return &Handler{}
}

func (h *Handler) RunSrv(addr string , handler http.Handler ) {
	
	http.HandleFunc("GET /checkapi" , h.CheckApi)
	
	//http.HandleFunc("GET /metrics" . )

	err := http.ListenAndServe(addr , handler)

	if err != nil {
		fmt.Println("error to initialize server")
	}

}
