FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go mod tidy && go build -o api .

EXPOSE 3002

CMD ["./api"]
