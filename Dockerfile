FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

VOLUME /app
CMD ["go", "test", "./...", "-v"]
