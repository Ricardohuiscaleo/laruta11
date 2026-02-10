FROM golang:1.21-alpine

WORKDIR /app

ENV GOSUMDB=off

COPY go.mod .
RUN go mod download

COPY . .
RUN GOSUMDB=off go build -o api main.go

EXPOSE 3001

CMD ["./api"]
