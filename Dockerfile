FROM golang:1.21.3-alpine3.17

WORKDIR /app

COPY src/go.mod .
COPY src/go.sum .
RUN go mod download

CMD ["go", "run", "main.go"]