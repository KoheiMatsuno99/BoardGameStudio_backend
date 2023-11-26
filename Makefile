fmt:
	go fmt ./...

inspect:
	golangci-lint run ./...

test:
	go test -v ./...

mock:
	go generate ./...

update-proto:
	go get github.com/KoheiMatsuno99/BoardGameStudio_gRPC && go mod tidy
