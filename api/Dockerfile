FROM golang:1.16.5-buster

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -v main.go

CMD ["./main"]