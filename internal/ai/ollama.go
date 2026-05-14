package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type OllamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type OllamaResponse struct {
	Response string `json:"response"`
}

func GenerateCommitMessage(diff string) (string, error) {

prompt := `
Generate a git commit message using conventional commits format.

Format examples:
- feat(auth): add jwt refresh token support
- fix(api): handle nil pointer error
- docs(readme): update installation guide

Rules:
- use conventional commits
- lowercase only
- max 72 characters
- no markdown
- no explanation
- output only commit message

Git diff:
` + diff

	body := OllamaRequest{
		Model:  "qwen2.5-coder:7b",
		Prompt: prompt,
		Stream: false,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(
		"POST",
		"http://localhost:11434/api/generate",
		bytes.NewBuffer(jsonBody),
	)

	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	var result OllamaResponse

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", err
	}

	if result.Response == "" {
		return "", fmt.Errorf("empty ai response")
	}

	return result.Response, nil
}