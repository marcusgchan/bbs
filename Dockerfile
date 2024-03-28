ARG GO_VERSION=1
FROM golang:${GO_VERSION}-alpine as builder

# Install system dependencies including 'make'
RUN apk update && apk add --no-cache gcc libc-dev make

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .

RUN make

RUN go build -v -o /run-app /usr/src/app/cmd/bbs

FROM alpine:latest

COPY --from=builder /run-app /usr/local/bin/
CMD ["run-app"]
