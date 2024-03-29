ARG GO_VERSION=1
FROM golang:${GO_VERSION}-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

RUN go install github.com/a-h/templ/cmd/templ@latest

COPY . .

RUN go env

RUN ls /go/bin

RUN templ generate

RUN go build -v -o /run-app /app/cmd/bbs

FROM alpine:latest

COPY --from=builder /run-app /usr/local/bin/

CMD ["run-app"]
