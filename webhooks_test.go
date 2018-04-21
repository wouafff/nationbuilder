package nationbuilder

import (
	"testing"
)

func TestWebhookString(t *testing.T) {
	w := &Webhook{
		ID:      "abc",
		Version: 2,
		URL:     "http://example.com",
		Event:   "person_created",
	}

	wString, expected := w.String(), "Webhook ID: abc, Version: 2, URL: http://example.com, Event: person_created"
	if wString != expected {
		t.Errorf("Expected webhook string to be %s but saw %s", expected, wString)
	}
}
