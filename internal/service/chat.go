package service

import (
	"AgenticDikti/internal/constants"
	"AgenticDikti/internal/model"
	"context"
)

func (s *Service) GetChatHistory(ctx context.Context, sessionId string) (chat []model.ChatHistory, err error) {
	chats, err := s.repository.SelectChatBySessionid(ctx, sessionId)
	if err != nil {
		return []model.ChatHistory{}, constants.ErrChatRetrieval
	}

	return chats, err
}

func (s *Service) InputChat(ctx context.Context, userLog model.ChatLogs, aiLog model.ChatLogs) (err error) {
	err = s.repository.InsertChat(ctx, userLog, aiLog)

	if err != nil {
		return constants.ErrChatInput
	}

	return nil
}
