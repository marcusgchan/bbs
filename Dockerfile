ARG GO_VERSION=1

FROM node:20-slim AS base
WORKDIR /app
COPY . /app
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable

FROM base AS build
RUN --mount=type=cache,id=s/86f41c16-56f7-4121-b03b-097e276cf191-pnpm,target=/pnpm/store pnpm install --frozen-lockfile
RUN pnpm run build


FROM golang:${GO_VERSION}-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

RUN go install github.com/a-h/templ/cmd/templ@latest

COPY . .

COPY --from=build /app/build /app/build

RUN templ generate

RUN go build -v -o /run-app /app/cmd/bbs


FROM alpine:latest

COPY --from=builder /run-app /usr/local/bin/

CMD ["run-app"]
