run:
	go build -o bin/api cmd/codegram/main.go && ./bin/api
fmt:
	go fmt *.go