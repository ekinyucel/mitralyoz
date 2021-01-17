build:
	env GOARCH=amd64 go build -o bin/mitralyoz main.go

run:
	go run main.go