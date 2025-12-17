package service

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/prompts"
	"github.com/tmc/langchaingo/tools"
)

// Custom server
func (s *Service) RunAgent(ctx context.Context, conversationString, userQuestion string) (string, error) {
	agent, err := initAgent()
	if err != nil {
		return "", err
	}
	messages := createLLMInputMessage(conversationString, userQuestion)
	executor := agents.NewExecutor(agent)
	result, err := executor.Call(ctx, map[string]any{
		"input": messages,
	})

	fmt.Println("Final Output:")
	fmt.Println(result["output"])
	bytes, _ := json.Marshal(result)
	return string(bytes), nil
}

func createLLMInputMessage(conversationString, userQuestion string) string {
	messages := fmt.Sprintf(`
	Conversation History : 
	%s

	Question: User: %s
	`, conversationString, userQuestion)
	return messages
}

func createPrompt() (prompts.PromptTemplate, error) {
	content, err := os.ReadFile("prompt.txt")
	if err != nil {
		return prompts.NewPromptTemplate(``, []string{}), err
	}

	prompt := prompts.NewPromptTemplate(
		string(content),
		[]string{},
	)
	return prompt, nil
}

func initAgent() (*agents.OneShotZeroAgent, error) {
	agentTools := []tools.Tool{
		GeneralFAQ{},
		PsychSchedule{},
		BookPsychologist{},
	}

	llm, err := ollama.New(
		ollama.WithServerURL("http://custom-server:11434"),
		ollama.WithModel("codellama"),
	)
	if err != nil {
		fmt.Println("Error loading ollama")
		return nil, err
	}

	prompt, err := createPrompt()
	if err != nil {
		fmt.Println("Error loading prompt")
		return nil, err
	}

	agent := agents.NewOneShotAgent(
		llm,
		agentTools,
		agents.WithPrompt(prompt),
	)
	return agent, nil
}
