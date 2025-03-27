.PHONY: setup
setup:
	go mod tidy
	go install github.com/air-verse/air@latest

.PHONY: clean
clean:
	rm -rf build/ tmp/

.PHONY: build
build: setup
	go build -o ./build/server ./cmd/server

.PHONY: run
run: setup
	go run ./cmd/server

.PHONY: air
air: setup
	air

.PHONY: dev
dev: air

