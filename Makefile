.PHONY: test

generate:
	go generate ./...

test:
	go mod tidy
	go test -v ./... -cover