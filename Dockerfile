FROM golang:1.22.3-alpine3.19 AS step1
WORKDIR /app
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod go mod download
COPY . .
RUN --mount=type=cache,target=/go/pkg/mod go mod tidy
RUN CGO_ENABLED=0 go build -o api server.go

FROM alpine:3.19
WORKDIR /app
COPY --from=step1 /app/api .
EXPOSE 8080
CMD ["./api"]