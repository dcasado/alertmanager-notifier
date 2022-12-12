FROM golang:1.19.4-alpine3.16 AS tester

WORKDIR /app

# Copy files
COPY go.mod ./
COPY alertmanager ./alertmanager
COPY notifier ./notifier
COPY main.go ./

# Run tests
RUN CGO_ENABLED=0 go test -v -timeout 30s


FROM tester AS builder

# Add group and user
RUN addgroup -S alertmanager-notifier && adduser -S alertmanager-notifier -G alertmanager-notifier

# Build binary
RUN CGO_ENABLED=0 go build -o alertmanager-notifier


FROM scratch

ENV LISTEN_ADDRESS="0.0.0.0"

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/alertmanager-notifier /usr/bin/alertmanager-notifier

USER alertmanager-notifier

EXPOSE 8080

ENTRYPOINT ["/usr/bin/alertmanager-notifier"]
