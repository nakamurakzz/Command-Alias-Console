.PHONY: start build help
DEFAULT: help
dev:
	air -c .air.toml

start:
	go run command_alias_console.go

build:
	go build -o bin/ComandAliasConsole command_alias_console.go

test:
	go test -v ./...

lint:
	golangci-lint run

lint-fix:
	golangci-lint run --fix

help:
	@echo "start - run the program. arg: p: profile path, ex) p=~/.zshrc"
	@echo "build - build the program"
	@echo "test - run test"
	@echo "lint - run lint"
	@echo "lint-fix - run lint and fix"
	@echo "help - show this help"