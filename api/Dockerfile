FROM golang:1.19.0-bullseye AS builder

WORKDIR /app

COPY . .

RUN go build -o main .

FROM debian:bullseye AS runtime

WORKDIR /app

RUN apt-get update && apt-get upgrade --no-install-recommends -y \
    && apt-get install -y curl ca-certificates \
    && apt-get clean && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/main .

ENV PORT=5000
ENV HOSTNAME=0.0.0.0
ENV ENVIRONMENT=production

EXPOSE ${PORT}

ENTRYPOINT [ "/app/main" ]