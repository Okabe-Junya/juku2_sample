.PHONY: build vet

DEFAULT_GOAL := build

build: vet
	@echo $(BINARY_NAME)
	go build -o bin/$(BINARY_NAME) -v

vet:
	go vet ./...
