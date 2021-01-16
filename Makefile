build:
	env GOARCH=amd64 go build -o bin/main main.go

run:
	go run main.go