# This Makefile is used to easy the commands used

deps:
	@go install github.com/air-verse/air@latest
	@go get

tidy:
	@go mod tidy

local_dev:
	@air

local:
	@go install ./cmd/jennai
	@jenn-ai

tests:
	@go test -cover ./...

# default value is set to docker-compose.yaml
# if you have a gpu, you will want to run:
#
# make <command> COMPOSE_FILE=docker-compose-gpu.yaml
COMPOSE_FILE ?= docker-compose.yaml

build_dev: $(COMPOSE_FILE)
	@docker-compose -f $< build jenn_ai_dev

up_dev: $(COMPOSE_FILE)
	@docker-compose -f $< up -d jenn_ai_dev

build: $(COMPOSE_FILE)
	@docker-compose -f $< build jenn_ai

up: $(COMPOSE_FILE)
	@docker-compose -f $< up -d jenn_ai

down: $(COMPOSE_FILE)
	@docker-compose -f $< down

# exec is used to enter the ollama container to install models
exec:
	@docker exec -it ollama bash
