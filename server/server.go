package server

import (
	"fmt"
	"net/http"
)


type Server struct{
	handler *Handler

}

func Init (handler *Handler) *Server {
	return &Server{
		handler: handler,
	}
}


func (s *Server) RunSrv(addr string , handler http.Handler ) {
	
	
	http.HandleFunc("GET /checkapi" , s.handler.CheckApi)

	http.HandleFunc("GET /metrics" , s.handler.PromCheckMetrics)

	err := http.ListenAndServe(addr , handler)

	if err != nil {
		fmt.Println("error to initialize server")
	}

}
