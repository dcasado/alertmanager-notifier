Alertmanager2gotify is an adapter from [Alertmanager](https://prometheus.io/docs/alerting/latest/alertmanager) webhook requests to [Gotify](https://gotify.net) messages. It transforms your alert manager alerts into Gotify notifications.

# Environment variables

| Name                  | Default value         | Description                                                                      |
|-----------------------|-----------------------|----------------------------------------------------------------------------------|
| GOTIFY_URL            | `http://localhost:8080` | Base Gotify URL                                                                  |
| GOTIFY_TOKEN          |                       | (Required) Token to use on the requests to Gotify                                           |
| GOTIFY_TIMEOUT_MILLIS | `5000`                  | Time limit for requests made to Gotify                                           |
| LISTEN_ADDRESS        | `127.0.0.1`             | Address where the service will listen on                                         |
| LISTEN_PORT           | `8080`                  | Port where the service will listen on                                            |
| DEFAULT_PRIORITY      | `5`                     | Priority to use for Gotify messages when no priority is set on the alert |


# Installation

## Docker

You can use the images form Docker Hub or Github Container Registry
- [davidcasado/alertmanager2gotify](https://hub.docker.com/r/davidcasado/alertmanager2gotify)
- [ghcr.io/dcasado/alertmanager2gotify](https://ghcr.io/dcasado/alertmanager2gotify)

To use with [Docker](https://docker.com) you only need to execute the following command replacing the environment variables with your own:

```bash
docker container run --name alertamanger2gotify -e GOTIFY_URL="http://gotify:8080" -e GOTIFY_TOKEN="token" ghcr.io/dcasado/alertmanager2gotify
```

You can also use it the following [docker-compose](https://docs.docker.com/compose/) file:

```yaml
services:
  alertmanager2gotify:
    image: ghcr.io/dcasado/alertmanager2gotify:latest
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
go build -o alertmanager2gotify
```

### Usage

To use the standalone binary previously built you can execute the following command replacing the environment variables with your own

```bash
GOTIFY_URL="http://127.0.0.1:8080" GOTIFY_TOKEN="token" ./alertmanager2gotify
```
