package notifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	urlPkg "net/url"
	"os"
	"strconv"
	"time"

	"github.com/dcasado/alertmanager-notifier/alertmanager"
)

const (
	gotifyURLEnvVariable             = "GOTIFY_URL"
	gotifyTokenEnvVariable           = "GOTIFY_TOKEN"
	gotifyTimeoutMillisEnvVariable   = "GOTIFY_TIMEOUT_MILLIS"
	gotifyDefaultPriorityEnvVariable = "GOTIFY_DEFAULT_PRIORITY"
)

type gotifyClient struct {
	url             string
	token           string
	defaultPriority int

	httpClient http.Client
}

type gotifyMessage struct {
	Title    string `json:"title"`
	Message  string `json:"message"`
	Priority int    `json:"priority"`
}

func newGotifyClient() *gotifyClient {
	urlJoined, _ := urlPkg.JoinPath(getGotifyURLEnvVariable(), "message")
	url, err := urlPkg.ParseRequestURI(urlJoined)
	if err != nil {
		log.Fatalf("new gotify client: %s", err)
	}

	token := getGotifyTokenEnvVariable()
	timeoutMillis := getGotifyTimeoutMillisEnvVariable()
	defaultPriority := getGotifyDefaultPriorityEnvVariable()

	httpClient := http.Client{
		Timeout: time.Duration(timeoutMillis) * time.Millisecond,
	}
	return &gotifyClient{url.String(), token, defaultPriority, httpClient}
}

func (g *gotifyClient) Notify(alert alertmanager.Alert) error {
	title, message, priority := alertmanager.ParseAlert(alert, g.defaultPriority)

	gm := gotifyMessage{Title: title, Message: message, Priority: priority}

	messageBytes, err := json.Marshal(gm)
	if err != nil {
		return fmt.Errorf("could not marshal gotify message: %s", err)
	}

	request, err := http.NewRequest(http.MethodPost, g.url, bytes.NewBuffer(messageBytes))
	if err != nil {
		return fmt.Errorf("error creating request: %s", err)
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-Gotify-Key", g.token)

	resp, err := g.httpClient.Do(request)
	if err != nil {
		return NewErrNotAvailable(g.url, err.Error())
	} else {
		if resp.StatusCode >= http.StatusBadRequest {
			return NewErrHTTPError(resp.StatusCode, http.StatusText(resp.StatusCode))
		}
	}
	return nil
}

func getGotifyURLEnvVariable() string {
	value := os.Getenv(gotifyURLEnvVariable)
	if len(value) != 0 {
		return value
	}
	return "http://localhost:8080"
}

func getGotifyTokenEnvVariable() string {
	gotifyToken := os.Getenv(gotifyTokenEnvVariable)
	if len(gotifyToken) == 0 {
		log.Fatalf("Gotify token is required")
	}
	return gotifyToken
}

func getGotifyTimeoutMillisEnvVariable() int {
	value := os.Getenv(gotifyTimeoutMillisEnvVariable)
	if len(value) != 0 {
		timeout, err := strconv.Atoi(value)
		if err != nil || timeout < 1 {
			log.Fatal("Invalid gotify timeout. Must be a number greater than 0")
		}
		return timeout
	}
	return 5000
}

func getGotifyDefaultPriorityEnvVariable() int {
	value := os.Getenv(gotifyDefaultPriorityEnvVariable)
	if len(value) != 0 {
		defaultPriorityInt, err := strconv.Atoi(value)
		if err != nil {
			log.Fatalln("Invalid default priority value")
		}
		return defaultPriorityInt
	}
	return 5
}
