package service

import (
	"context"
	"fmt"

	"github.com/tmc/langchaingo/agents"
)

func (s *Service) RunAgent(ctx context.Context, conversationString, userQuestion string) (string, error) {
	messages := createLLMInputMessage(conversationString, userQuestion)
	fmt.Println(messages)
	executor := agents.NewExecutor(s.agent)

	result, err := executor.Call(ctx, map[string]any{
		"input": messages,
	})

	if err != nil {
		return "", err
	}

	// error handling

	fmt.Println(result["output"])
	raw := result["output"].(string)

	return raw, nil
}

func createLLMInputMessage(conversationString, userQuestion string) string {
	messages := fmt.Sprintf(`
	Conversation History : 
	%s

	Question: 
	User: %s
	`, conversationString, userQuestion)
	return messages
}
