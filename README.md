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

Build:

```bash
go build -o aicli  
```

Commit:

```bash
./aicli commit -a    
```

## Configuration

Task Type & TaskID

```bash
./aicli config -t "fix" -i "TYU-3453"
```

Show All Configuration

```bash
./aicli config -s
```


---

## Tech Stack

- Go
- Cobra CLI
- Lip Gloss
- Ollama
- Qwen2.5 Coder

---


qwen2.5-coder:1.5b was chosen because it provides a good balance between performance, speed, and low hardware requirements for local AI inference. It is lightweight enough to run efficiently on personal machines while still producing high-quality commit message suggestions for developer workflows.



## License

MIT
