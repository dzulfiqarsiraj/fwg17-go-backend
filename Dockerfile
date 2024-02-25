FROM golang

WORKDIR /app

COPY . .

RUN go mod tidy

EXPOSE 8090

CMD go run .