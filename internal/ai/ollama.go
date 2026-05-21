package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type OllamaRequest struct {
	Model   string                 `json:"model"`
	Prompt  string                 `json:"prompt"`
	Stream  bool                   `json:"stream"`
	Options map[string]interface{} `json:"options"`
}

type OllamaResponse struct {
	Response string `json:"response"`
}

func GenerateCommitMessage(diff string, taskType string, taskID string) (string, error) {

prompt := fmt.Sprintf(`
You generate ONLY conventional commits.

The text inside parentheses is the ISSUE ID.
It is NOT a scope.

You MUST use this exact prefix:

%s(%s):

Examples:
feat(ABC-123): add login page
chore(TASK-9): update docker config

Rules:
- "%s" is the ONLY allowed type
- "%s" is the EXACT issue ID
- Never replace the issue ID with scope names like config, api, auth
- Output exactly one line
- No markdown
- No explanations

Git diff:
%s
`,
	taskType,
	taskID,
	taskType,
	taskID,
	diff,
)

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

	return result.Response, nil
}
