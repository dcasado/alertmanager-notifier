FROM golang:1.17.5-alpine3.15 AS builder

WORKDIR /app

COPY go.mod ./

COPY main.go ./
RUN CGO_ENABLED=0 go build -o alertmanager2gotify


FROM scratch

ENV LISTEN_ADDRESS="0.0.0.0"

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/alertmanager2gotify /usr/bin/alertmanager2gotify

EXPOSE 8080

ENTRYPOINT ["/usr/bin/alertmanager2gotify"]
