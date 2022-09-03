FROM golang:1.19.0-bullseye AS builder

WORKDIR /app

COPY . .

RUN go build -o main .

FROM debian:bullseye AS runtime

WORKDIR /app

COPY --from=builder /app/main .

ENV PORT=5000
ENV HOSTNAME=0.0.0.0
ENV ENVIRONMENT=production

EXPOSE ${PORT}

ENTRYPOINT [ "/app/main" ]