package messaging

// Code generated by jsonapi. DO NOT EDIT.
// Source: github.com/pentops/o5-test-app/internal/genclient/j5/messaging/v1/messaging

import ()

// MessageCause Proto: MessageCause
type MessageCause struct {
	Method    string `json:"method"`
	MessageId string `json:"messageId,omitempty"`
	SourceApp string `json:"sourceApp,omitempty"`
	SourceEnv string `json:"sourceEnv,omitempty"`
}
