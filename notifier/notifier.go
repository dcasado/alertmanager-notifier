package notifier

import (
	"fmt"
	"log"

	"github.com/dcasado/alertmanager-notifier/alertmanager"
)

const (
	GotifyType string = "gotify"
	NTFYType   string = "ntfy"
)

type ErrNotAvailable struct {
	url string
}

func NewErrNotAvailable(url string) error {
	return ErrNotAvailable{url: url}
}

func (e ErrNotAvailable) Error() string {
	return fmt.Sprintf("notifier %s not available", e.url)
}

type ErrHTTPError struct {
	code int
	msg  string
}

func NewErrHTTPError(code int, msg string) error {
	return ErrHTTPError{code: code, msg: msg}
}

func (e ErrHTTPError) Error() string {
	return fmt.Sprintf("destination returned and error. Code: %d Reason: %s", e.code, e.msg)
}

type Notifier interface {
	Notify(alert alertmanager.Alert) error
}

func New(notifierType string) Notifier {
	switch notifierType {
	case GotifyType:
		return newGotifyClient()
	case NTFYType:
		return newNTFYClient()
	default:
		log.Fatalf("Wrong notifier type %s", notifierType)
		return nil
	}
}
