package domain

import "time"

// NullableTime returns nil if the time is zero, otherwise returns a pointer to the time
func NullableTime(t time.Time) *time.Time {
	if t.IsZero() {
		return nil
	}
	return &t
}

// NullableString returns nil if the string is empty, otherwise returns a pointer to the string
func NullableString(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
