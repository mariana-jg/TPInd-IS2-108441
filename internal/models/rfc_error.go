// Package models has the structs that are used to represent the data in the application.
package models

// RFCError is the struct that represents an error in the RFC format
type RFCError struct {
	Type     string `json:"type"`
	Title    string `json:"title"`
	Status   int    `json:"status"`
	Detail   string `json:"detail"`
	Instance string `json:"instance"`
}

// NewRFCError is the function that creates a new RFCError
func NewRFCError(status int, title string, detail string, instance string) RFCError {
	return RFCError{
		Type:     "about:blank",
		Title:    title,
		Status:   status,
		Detail:   detail,
		Instance: instance,
	}
}
