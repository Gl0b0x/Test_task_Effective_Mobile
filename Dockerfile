# Step 1: Modules caching
FROM golang:1.23.1-alpine as modules
COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

# Step 2: Builder
FROM golang:1.23.1-alpine as builder
COPY --from=modules /go/pkg /go/pkg
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /bin/app ./cmd

FROM alpine:latest
COPY --from=builder /app/config /config
COPY --from=builder /bin/app /app

CMD ["./app"]