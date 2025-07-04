ARG GO_VERSION=1.24.1

# fetch
FROM golang:${GO_VERSION} AS fetch-stage
COPY go.mod go.sum /app
WORKDIR /app
RUN go mod download && go mod verify

# generate
FROM ghcr.io/a-h/templ:v0.3.833 AS generate-stage
COPY --chown=65532:65532 . /app
WORKDIR /app
RUN ["templ", "generate"]

# npm
FROM node:22 AS npm-stage
WORKDIR /app
COPY . .
RUN npm install
RUN npm run build

# build
FROM golang:${GO_VERSION} AS build-stage
COPY --from=generate-stage /app /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -v -o /app/server

# release
FROM debian:bookworm-slim AS release
COPY --from=build-stage /app /app
COPY --from=npm-stage /app/assets/public /app/assets/public
WORKDIR /app

CMD ["/app/server"]
