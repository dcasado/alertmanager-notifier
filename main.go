package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dcasado/alertmanager-notifier/alertmanager"
	"github.com/dcasado/alertmanager-notifier/notifier"
)

const (
	listenAddressEnvVariable = "LISTEN_ADDRESS"
	listenPortEnvVariable    = "LISTEN_PORT"
	notifierTypeEnvVariable  = "NOTIFIER_TYPE"

	alertsEndpoint = "/alerts"
	healthEndpoint = "/health"
)

var s notifier.Notifier

func main() {
	listenAddress := getListenAddressEnvVariable()
	listenPort := getListenPortEnvVariable()
	notifierType := getNotifierTypeEnvVariable()

	s = notifier.New(notifierType)

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

func getListenAddressEnvVariable() string {
	value := os.Getenv(listenAddressEnvVariable)
	if len(value) != 0 {
		return value
	}
	return "127.0.0.1"
}

func getListenPortEnvVariable() string {
	value := os.Getenv(listenPortEnvVariable)
	if len(value) != 0 {
		return value
	}
	return "8080"
}

func getNotifierTypeEnvVariable() string {
	value := os.Getenv(notifierTypeEnvVariable)
	if len(value) != 0 {
		return value
	}
	return notifier.GotifyType
}

func handleAlerts(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		log.Printf("Method %s not allowed on %s endpoint", request.Method, request.URL)
		http.Error(responseWriter, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	var body alertmanager.RequestBody

	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&body)
	if err != nil {
		log.Println("Malformed request received from alertmanager")
		http.Error(responseWriter, "Malformed request", http.StatusBadRequest)
		return
	}

	for _, alert := range body.Alerts {
		err := s.Notify(alert)
		if err != nil {
			log.Printf("Error from notifier: %s", err)
			switch err.(type) {
			case notifier.ErrNotAvailable:
				http.Error(responseWriter, err.Error(), http.StatusGatewayTimeout)
				return
			case notifier.ErrHTTPError:
				http.Error(responseWriter, err.Error(), http.StatusBadGateway)
				return
			default:
				http.Error(responseWriter, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
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
