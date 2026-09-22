.DEFAULT_GOAL := test
.PHONY: pretty test next prev

GOBIN := $(shell go env GOPATH)/bin
GOLANGCI_LINT := $(GOBIN)/golangci-lint

NEXT := ../9.3. Создание инстанса
PREV := ../9.1. Зачем нужны Struct

pretty: ## Отформатировать код
	$(GOLANGCI_LINT) fmt ./...
	$(GOLANGCI_LINT) run ./...

test: ## Запустить тесты
	go test .

ifeq ($(strip $(NEXT)),)
next: ## Перейти в следующий урок
	@echo "Это последний урок."
	@exit 1
else
next: ## Перейти в следующий урок
	cd "$(NEXT)"
endif

ifeq ($(strip $(PREV)),)
prev: ## Перейти в предыдущий урок
	@echo "Это первый урок."
	@exit 1
else
prev: ## Перейти в предыдущий урок
	cd "$(PREV)"
endif
