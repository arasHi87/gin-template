.PHONY: init

init:
	go mod download
	go install github.com/swaggo/swag/cmd/swag@v1.8.1

format:
	gofmt -w .

gendoc:
	${HOME}/go/bin/swag init

run: gendoc
	go run main.go

build: gendoc
	go build -o app