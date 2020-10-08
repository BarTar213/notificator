.PHONY: install test binary cover cover-total

binary:
	@go build -o notificator ./cmd/notificator

install:
	go install ./...

test:
	@go test ./... -cover

cover:
	@go test -coverprofile=cover.out ./...

cover-total:
	@go test -coverprofile=cover.out ./...
	@go tool cover -func cover.out | grep total: