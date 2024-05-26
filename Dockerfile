# Build Stage
FROM golang:1.22.3-alpine3.20 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN  go build -o ./cov-shop ./main.go

# Run Stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/cov-shop .
COPY .env .
EXPOSE 8181
ENTRYPOINT [ "./cov-shop" ]