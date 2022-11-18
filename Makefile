.PHONY: run
run: # билдить проект с помощью команды make
	go run ./cmd/drug

.DEFAULT_GOAL := run