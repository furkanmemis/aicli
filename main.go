package main

import (
	"github.com/furkanmemis/aicli/cmd"
	"github.com/furkanmemis/aicli/internal/ai"
)

func main() {
	ai.EnsureModelExists("qwen2.5-coder:1.5b")
	cmd.Execute()
}