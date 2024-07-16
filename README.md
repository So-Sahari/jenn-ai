# JennAI

JennAI is an open-source chat interface to interact with LLMs.
Part of the inspiration of the project was getting a Go generative AI project running.
After getting a small Go + HTMX + Ollama project setup, I wanted to add more features to it.
JennAI is the generative AI helper that gives both cli and UI access to LLMs.

## Features

Support for models is currently limited to:
- Bedrock Foundational Models
- Ollama Models

## Dependencies
- Bedrock requires model access set as default profile or leveraging `AWS_PROFILE`
- Ollama requires installation [link](https://ollama.com/)

## Installation
```bash
git clone <repo>
go install ./cmd/jenn-ai
```

## Local Development

If you have go installed, ensure that you have air
```bash
make deps
```

Run in development to perform live edits
```bash
make local_dev
```

Run locally
```bash
make local
```

### Docker

Run locally (in development)
```bash
make build_dev
make up_dev
```

Run locally
```bash
make build
make up
```

To bring it down, run
```bash
make down
```

If you have a GPU, then you need to modify the input
```bash
make up COMPOSE_FILE=docker-compose-gpu.yaml
```
Make sure to include the variable when running `down`
