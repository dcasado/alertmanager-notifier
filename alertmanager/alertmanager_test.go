package alertmanager

import (
	"testing"
)

func Test_parseAlert(t *testing.T) {
	expectedTitle := "[FIRING][WARNING] Summary"
	expectedMessage := "[instance] Description"
	expectedPriority := 2

	var alert Alert
	alert.Status = "firing"
	alert.Labels.Alertname = "Test alert"
	alert.Labels.Instance = "instance"
	alert.Labels.Severity = "warning"
	alert.Annotations.Summary = "Summary"
	alert.Annotations.Description = "Description"
	alert.Annotations.Priority = "2"
	actualTitle, actualMessage, actualPriority := ParseAlert(alert, 5)

	if expectedTitle != actualTitle {
		t.Errorf("Title was incorrect want: \"%+v\", but got: \"%+v\"", expectedTitle, actualTitle)
	}
	if expectedMessage != actualMessage {
		t.Errorf("Message was incorrect want: %+v, but got: %+v", expectedMessage, actualMessage)
	}
	if expectedPriority != actualPriority {
		t.Errorf("Priority was incorrect want: %+v, but got: %+v", expectedPriority, actualPriority)
	}
}

func Test_parseAlert_noInstance(t *testing.T) {
	expectedTitle := "[FIRING][WARNING] Summary"
	expectedMessage := "Description"
	expectedPriority := 2

	var alert Alert
	alert.Status = "firing"
	alert.Labels.Alertname = "Test alert"
	alert.Labels.Severity = "warning"
	alert.Annotations.Summary = "Summary"
	alert.Annotations.Description = "Description"
	alert.Annotations.Priority = "2"
	actualTitle, actualMessage, actualPriority := ParseAlert(alert, 5)

	if expectedTitle != actualTitle {
		t.Errorf("Title was incorrect want: \"%+v\", but got: \"%+v\"", expectedTitle, actualTitle)
	}
	if expectedMessage != actualMessage {
		t.Errorf("Message was incorrect want: %+v, but got: %+v", expectedMessage, actualMessage)
	}
	if expectedPriority != actualPriority {
		t.Errorf("Priority was incorrect want: %+v, but got: %+v", expectedPriority, actualPriority)
	}
}

func Test_parseAlarm_noSummary(t *testing.T) {
	expectedTitle := "[FIRING][WARNING] "
	expectedMessage := "[instance] Description"
	expectedPriority := 2

	var alert Alert
	alert.Status = "firing"
	alert.Labels.Alertname = "Test alert"
	alert.Labels.Instance = "instance"
	alert.Labels.Severity = "warning"
	alert.Annotations.Description = "Description"
	alert.Annotations.Priority = "2"
	actualTitle, actualMessage, actualPriority := ParseAlert(alert, 5)

	if expectedTitle != actualTitle {
		t.Errorf("Title was incorrect want: \"%+v\", but got: \"%+v\"", expectedTitle, actualTitle)
	}
	if expectedMessage != actualMessage {
		t.Errorf("Message was incorrect want: %+v, but got: %+v", expectedMessage, actualMessage)
	}
	if expectedPriority != actualPriority {
		t.Errorf("Priority was incorrect want: %+v, but got: %+v", expectedPriority, actualPriority)
	}
}

func Test_parseAlert_noDescription(t *testing.T) {
	expectedTitle := "[FIRING][WARNING] Summary"
	expectedMessage := "[instance] "
	expectedPriority := 2

	var alert Alert
	alert.Status = "firing"
	alert.Labels.Alertname = "Test alert"
	alert.Labels.Instance = "instance"
	alert.Labels.Severity = "warning"
	alert.Annotations.Summary = "Summary"
	alert.Annotations.Priority = "2"
	actualTitle, actualMessage, actualPriority := ParseAlert(alert, 5)

	if expectedTitle != actualTitle {
		t.Errorf("Title was incorrect want: \"%+v\", but got: \"%+v\"", expectedTitle, actualTitle)
	}
	if expectedMessage != actualMessage {
		t.Errorf("Message was incorrect want: %+v, but got: %+v", expectedMessage, actualMessage)
	}
	if expectedPriority != actualPriority {
		t.Errorf("Priority was incorrect want: %+v, but got: %+v", expectedPriority, actualPriority)
	}
}

func Test_parseAlert_noPriority(t *testing.T) {
	expectedTitle := "[FIRING][WARNING] Summary"
	expectedMessage := "[instance] Description"
	expectedPriority := 5

	var alert Alert
	alert.Status = "firing"
	alert.Labels.Alertname = "Test alert"
	alert.Labels.Instance = "instance"
	alert.Labels.Severity = "warning"
	alert.Annotations.Summary = "Summary"
	alert.Annotations.Description = "Description"
	actualTitle, actualMessage, actualPriority := ParseAlert(alert, 5)

	if expectedTitle != actualTitle {
		t.Errorf("Title was incorrect want: \"%+v\", but got: \"%+v\"", expectedTitle, actualTitle)
	}
	if expectedMessage != actualMessage {
		t.Errorf("Message was incorrect want: %+v, but got: %+v", expectedMessage, actualMessage)
	}
	if expectedPriority != actualPriority {
		t.Errorf("Priority was incorrect want: %+v, but got: %+v", expectedPriority, actualPriority)
	}
}

func Test_parseAlert_wrongPriority(t *testing.T) {
	expectedTitle := "[FIRING][WARNING] Summary"
	expectedMessage := "[instance] Description"
	expectedPriority := 5

	var alert Alert
	alert.Status = "firing"
	alert.Labels.Alertname = "Test alert"
	alert.Labels.Instance = "instance"
	alert.Labels.Severity = "warning"
	alert.Annotations.Summary = "Summary"
	alert.Annotations.Description = "Description"
	alert.Annotations.Priority = "wrong"
	actualTitle, actualMessage, actualPriority := ParseAlert(alert, 5)

	if expectedTitle != actualTitle {
		t.Errorf("Title was incorrect want: \"%+v\", but got: \"%+v\"", expectedTitle, actualTitle)
	}
	if expectedMessage != actualMessage {
		t.Errorf("Message was incorrect want: %+v, but got: %+v", expectedMessage, actualMessage)
	}
	if expectedPriority != actualPriority {
		t.Errorf("Priority was incorrect want: %+v, but got: %+v", expectedPriority, actualPriority)
	}
}
