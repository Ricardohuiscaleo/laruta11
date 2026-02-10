FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod .
RUN go mod download

COPY . .
RUN go build -o api main.go

EXPOSE 3001

CMD ["./api"]
