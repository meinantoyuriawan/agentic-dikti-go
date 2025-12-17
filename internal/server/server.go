package server

import (
	"AgenticDikti/internal/model"
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
)

type Service interface {
	GetChatHistory(ctx context.Context, sessionId string) (chat []model.ChatHistory, err error)
	InputChat(ctx context.Context, userLog model.ChatLogs, aiLog model.ChatLogs) (err error)
	InputBooking(ctx context.Context, userBookData model.BookingData) (err error)
	RunAgent(ctx context.Context, conversationString, userQuestion string) (string, error)
}

type Server struct {
	port      int
	service   Service
	validator *validator.Validate
}

func NewServer(service Service) *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	v := validator.New()
	NewServer := &Server{
		port:      port,
		service:   service,
		validator: v,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
