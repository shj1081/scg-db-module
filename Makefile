.PHONY: setup
setup:
	go mod tidy
	go install github.com/air-verse/air@latest

.PHONY: swagger
swagger:
	swag init -g ./cmd/server/main.go

.PHONY: clean
clean:
	rm -rf build/ tmp/ docs/

.PHONY: build
build: setup swagger
	go build -o ./build/server ./cmd/server

.PHONY: run
run: setup swagger
	go run ./cmd/server

.PHONY: air
air: setup swagger
	air

.PHONY: dev
dev: air


