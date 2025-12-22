package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

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
	fmt.Println(messages)
	executor := agents.NewExecutor(agent)

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

type AgentOutput struct {
	Response  string `json:"response"`
	Emergency bool   `json:"emergency"`
}

func ParseAgentFinalAnswer(raw string) (*AgentOutput, error) {
	if !strings.HasPrefix(raw, "{") || !strings.HasSuffix(raw, "}") {
		return nil, errors.New("final answer does not contain valid JSON object")
	}

	// 4. Unmarshal
	var out AgentOutput
	if err := json.Unmarshal([]byte(raw), &out); err != nil {
		return nil, err
	}

	// 5. Validate required fields
	if out.Response == "" {
		return nil, errors.New("missing or empty 'response' field")
	}

	return &out, nil
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

func createPrompt() (prompts.PromptTemplate, error) {
	content, err := os.ReadFile("internal/service/prompt.txt")
	if err != nil {
		return prompts.NewPromptTemplate(``, []string{}), err
	}

	prompt := prompts.NewPromptTemplate(
		string(content),
		[]string{},
	)
	return prompt, nil
}

func initAgent() (*agents.OpenAIFunctionsAgent, error) {
	agentTools := []tools.Tool{
		GeneralFAQ{},
		PsychSchedule{},
		BookPsychologist{},
	}

	llm, err := ollama.New(
		ollama.WithServerURL("http://8.215.1.191:11434/"),
		ollama.WithModel("gpt-oss:20b"),
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

	agent := agents.NewOpenAIFunctionsAgent(
		llm,
		agentTools,
		agents.WithPrompt(prompt),
	)

	// agent := agents.NewOneShotAgent(
	// 	llm,
	// 	agentTools,
	// 	agents.WithPrompt(prompt),
	// 	agents.WithMaxIterations(3),
	// )
	return agent, nil
}
