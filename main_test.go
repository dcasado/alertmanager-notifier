package main

import "testing"

func Test_buildGotifyMessage(t *testing.T) {
	want := GotifyMessage{"[FIRING][WARNING] Summary", "[instance] Description", 2}

	var alert Alert
	alert.Status = "firing"
	alert.Labels.Alertname = "Test alert"
	alert.Labels.Instance = "instance"
	alert.Labels.Severity = "warning"
	alert.Annotations.Summary = "Summary"
	alert.Annotations.Description = "Description"
	alert.Annotations.Priority = "2"
	got := buildGotifyMessage(alert)

	if want != got {
		t.Errorf("Gotify message was incorrect want: %+v, but got: %+v", want, got)
	}
}

func Test_buildGotifyMessage_noInstance(t *testing.T) {
	want := GotifyMessage{"[FIRING][WARNING] Summary", "Description", 2}

	var alert Alert
	alert.Status = "firing"
	alert.Labels.Alertname = "Test alert"
	alert.Labels.Severity = "warning"
	alert.Annotations.Summary = "Summary"
	alert.Annotations.Description = "Description"
	alert.Annotations.Priority = "2"
	got := buildGotifyMessage(alert)

	if want != got {
		t.Errorf("Gotify message was incorrect want: %+v, but got: %+v", want, got)
	}
}

func Test_buildGotifyMessage_noSummary(t *testing.T) {
	want := GotifyMessage{"[FIRING][WARNING] ", "[instance] Description", 2}

	var alert Alert
	alert.Status = "firing"
	alert.Labels.Alertname = "Test alert"
	alert.Labels.Instance = "instance"
	alert.Labels.Severity = "warning"
	alert.Annotations.Description = "Description"
	alert.Annotations.Priority = "2"
	got := buildGotifyMessage(alert)

	if want != got {
		t.Errorf("Gotify message was incorrect want: %+v, but got: %+v", want, got)
	}
}

func Test_buildGotifyMessage_noDescription(t *testing.T) {
	want := GotifyMessage{"[FIRING][WARNING] Summary", "[instance] ", 2}

	var alert Alert
	alert.Status = "firing"
	alert.Labels.Alertname = "Test alert"
	alert.Labels.Instance = "instance"
	alert.Labels.Severity = "warning"
	alert.Annotations.Summary = "Summary"
	alert.Annotations.Priority = "2"
	got := buildGotifyMessage(alert)

	if want != got {
		t.Errorf("Gotify message was incorrect want: %+v, but got: %+v", want, got)
	}
}

func Test_buildGotifyMessage_noPriority(t *testing.T) {
	want := GotifyMessage{"[FIRING][WARNING] Summary", "[instance] Description", 5}

	var alert Alert
	alert.Status = "firing"
	alert.Labels.Alertname = "Test alert"
	alert.Labels.Instance = "instance"
	alert.Labels.Severity = "warning"
	alert.Annotations.Summary = "Summary"
	alert.Annotations.Description = "Description"
	got := buildGotifyMessage(alert)

	if want != got {
		t.Errorf("Gotify message was incorrect want: %+v, but got: %+v", want, got)
	}
}

func Test_buildGotifyMessage_wrongPriority(t *testing.T) {
	want := GotifyMessage{"[FIRING][WARNING] Summary", "[instance] Description", 5}

	var alert Alert
	alert.Status = "firing"
	alert.Labels.Alertname = "Test alert"
	alert.Labels.Instance = "instance"
	alert.Labels.Severity = "warning"
	alert.Annotations.Summary = "Summary"
	alert.Annotations.Description = "Description"
	alert.Annotations.Priority = "wrong"
	got := buildGotifyMessage(alert)

	if want != got {
		t.Errorf("Gotify message was incorrect want: %+v, but got: %+v", want, got)
	}
}
