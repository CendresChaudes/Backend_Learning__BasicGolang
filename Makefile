GOBIN := $(shell go env GOPATH)/bin
GOLANGCI_LINT := $(GOBIN)/golangci-lint

run: ## Запустить программу
	go run .

pretty: ## Отформатировать код
	$(GOLANGCI_LINT) fmt ./...
	$(GOLANGCI_LINT) run ./...

test: ## Запустить тесты
	go test ./...

help: ## Показать доступные команды
	@echo "Команды:"
	@echo "  make lint   проверить код линтером"
	@echo "  make fmt    отформатировать код"
	@echo "  make run    запустить программу"
