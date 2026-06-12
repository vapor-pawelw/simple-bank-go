# syntax=docker/dockerfile:1

FROM golang:1.26 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/api/main.go

FROM gcr.io/distroless/base-debian12
COPY --from=builder /app/app /app
EXPOSE 8080
ENTRYPOINT ["/app"]