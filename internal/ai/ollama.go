package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

type OllamaRequest struct {
	Model   string                 `json:"model"`
	Prompt  string                 `json:"prompt"`
	Stream  bool                   `json:"stream"`
	Options map[string]interface{} `json:"options,omitempty"`
}

type OllamaResponse struct {
	Response string `json:"response"`
}

func GenerateCommitMessage(
	diff string,
	taskType string,
	taskID string,
) (string, error) {

	prompt := fmt.Sprintf(`
Generate a short git commit description based on this diff.

Rules:
- Output ONLY the description
- No commit type
- No issue id
- No colon
- No markdown
- Keep it short and meaningful

Examples:
add login endpoint
update docker config
fix auth validation

Git diff:
%s
`, diff)

	body := OllamaRequest{
		Model:  "qwen2.5-coder:1.5b",
		Prompt: prompt,
		Stream: false,
		Options: map[string]interface{}{
			"temperature": 0,
			"top_p":       0.1,
			"top_k":       10,
		},
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

	message := strings.TrimSpace(result.Response)

	message = regexp.MustCompile(
		`^(feat|fix|chore|docs|style|refactor|test|perf|build|ci|revert)(\([^)]+\))?:\s*`,
	).ReplaceAllString(message, "")

	message = strings.ReplaceAll(message, "\n", " ")
	message = strings.TrimSpace(message)

	finalCommit := fmt.Sprintf(
		"%s(%s): %s",
		taskType,
		taskID,
		message,
	)

	return finalCommit, nil
}