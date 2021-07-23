FROM postgres:13.3-alpine

WORKDIR /var/lib/postgresql

COPY . .

RUN apk add openssl --no-cache

RUN chmod +x self-signed-ssl
RUN ./self-signed-ssl

RUN chown postgres /var/lib/postgresql/server.key && \
    chmod 600 /var/lib/postgresql/server.key