package service

import (
	"AgenticDikti/internal/constants"
	"AgenticDikti/internal/model"
	"context"
)

// def _generate_chat_id(self):
// unique_string = f"{self.sessionId}-{self.userMessage}-{time.time()}"
// chatId = uuid.uuid5(uuid.NAMESPACE_DNS, unique_string)
// return str(chatId)

// GetChatService
// 1. get history
// 2.

func (s *Service) GetChatHistory(ctx context.Context, sessionId string) (chat []model.ChatHistory, err error) {
	chats, err := s.repository.SelectChatBySessionid(ctx, sessionId)
	if err != nil {
		return []model.ChatHistory{}, constants.ErrChatRetrieval
	}

	return chats, err
}

func (s *Service) InputChat(ctx context.Context, userLog model.ChatLogs, aiLog model.ChatLogs) (chatId string, err error) {
	chatId, err = s.repository.InsertChat(ctx, userLog, aiLog)

	if err != nil {
		return "", constants.ErrChatInput
	}

	return chatId, nil
}
