test:
	go test -race -coverprofile="coverage.out" -covermode=atomic ./...
	go tool cover -html="coverage.out"

lint:
	golangci-lint run

run:
	go run ./cmd/healths_app/main.go