FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.* .

RUN go mod download

# copy folders with content
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/main ./cmd/main.go

# Path: Dockerfile
FROM alpine:latest

WORKDIR /app

# install psql client
RUN apk add --no-cache postgresql-client

COPY --from=builder /app/bin .
COPY entrypoint.sh .

RUN /bin/chmod +x entrypoint.sh

ENTRYPOINT ["/app/entrypoint.sh"]

