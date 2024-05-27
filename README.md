alertmanager-notifier is an adapter from [Alertmanager](https://prometheus.io/docs/alerting/latest/alertmanager) webhook requests to [Gotify](https://gotify.net) or [NTFY](https://ntfy.sh/). It transforms your alert manager alerts into notifications.

# Environment variables

| Name                    | Default value           | Description                                                                      |
|-------------------------|-------------------------|----------------------------------------------------------------------------------|
| LISTEN_ADDRESS          | `127.0.0.1`             | Address where the service will listen on                                         |
| LISTEN_PORT             | `8080`                  | Port where the service will listen on                                            |
| NOTIFIER_TYPE           | `gotify`                | Which notifier to use. Valid values are: `gotify` or `ntfy`                      |
| GOTIFY_URL              | `http://localhost:8080` | Base Gotify URL                                                                  |
| GOTIFY_TOKEN            |                         | (Required) Token to use on the requests to Gotify                                |
| GOTIFY_TIMEOUT_MILLIS   | `5000`                  | Time limit for requests made to Gotify                                           |
| GOTIFY_DEFAULT_PRIORITY | `5`                     | Priority to use for Gotify messages when no priority is set on the alert         |
| NTFY_URL                | `http://localhost:8080` | Base NTFY URL                                                                    |
| NTFY_TOPIC              | `alertmanager`          | Topic to which notifications will be sent                                        |
| NTFY_USER               |                         | User to use if authentication is set on NTFY server                              |
| NTFY_PASSWORD           |                         | Password to use if authentication is set on NTFY server                          |
| NTFY_TIMEOUT_MILLIS     | `5000`                  | Time limit for requests made to NTFY                                             |
| NTFY_DEFAULT_PRIORITY   | `3`                     | Priority to use for NTFY notifications when no priority is set on the alert      |


# Installation

## Docker

You can use the images form Docker Hub or Github Container Registry
- [davidcasado/alertmanager-notifier](https://hub.docker.com/r/davidcasado/alertmanager-notifier)
- [ghcr.io/dcasado/alertmanager-notifier](https://ghcr.io/dcasado/alertmanager-notifier)

To use with [Docker](https://docker.com) you only need to execute the following command replacing the environment variables with your own:

```bash
docker container run --name alertamanger2gotify -e GOTIFY_URL="http://gotify:8080" -e GOTIFY_TOKEN="token" ghcr.io/dcasado/alertmanager-notifier
```

You can also use it the following [docker-compose](https://docs.docker.com/compose/) file:

```yaml
services:
  alertmanager-notifier:
    image: ghcr.io/dcasado/alertmanager-notifier:latest
    environment:
      - GOTIFY_URL=http://gotify:8080
      - GOTIFY_TOKEN=token
    restart: unless-stopped
    ports:
      - "8080:8080"
```

## Build from source

To build from source you must have [Go](https://golang.org/) installed on your local machine:

```bash
go build -o alertmanager-notifier
```

### Usage

To use the standalone binary previously built you can execute the following command replacing the environment variables with your own

```bash
GOTIFY_URL="http://127.0.0.1:8080" GOTIFY_TOKEN="token" ./alertmanager-notifier
```
