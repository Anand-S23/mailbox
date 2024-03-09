build:
	@go build -o bin/mailbox_backend cmd/api/main.go

run: build
	@./bin/mailbox_backend

test:
	@go test -v ./...
