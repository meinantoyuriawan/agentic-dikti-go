package service

import (
	"fmt"

	"github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/prompts"
	"github.com/tmc/langchaingo/tools"
)

func createPrompt() prompts.PromptTemplate {

	content := string(prompt_text)

	prompt := prompts.NewPromptTemplate(
		string(content),
		[]string{},
	)
	return prompt
}

func initAgent(repo Repository) (*agents.OpenAIFunctionsAgent, error) {
	agentTools := []tools.Tool{
		GeneralFAQ{},
		PsychSchedule{repo},
		BookPsychologist{repo},
	}

	llm, err := ollama.New(
		ollama.WithServerURL("http://8.215.1.191:11434/"),
		ollama.WithModel("gpt-oss:20b"),
		ollama.WithKeepAlive("-1m"),
	)
	if err != nil {
		fmt.Println("Error loading ollama")
		return nil, err
	}

	prompt := createPrompt()
	if err != nil {
		fmt.Println("Error loading prompt")
		return nil, err
	}

	agent := agents.NewOpenAIFunctionsAgent(
		llm,
		agentTools,
		agents.WithPrompt(prompt),
	)

	fmt.Println("agent created loh ya")

	return agent, nil
}
