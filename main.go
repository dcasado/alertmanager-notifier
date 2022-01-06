package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type AlertManagerRequestBody struct {
	Alerts []Alert `json:"alerts"`
}

type Alert struct {
	Status string `json:"status"`
	Labels struct {
		Alertname string `json:"alertname"`
		Instance  string `json:"instance"`
		Severity  string `json:"severity"`
	} `json:"labels"`
	Annotations struct {
		Summary     string `json:"summary"`
		Description string `json:"description"`
		Priority    string `json:"priority"`
	} `json:"annotations"`
}

type GotifyMessage struct {
	Title    string `json:"title"`
	Message  string `json:"message"`
	Priority int    `json:"priority"`
}

const (
	gotifyURLEnvVariable           = "GOTIFY_URL"
	gotifyTokenEnvVariable         = "GOTIFY_TOKEN"
	gotifyTimeoutMillisEnvVariable = "GOTIFY_TIMEOUT_MILLIS"
	listenAddressEnvVariable       = "LISTEN_ADDRESS"
	listenPortEnvVariable          = "LISTEN_PORT"
	defaultPriorityEnvVariable     = "DEFAULT_PRIORITY"

	alertsEndpoint = "/alerts"
	healthEndpoint = "/health"
)

var (
	gotifyURL           string = "http://localhost:8080"
	gotifyToken         string
	gotifyTimeoutMillis int    = 5000
	listenAddress       string = "127.0.0.1"
	listenPort          string = "8080"
	defaultPriority     int    = 5
)

func main() {
	setupEnvVariables()

	log.Printf("Starting server listening on %s:%s", listenAddress, listenPort)

	serveMux := http.NewServeMux()
	serveMux.HandleFunc(alertsEndpoint, handleAlerts)
	serveMux.HandleFunc(healthEndpoint, handleHealth)

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", listenAddress, listenPort),
		Handler: serveMux,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Error starting the server: %s", err)
	}
}

func setupEnvVariables() {
	setupGotifyURLEnvVariable()
	setupGotifyTokenEnvVariable()
	setupGotifyTimeoutMillisEnvVariable()
	setupListenAddressEnvVariable()
	setupListenPortEnvVariable()
	setupDefaultPriorityEnvVariable()
}

func setupGotifyURLEnvVariable() {
	value := os.Getenv(gotifyURLEnvVariable)
	if len(value) != 0 {
		gotifyURL = value
	}
}

func setupGotifyTokenEnvVariable() {
	gotifyToken = os.Getenv(gotifyTokenEnvVariable)
	if len(gotifyToken) == 0 {
		log.Fatalf("Gotify token is required")
	}
}

func setupGotifyTimeoutMillisEnvVariable() {
	value := os.Getenv(gotifyTimeoutMillisEnvVariable)
	if len(value) != 0 {
		timeout, err := strconv.Atoi(value)
		if err != nil || timeout < 1 {
			log.Fatal("Invalid gotify timeout. Must be a number greater than 0")
		}
		gotifyTimeoutMillis = timeout
	}
}

func setupListenAddressEnvVariable() {
	value := os.Getenv(listenAddressEnvVariable)
	if len(value) != 0 {
		listenAddress = value
	}
}

func setupListenPortEnvVariable() {
	value := os.Getenv(listenPortEnvVariable)
	if len(value) != 0 {
		listenPort = value
	}
}

func setupDefaultPriorityEnvVariable() {
	value := os.Getenv(defaultPriorityEnvVariable)
	if len(value) != 0 {
		defaultPriorityInt, err := strconv.Atoi(value)
		if err != nil {
			log.Fatalln("Invalid default priority value")
		}
		defaultPriority = defaultPriorityInt
	}
}

func handleAlerts(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		log.Printf("Method %s not allowed on %s endpoint", request.Method, request.URL)
		http.Error(responseWriter, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	var alertManagerBody AlertManagerRequestBody

	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&alertManagerBody)
	if err != nil {
		log.Println("Malformed request received from alertmanager")
		http.Error(responseWriter, "Malformed request", http.StatusBadRequest)
		return
	}

	for _, alert := range alertManagerBody.Alerts {
		gotifyMessage := buildGotifyMessage(alert)
		message, err := json.Marshal(gotifyMessage)
		if err != nil {
			log.Printf("Could not marshal gotify message: %s", err)
			http.Error(responseWriter, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		client := http.Client{
			Timeout: time.Duration(gotifyTimeoutMillis) * time.Millisecond,
		}

		request, err := http.NewRequest("POST", fmt.Sprintf("%s/message", gotifyURL), bytes.NewBuffer(message))
		if err != nil {
			log.Printf("Error creating new request: %s", err)
			http.Error(responseWriter, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("X-Gotify-Key", gotifyToken)

		resp, err := client.Do(request)
		if err != nil {
			log.Printf("Error sending request to Gotify: %s", err)
			http.Error(responseWriter, fmt.Sprintf("Gotify not available: %s", err), http.StatusGatewayTimeout)
			return
		} else {
			if resp.StatusCode != http.StatusOK {
				log.Printf("Response from gotify returned error. Code: %d, Status: %s", resp.StatusCode, resp.Status)
				http.Error(responseWriter, fmt.Sprintf("Response from gotify returned error: %s", http.StatusText(resp.StatusCode)), http.StatusInternalServerError)
				return
			}
		}
	}
}

func handleHealth(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		log.Printf("Method %s not allowed on %s endpoint", request.Method, request.URL)
		http.Error(responseWriter, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Header().Set("Content-Type", "application/text")
	responseWriter.Write([]byte("Ok"))
}

func buildGotifyMessage(alert Alert) GotifyMessage {
	var gotifyMessage GotifyMessage

	gotifyMessage.Title = fmt.Sprintf("[%s][%s] ", strings.ToUpper(alert.Status), strings.ToUpper(alert.Labels.Severity))
	if summary := alert.Annotations.Summary; len(summary) != 0 {
		gotifyMessage.Title += summary
	} else {
		log.Printf("Summary annotation not set in alert %s", alert.Labels.Alertname)
	}

	if instance := alert.Labels.Instance; len(instance) != 0 {
		gotifyMessage.Message = fmt.Sprintf("[%s] ", alert.Labels.Instance)
	}

	if description := alert.Annotations.Description; len(description) != 0 {
		gotifyMessage.Message += description
	} else {
		log.Printf("Description annotation not set in alert %s", alert.Labels.Alertname)
	}

	if priorityValue := alert.Annotations.Priority; len(priorityValue) != 0 {
		priority, err := strconv.Atoi(priorityValue)
		if err != nil {
			log.Printf("Priority annotation value not valid in alert %s. Defaults to %d", alert.Labels.Alertname, defaultPriority)
			gotifyMessage.Priority = defaultPriority
		} else {
			gotifyMessage.Priority = priority
		}
	} else {
		log.Printf("Priority annotation not set in alert %s. Defaults to %d", alert.Labels.Alertname, defaultPriority)
		gotifyMessage.Priority = defaultPriority
	}

	return gotifyMessage
}
