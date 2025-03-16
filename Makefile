init: install services

install:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.63.4

.PHONY: services

services:
	cd ./todo-ui && npm run format && npm run lint
	cd ./todo-api && $(MAKE) init

build:
	docker compose build

up:
	docker compose up