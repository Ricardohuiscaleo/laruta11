FROM golang:1.21-alpine

WORKDIR /app

COPY . .
RUN go mod download
RUN go build -o api main.go

EXPOSE 3001

CMD ["./api"]
