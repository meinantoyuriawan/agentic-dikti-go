package service

import (
	"AgenticDikti/internal/model"
	"context"
)

type Service struct {
	repository Repository
}

type Repository interface {
	SelectChatBySessionid(ctx context.Context, sessionId string) (res []model.ChatHistory, err error)
	InsertChat(ctx context.Context, userLog model.ChatLogs, aiLog model.ChatLogs) (err error)
	InsertBooking(ctx context.Context, userBookData model.BookingData) (err error)
}

func New(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}
