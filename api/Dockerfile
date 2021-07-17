FROM golang:1.16.6-buster

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -v main.go

EXPOSE ${PORT}

CMD ["./main"]