package state

// Code generated by jsonapi. DO NOT EDIT.
// Source: github.com/pentops/o5-test-app/internal/genclient/psm/state/v1/state

import (
	time "time"

	auth "github.com/pentops/o5-test-app/internal/genclient/o5/auth/v1/auth"
)

// CommandCause Proto: psm.state.v1.CommandCause
type CommandCause struct {
	MethodName string      `json:"methodName,omitempty"`
	Actor      *auth.Actor `json:"actor,omitempty"`
}

// ExternalEventCause Proto: psm.state.v1.ExternalEventCause
type ExternalEventCause struct {
	SystemName string  `json:"systemName,omitempty"`
	EventName  string  `json:"eventName,omitempty"`
	ExternalId *string `json:"externalId,omitempty"`
}

// ReplyCause Proto: psm.state.v1.ReplyCause
type ReplyCause struct {
	Request *PSMEventCause `json:"request,omitempty"`
	Async   bool           `json:"async"`
}

// StateMetadata Proto: psm.state.v1.StateMetadata
type StateMetadata struct {
	CreatedAt    *time.Time `json:"createdAt,omitempty"`
	UpdatedAt    *time.Time `json:"updatedAt,omitempty"`
	LastSequence int64      `json:"lastSequence,omitempty"`
}

// EventMetadata Proto: psm.state.v1.EventMetadata
type EventMetadata struct {
	EventId   string     `json:"eventId,omitempty"`
	Sequence  int64      `json:"sequence,omitempty"`
	Timestamp *time.Time `json:"timestamp"`
	Cause     *Cause     `json:"cause,omitempty"`
}

// Cause Proto: psm.state.v1.Cause
type Cause struct {
	Type *Cause_type `json:"type,omitempty"`
}

// Cause_type Proto: psm.state.v1.Cause.type
type Cause_type struct {
	PsmEvent      *PSMEventCause      `json:"psmEvent,omitempty"`
	Command       *CommandCause       `json:"command,omitempty"`
	ExternalEvent *ExternalEventCause `json:"externalEvent,omitempty"`
	Reply         *ReplyCause         `json:"reply,omitempty"`
}

func (s Cause_type) OneofKey() string {
	if s.PsmEvent != nil {
		return "psmEvent"
	}
	if s.Command != nil {
		return "command"
	}
	if s.ExternalEvent != nil {
		return "externalEvent"
	}
	if s.Reply != nil {
		return "reply"
	}
	return ""
}

func (s Cause_type) Type() interface{} {
	if s.PsmEvent != nil {
		return s.PsmEvent
	}
	if s.Command != nil {
		return s.Command
	}
	if s.ExternalEvent != nil {
		return s.ExternalEvent
	}
	if s.Reply != nil {
		return s.Reply
	}
	return nil
}

// PSMEventCause Proto: psm.state.v1.PSMEventCause
type PSMEventCause struct {
	EventId      string `json:"eventId,omitempty"`
	StateMachine string `json:"stateMachine,omitempty"`
}
