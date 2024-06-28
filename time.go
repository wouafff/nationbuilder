package nationbuilder

import (
	"strings"
	"time"
)

// The time format used by nationbuilder - usable in a call to format
// const TimeFormat = "2006-01-02T15:04:05-07:00"
const (
	TimeFormatWithOffset = "2006-01-02T15:04:05-07:00"
	TimeFormatUTC        = "2006-01-02T15:04:05Z"
)

// A wrapper around time.Time to allow deserialising a string into a time object
type Date struct {
	Time *time.Time
}

// A nationbuilder formatted representation of the time
func (n *Date) String() string {
	if n.Time == nil {
		return ""
	}

	return n.Time.Format(TimeFormatWithOffset)
}

// Implement json.Marshaller
func (n *Date) MarshalJSON() ([]byte, error) {
	return []byte(`"` + n.String() + `"`), nil
}

// Implement json.Unmarshaller
func (n *Date) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	if s == "" {
		return nil
	}
	var err error
	var t time.Time

	// Try parsing with each known format
	formats := []string{TimeFormatWithOffset, TimeFormatUTC}
	for _, format := range formats {
		t, err = time.Parse(format, s)
		if err == nil {
			n.Time = &t
			return nil
		}
	}

	// If none of the formats worked, return the last error
	return err
}

// Create a new Date from a string representation of a date
// which follows nationbuilder's date format
func NewDate(date string) (*Date, error) {
	t, err := time.Parse(TimeFormatWithOffset, date)
	if err != nil {
		t, err = time.Parse(TimeFormatUTC, date)
		if err != nil {
			return nil, err
		}
	}
	return &Date{
		Time: &t,
	}, nil
}

// Shorthand function to return a Date wrapper around a time.Time object
func NewDateFromTime(t time.Time) *Date {
	return &Date{
		Time: &t,
	}
}
