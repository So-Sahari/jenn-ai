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
go install
```
