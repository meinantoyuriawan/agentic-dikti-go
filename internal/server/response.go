package server

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error,omitempty"`
}

func sendResponse(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if body == nil {
		return
	}

	if err := json.NewEncoder(w).Encode(body); err != nil {
		log.Printf("failed to encode response: %v", err)
	}
}

func sendErrorResponse(w http.ResponseWriter, statusCode int, error string) {
	resp := ErrorResponse{
		Error: error,
	}

	sendResponse(w, statusCode, resp)
}

type DiktiResponse struct {
	SessionId    string `json:"sessionid"`
	ChatId       string `json:"chatid"`
	ChatInput    string `json:"chatinput"`
	Timestamp    string `json:"timestamp"`
	Role         string `json:"role"`
	Emergency    bool   `json:"emergency"`
	UniversityId string `json:"universityid"`
}

//     def __init__(self, sessionid, chatid, chatinput, timestamp, role, emergency, universityid):
// self.sessionid = sessionid
// self.chatid = chatid
// self.chatinput = chatinput
// self.timestamp = timestamp
// self.role = role
// self.emergency = emergency
// self.universityid = universityid
