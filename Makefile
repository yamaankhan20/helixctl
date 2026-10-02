build:
	go build -o bin/helixctl ./cmd/helixctl
	go build -o bin/helixd ./cmd/helixd
	go build -o bin/helix-agent ./cmd/helix-agent

test:
	go test ./...

test-race:
	go test -race ./...

fmt:
	gofmt -w .

vet:
	go vet ./...
