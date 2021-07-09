FROM golang:1.16.5-buster

WORKDIR /app

COPY . .

RUN go mod download -v
RUN go build -v main.go

CMD ["main"]