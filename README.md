# aicli

AI-powered Git CLI that generates clean and meaningful conventional commit messages using local LLMs with Ollama.

## Features

- AI-generated commit messages
- Conventional commit support
- Local LLM integration with Ollama
- No API key required
- Interactive commit confirmation
- Colored terminal output
- Built with Go and Cobra CLI

---

## Requirements

- Go 1.24+
- Ollama

Install Ollama:

```bash
brew install ollama
```

Start Ollama server:

```bash
ollama serve
```

Download model:

```bash
ollama pull qwen2.5-coder:1.5b
```

List installed models:

```bash
ollama list
```

Remove model:

```bash
ollama rm qwen2.5-coder:1.5b
```

Stop running model:

```bash
ollama stop qwen2.5-coder:1.5b
```

---

## Usage

Stage your changes:

```bash
git add .
```

Generate commit message:

```bash
go run . commit
```

---

## Example Output

```bash
Generated commit message in 0.52s

Suggested commit:
feat(cli): add ai-powered commit generation

Commit? (y/n):
```

---

## Tech Stack

- Go
- Cobra CLI
- Lip Gloss
- Ollama
- Qwen2.5 Coder

---


## License

MIT