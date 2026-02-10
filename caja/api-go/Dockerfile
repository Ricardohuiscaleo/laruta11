FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod .
RUN go mod download && go mod tidy

COPY . .
RUN go build -o api .

EXPOSE 3002

CMD ["./api"]
