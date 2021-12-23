FROM golang:1.17.5-alpine3.15 AS builder

# Add group and user
RUN addgroup -S alertmanager2gotify && adduser -S alertmanager2gotify -G alertmanager2gotify

WORKDIR /app

COPY go.mod ./

COPY main.go ./
RUN CGO_ENABLED=0 go build -o alertmanager2gotify


FROM scratch

ENV LISTEN_ADDRESS="0.0.0.0"

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/alertmanager2gotify /usr/bin/alertmanager2gotify

USER alertmanager2gotify

EXPOSE 8080

ENTRYPOINT ["/usr/bin/alertmanager2gotify"]
