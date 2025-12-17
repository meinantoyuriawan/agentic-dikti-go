package server

import (
	"AgenticDikti/internal/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func generateDeterministicUUID(sessionId, userMessage string) string {
	data := fmt.Sprintf(`%s-%s-%s`, sessionId, userMessage, time.Now().UTC().Unix())
	return uuid.NewSHA1(uuid.NameSpaceURL, []byte(data)).String()
}

func (s *Server) chatHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
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

	// generate chat id
	// GetChatHistory
	// transform chat history to string format of
	// - User:
	// - Assistant:
	// hit RunAgent(ctx, convstring, req.chatinput)

	// set the return of agent as ai, chatinput as user and save both use chat id
	// call InputChat

	// implementation

	chat, err := s.service.GetChatHistory(ctx, req.SessionId)
	if err != nil {
		log.Printf("error getting chat history: %s\n", err.Error())
		sendErrorResponse(w, http.StatusInternalServerError, "Terjadi kesalahan, Mohon input ulang")
		return
	}

	history := ``

	for _, d := range chat {
		chatRole := `User`
		if d.Role.String == "ai" {
			chatRole = `Assistant`
		}
		history += fmt.Sprintf(`%s : %s \n`, chatRole, d.ChatInput.String)
	}

	agentResponse, err := s.service.RunAgent(ctx, history, req.ChatInput)
	if err != nil {
		log.Printf("error hit agent: %s\n", err.Error())
		sendErrorResponse(w, http.StatusInternalServerError, "Terjadi kesalahan, Mohon input ulang")
		return
	}

	userChat := model.ChatLogs{
		// req.SessionId,
		// generateDeterministicUUID(req.SessionId, req.ChatInput),
		// req.ChatInput,
		// "timestamp",
		// "user",
		// false,
		// req.UniversityId,
	}
	aiChat := model.ChatLogs{
		// req.SessionId,
		// generateDeterministicUUID(req.SessionId, agentResponse),
		// agentResponse,
		// "timestamp",
		// "ai",
		// false,
		// req.UniversityId,
	}

	err = s.service.InputChat(ctx, userChat, aiChat)
	if err != nil {
		log.Printf("error saving chat: %s\n", err.Error())
		sendErrorResponse(w, http.StatusInternalServerError, "Terjadi kesalahan, Mohon input ulang")
		return
	}

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
