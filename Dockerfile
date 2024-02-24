FROM golang:1.21-buster as build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build \
    -ldflags="-linknode external -extldflags -static" \
    -tags metgo \
    -o bin/api cmd/main.go && ./bin/api

# ....

FROM scratch


COPY --from=build /app/bin/api api

EXPOSE 8080

CMD [ "/api" ]