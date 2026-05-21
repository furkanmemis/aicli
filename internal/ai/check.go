package ai

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func EnsureModelExists(model string) error {

	// ollama kurulu mu?
	_, err := exec.LookPath("ollama")
	if err != nil {
		return fmt.Errorf("ollama is not installed")
	}

	// ollama serve çalışıyor mu?
	if !isOllamaRunning() {

		fmt.Println("Ollama is not running. Starting ollama serve...")

		err := startOllama()
		if err != nil {
			return err
		}

		// ayağa kalkmasını bekle
		for i := 0; i < 10; i++ {

			if isOllamaRunning() {
				break
			}

			time.Sleep(1 * time.Second)
		}

		if !isOllamaRunning() {
			return fmt.Errorf("failed to start ollama")
		}
	}

	// model var mı?
	listCmd := exec.Command("ollama", "list")

	output, err := listCmd.Output()
	if err != nil {
		return fmt.Errorf("failed to check ollama models: %w", err)
	}

	if strings.Contains(string(output), model) {
		fmt.Printf("Model already exists: %s\n", model)
		return nil
	}

	fmt.Printf("Model not found. Pulling model: %s\n", model)

	pullCmd := exec.Command("ollama", "pull", model)

	pullCmd.Stdout = os.Stdout
	pullCmd.Stderr = os.Stderr

	err = pullCmd.Run()
	if err != nil {
		return fmt.Errorf("failed to pull model: %w", err)
	}

	fmt.Printf("Model downloaded successfully: %s\n", model)

	return nil
}

func isOllamaRunning() bool {

	client := http.Client{
		Timeout: 2 * time.Second,
	}

	resp, err := client.Get("http://localhost:11434/api/tags")
	if err != nil {
		return false
	}

	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK
}

func startOllama() error {

	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {

		cmd = exec.Command("cmd", "/C", "start", "ollama", "serve")

	} else {

		cmd = exec.Command("ollama", "serve")
	}

	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("failed to start ollama serve: %w", err)
	}

	return nil
}