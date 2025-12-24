package service

import (
	"AgenticDikti/internal/model"
	"context"

	"github.com/tmc/langchaingo/agents"
)

type Service struct {
	repository Repository
	agent      *agents.OpenAIFunctionsAgent
}

type Repository interface {
	SelectChatBySessionid(ctx context.Context, sessionId string) (res []model.ChatHistory, err error)
	InsertChat(ctx context.Context, userLog model.ChatLogs, aiLog model.ChatLogs) (chatId string, err error)
	InsertBooking(ctx context.Context, userBookData model.BookingData) (err error)
	SelectJadwalPsikolog(ctx context.Context) (res []model.JadwalPsikolog, err error)
}

func New(repository Repository) *Service {
	agent, err := initAgent(repository)
	if err != nil {
		panic(err)
	}
	return &Service{
		repository: repository,
		agent:      agent,
	}
}
