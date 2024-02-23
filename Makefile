run:
	go build -o bin/api cmd/main.go && ./bin/api
fmt:
	go fmt *.go

buildwin:
	GOOS=windows GOARCH=amd64 go build -o bin/win/ cmd/main.go