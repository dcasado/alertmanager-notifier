package alertmanager

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type RequestBody struct {
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

func ParseAlert(alert Alert, defaultPriority int) (string, string, int) {
	title := fmt.Sprintf("[%s][%s] ", strings.ToUpper(alert.Status), strings.ToUpper(alert.Labels.Severity))
	if summary := alert.Annotations.Summary; len(summary) != 0 {
		title += summary
	} else {
		log.Printf("Summary annotation not set in alert %s", alert.Labels.Alertname)
	}

	message := ""
	if instance := alert.Labels.Instance; len(instance) != 0 {
		message += fmt.Sprintf("[%s] ", alert.Labels.Instance)
	}

	if description := alert.Annotations.Description; len(description) != 0 {
		message += description
	} else {
		log.Printf("Description annotation not set in alert %s", alert.Labels.Alertname)
	}

	priority := 0
	if priorityValue := alert.Annotations.Priority; len(priorityValue) != 0 {
		p, err := strconv.Atoi(priorityValue)
		if err != nil {
			log.Printf("Priority annotation value not valid in alert %s", alert.Labels.Alertname)
			priority = defaultPriority
		} else {
			priority = p
		}
	} else {
		log.Printf("Priority annotation not set in alert %s", alert.Labels.Alertname)
		priority = defaultPriority
	}

	return title, message, priority
}
