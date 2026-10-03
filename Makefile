build:
	go build -o bin/helixctl ./cmd/helixctl
	go build -o bin/helixd ./cmd/helixd
	go build -o bin/helix-agent ./cmd/helix-agent

proto:
	@if ! command -v protoc > /dev/null; then \
		echo "protoc not found. Please install protoc and protoc-gen-go."; \
		exit 1; \
	fi
	protoc --go_out=. --go-grpc_out=. api/proto/helix/v1/*.proto

test:
	go test ./...

test-race:
	go test -race ./...

fmt:
	gofmt -w .

vet:
	go vet ./...