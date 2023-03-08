package notifier

import (
	"bytes"
	"encoding/base64"
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
	ntfyURLEnvVariable             = "NTFY_URL"
	ntfyTopicEnvVariable           = "NTFY_TOPIC"
	ntfyUserEnvVariable            = "NTFY_USER"
	ntfyPasswordEnvVariable        = "NTFY_PASSWORD"
	ntfyTimeoutMillisEnvVariable   = "NTFY_TIMEOUT_MILLIS"
	ntfyDefaultPriorityEnvVariable = "NTFY_DEFAULT_PRIORITY"
)

type ntfyClient struct {
	url             string
	user            string
	password        string
	defaultPriority int

	httpClient http.Client
}

func newNTFYClient() *ntfyClient {
	urlJoined, _ := urlPkg.JoinPath(getNTFYURLEnvVariable(), getNTFYTopicEnvVariable())
	url, err := urlPkg.ParseRequestURI(urlJoined)
	if err != nil {
		log.Fatalf("new ntfy client: %s", err)
	}

	user := getNTFYUserEnvVariable()
	password := getNTFYPasswordEnvVariable()
	timeoutMillis := getNTFYTimeoutMillisEnvVariable()
	defaultPriority := getNTFYDefaultPriorityEnvVariable()

	httpClient := http.Client{
		Timeout: time.Duration(timeoutMillis) * time.Millisecond,
	}

	return &ntfyClient{url.String(), user, password, defaultPriority, httpClient}
}

func (n *ntfyClient) Notify(alert alertmanager.Alert) error {
	title, message, priority := alertmanager.ParseAlert(alert, n.defaultPriority)

	request, err := http.NewRequest(http.MethodPost, n.url, bytes.NewBufferString(message))
	if err != nil {
		return fmt.Errorf("error creating request: %s", err)
	}
	request.Header.Set("Title", title)
	request.Header.Set("Priority", strconv.Itoa(priority))
	if n.user != "" {
		encodedCredentials := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", n.user, n.password)))
		request.Header.Set("Authorization", fmt.Sprintf("Basic %s", encodedCredentials))
	}

	resp, err := n.httpClient.Do(request)
	if err != nil {
		return NewErrNotAvailable(n.url, err.Error())
	} else {
		if resp.StatusCode >= http.StatusBadRequest {
			return NewErrHTTPError(resp.StatusCode, http.StatusText(resp.StatusCode))
		}
	}
	return nil
}

func getNTFYURLEnvVariable() string {
	value := os.Getenv(ntfyURLEnvVariable)
	if len(value) != 0 {
		return value
	}
	return "http://localhost:8080"
}

func getNTFYTopicEnvVariable() string {
	value := os.Getenv(ntfyTopicEnvVariable)
	if len(value) != 0 {
		return value
	}
	return "alertmanager"
}

func getNTFYUserEnvVariable() string {
	value := os.Getenv(ntfyUserEnvVariable)
	if len(value) != 0 {
		return value
	}
	return ""
}

func getNTFYPasswordEnvVariable() string {
	value := os.Getenv(ntfyPasswordEnvVariable)
	if len(value) != 0 {
		return value
	}
	return ""
}

func getNTFYTimeoutMillisEnvVariable() int {
	defaultTimeoutMillis := 5000
	value := os.Getenv(ntfyTimeoutMillisEnvVariable)
	if len(value) != 0 {
		timeout, err := strconv.Atoi(value)
		if err != nil || timeout < 1 {
			log.Printf("Invalid NTFY timeout. Defaults to: %d", defaultTimeoutMillis)
			return defaultTimeoutMillis
		}
		return timeout
	}
	return defaultTimeoutMillis
}

func getNTFYDefaultPriorityEnvVariable() int {
	defaultPriority := 3
	value := os.Getenv(ntfyDefaultPriorityEnvVariable)
	if len(value) != 0 {
		userDefaultPriorityInt, err := strconv.Atoi(value)
		if err != nil {
			log.Printf("Invalid default priority value. Defaults to: %d", defaultPriority)
			return defaultPriority
		}
		if userDefaultPriorityInt > 0 && userDefaultPriorityInt < 6 {
			return userDefaultPriorityInt
		} else {
			log.Printf("Default priority should be between 1 and 5 both included")
		}
	}
	return defaultPriority
}
