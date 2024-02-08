run:
	go build -o bin/api cmd/main.go && ./bin/api
fmt:
	go fmt *.go