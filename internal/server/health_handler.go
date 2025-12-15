package server

import (
	"net/http"
)

func (s *Server) helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	sendResponse(w, http.StatusOK, "Hello World!")
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	sendResponse(w, http.StatusOK, "OK")
}
