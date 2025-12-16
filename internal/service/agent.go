package service

import (
	"context"
	"fmt"

	"github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/prompts"
	"github.com/tmc/langchaingo/tools"
)

// Custom server

func run(ctx context.Context) error {

	llm, err := ollama.New(
		ollama.WithServerURL("http://custom-server:11434"),
		ollama.WithModel("codellama"),
	)
	if err != nil {
		return err
	}
	prompt := prompts.NewPromptTemplate(
		`You are a helpful assistant.
		You may use tools if needed.
		
		Question: {{.input}}`,
		[]string{"input"},
	)

	agentTools := []tools.Tool{
		GeneralFAQ{},
		PsychSchedule{},
	}

	agent := agents.NewOneShotAgent(
		llm,
		agentTools,
		agents.WithPrompt(prompt),
	)
	executor := agents.NewExecutor(agent)

	result, err := executor.Call(ctx, map[string]any{
		"input": "What is 12 plus 30?",
	})

	if err != nil {
		return err
	}

	fmt.Println("Final Output:")
	fmt.Println(result["output"])
	return err
}
