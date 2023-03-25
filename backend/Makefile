.DEFAULT_GOAL := run

build:
	go mod download && go build -o ./bin/app ./main.go

run: build
	./bin/app

server-build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o $(OUTPUT) ./main.go
