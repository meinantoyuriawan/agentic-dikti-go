package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func (s *Server) chatHandler(w http.ResponseWriter, r *http.Request) {
	// ctx := r.Context()
	var req DiktiRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("invalid chat request")
		sendErrorResponse(w, http.StatusBadRequest, "Terjadi kesalahan, Mohon input ulang")
		return
	}

	err = s.validator.Struct(req)
	if err != nil {
		log.Printf("invalid chat request: %s\n", err.Error())
		sendErrorResponse(w, http.StatusBadRequest, "Terjadi kesalahan, Mohon input ulang")
		return
	}

	// chat, err := s.service.GetChatHistory(ctx, req.SessionId)
	// if err != nil {
	// 	log.Printf("error getting chat history: %s\n", err.Error())
	// 	sendErrorResponse(w, http.StatusInternalServerError, "Terjadi kesalahan, Mohon input ulang")
	// 	return
	// }

	//agent calling
	resp := DiktiResponse{
		SessionId:    req.SessionId,
		ChatId:       "chatIDFromService",
		ChatInput:    "chatInputFromAgent",
		Timestamp:    "timestamp from service layer",
		Role:         "ai",
		Emergency:    false,
		UniversityId: req.UniversityId,
	}
	sendResponse(w, http.StatusOK, resp)

}
