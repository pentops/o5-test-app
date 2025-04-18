// Code generated by protoc-gen-go-psm. DO NOT EDIT.

package test_pb

import (
	context "context"
	fmt "fmt"
	psm_j5pb "github.com/pentops/j5/gen/j5/state/v1/psm_j5pb"
	psm "github.com/pentops/protostate/psm"
	sqrlx "github.com/pentops/sqrlx.go/sqrlx"
)

// PSM GreetingPSM

type GreetingPSM = psm.StateMachine[
	*GreetingKeys,    // implements psm.IKeyset
	*GreetingState,   // implements psm.IState
	GreetingStatus,   // implements psm.IStatusEnum
	*GreetingData,    // implements psm.IStateData
	*GreetingEvent,   // implements psm.IEvent
	GreetingPSMEvent, // implements psm.IInnerEvent
]

type GreetingPSMDB = psm.DBStateMachine[
	*GreetingKeys,    // implements psm.IKeyset
	*GreetingState,   // implements psm.IState
	GreetingStatus,   // implements psm.IStatusEnum
	*GreetingData,    // implements psm.IStateData
	*GreetingEvent,   // implements psm.IEvent
	GreetingPSMEvent, // implements psm.IInnerEvent
]

type GreetingPSMEventSpec = psm.EventSpec[
	*GreetingKeys,    // implements psm.IKeyset
	*GreetingState,   // implements psm.IState
	GreetingStatus,   // implements psm.IStatusEnum
	*GreetingData,    // implements psm.IStateData
	*GreetingEvent,   // implements psm.IEvent
	GreetingPSMEvent, // implements psm.IInnerEvent
]

type GreetingPSMEventKey = string

const (
	GreetingPSMEventNil       GreetingPSMEventKey = "<nil>"
	GreetingPSMEventInitiated GreetingPSMEventKey = "initiated"
	GreetingPSMEventReplied   GreetingPSMEventKey = "replied"
)

// EXTEND GreetingKeys with the psm.IKeyset interface

// PSMIsSet is a helper for != nil, which does not work with generic parameters
func (msg *GreetingKeys) PSMIsSet() bool {
	return msg != nil
}

// PSMFullName returns the full name of state machine with package prefix
func (msg *GreetingKeys) PSMFullName() string {
	return "test.v1.greeting"
}
func (msg *GreetingKeys) PSMKeyValues() (map[string]string, error) {
	keyset := map[string]string{
		"greeting_id": msg.GreetingId,
	}
	return keyset, nil
}

// EXTEND GreetingState with the psm.IState interface

// PSMIsSet is a helper for != nil, which does not work with generic parameters
func (msg *GreetingState) PSMIsSet() bool {
	return msg != nil
}

func (msg *GreetingState) PSMMetadata() *psm_j5pb.StateMetadata {
	if msg.Metadata == nil {
		msg.Metadata = &psm_j5pb.StateMetadata{}
	}
	return msg.Metadata
}

func (msg *GreetingState) PSMKeys() *GreetingKeys {
	return msg.Keys
}

func (msg *GreetingState) SetStatus(status GreetingStatus) {
	msg.Status = status
}

func (msg *GreetingState) SetPSMKeys(inner *GreetingKeys) {
	msg.Keys = inner
}

func (msg *GreetingState) PSMData() *GreetingData {
	if msg.Data == nil {
		msg.Data = &GreetingData{}
	}
	return msg.Data
}

// EXTEND GreetingData with the psm.IStateData interface

// PSMIsSet is a helper for != nil, which does not work with generic parameters
func (msg *GreetingData) PSMIsSet() bool {
	return msg != nil
}

// EXTEND GreetingEvent with the psm.IEvent interface

// PSMIsSet is a helper for != nil, which does not work with generic parameters
func (msg *GreetingEvent) PSMIsSet() bool {
	return msg != nil
}

func (msg *GreetingEvent) PSMMetadata() *psm_j5pb.EventMetadata {
	if msg.Metadata == nil {
		msg.Metadata = &psm_j5pb.EventMetadata{}
	}
	return msg.Metadata
}

func (msg *GreetingEvent) PSMKeys() *GreetingKeys {
	return msg.Keys
}

func (msg *GreetingEvent) SetPSMKeys(inner *GreetingKeys) {
	msg.Keys = inner
}

// PSMEventKey returns the GreetingPSMEventPSMEventKey for the event, implementing psm.IEvent
func (msg *GreetingEvent) PSMEventKey() GreetingPSMEventKey {
	tt := msg.UnwrapPSMEvent()
	if tt == nil {
		return GreetingPSMEventNil
	}
	return tt.PSMEventKey()
}

// UnwrapPSMEvent implements psm.IEvent, returning the inner event message
func (msg *GreetingEvent) UnwrapPSMEvent() GreetingPSMEvent {
	if msg == nil {
		return nil
	}
	if msg.Event == nil {
		return nil
	}
	switch v := msg.Event.Type.(type) {
	case *GreetingEventType_Initiated_:
		return v.Initiated
	case *GreetingEventType_Replied_:
		return v.Replied
	default:
		return nil
	}
}

// SetPSMEvent sets the inner event message from a concrete type, implementing psm.IEvent
func (msg *GreetingEvent) SetPSMEvent(inner GreetingPSMEvent) error {
	if msg.Event == nil {
		msg.Event = &GreetingEventType{}
	}
	switch v := inner.(type) {
	case *GreetingEventType_Initiated:
		msg.Event.Type = &GreetingEventType_Initiated_{Initiated: v}
	case *GreetingEventType_Replied:
		msg.Event.Type = &GreetingEventType_Replied_{Replied: v}
	default:
		return fmt.Errorf("invalid type %T for GreetingEventType", v)
	}
	return nil
}

type GreetingPSMEvent interface {
	psm.IInnerEvent
	PSMEventKey() GreetingPSMEventKey
}

// EXTEND GreetingEventType_Initiated with the GreetingPSMEvent interface

// PSMIsSet is a helper for != nil, which does not work with generic parameters
func (msg *GreetingEventType_Initiated) PSMIsSet() bool {
	return msg != nil
}

func (*GreetingEventType_Initiated) PSMEventKey() GreetingPSMEventKey {
	return GreetingPSMEventInitiated
}

// EXTEND GreetingEventType_Replied with the GreetingPSMEvent interface

// PSMIsSet is a helper for != nil, which does not work with generic parameters
func (msg *GreetingEventType_Replied) PSMIsSet() bool {
	return msg != nil
}

func (*GreetingEventType_Replied) PSMEventKey() GreetingPSMEventKey {
	return GreetingPSMEventReplied
}

func GreetingPSMBuilder() *psm.StateMachineConfig[
	*GreetingKeys,    // implements psm.IKeyset
	*GreetingState,   // implements psm.IState
	GreetingStatus,   // implements psm.IStatusEnum
	*GreetingData,    // implements psm.IStateData
	*GreetingEvent,   // implements psm.IEvent
	GreetingPSMEvent, // implements psm.IInnerEvent
] {
	return &psm.StateMachineConfig[
		*GreetingKeys,    // implements psm.IKeyset
		*GreetingState,   // implements psm.IState
		GreetingStatus,   // implements psm.IStatusEnum
		*GreetingData,    // implements psm.IStateData
		*GreetingEvent,   // implements psm.IEvent
		GreetingPSMEvent, // implements psm.IInnerEvent
	]{}
}

func GreetingPSMMutation[SE GreetingPSMEvent](cb func(*GreetingData, SE) error) psm.TransitionMutation[
	*GreetingKeys,    // implements psm.IKeyset
	*GreetingState,   // implements psm.IState
	GreetingStatus,   // implements psm.IStatusEnum
	*GreetingData,    // implements psm.IStateData
	*GreetingEvent,   // implements psm.IEvent
	GreetingPSMEvent, // implements psm.IInnerEvent
	SE,               // Specific event type for the transition
] {
	return psm.TransitionMutation[
		*GreetingKeys,    // implements psm.IKeyset
		*GreetingState,   // implements psm.IState
		GreetingStatus,   // implements psm.IStatusEnum
		*GreetingData,    // implements psm.IStateData
		*GreetingEvent,   // implements psm.IEvent
		GreetingPSMEvent, // implements psm.IInnerEvent
		SE,               // Specific event type for the transition
	](cb)
}

type GreetingPSMHookBaton = psm.HookBaton[
	*GreetingKeys,    // implements psm.IKeyset
	*GreetingState,   // implements psm.IState
	GreetingStatus,   // implements psm.IStatusEnum
	*GreetingData,    // implements psm.IStateData
	*GreetingEvent,   // implements psm.IEvent
	GreetingPSMEvent, // implements psm.IInnerEvent
]

func GreetingPSMLogicHook[SE GreetingPSMEvent](cb func(context.Context, GreetingPSMHookBaton, *GreetingState, SE) error) psm.TransitionLogicHook[
	*GreetingKeys,    // implements psm.IKeyset
	*GreetingState,   // implements psm.IState
	GreetingStatus,   // implements psm.IStatusEnum
	*GreetingData,    // implements psm.IStateData
	*GreetingEvent,   // implements psm.IEvent
	GreetingPSMEvent, // implements psm.IInnerEvent
	SE,               // Specific event type for the transition
] {
	return psm.TransitionLogicHook[
		*GreetingKeys,    // implements psm.IKeyset
		*GreetingState,   // implements psm.IState
		GreetingStatus,   // implements psm.IStatusEnum
		*GreetingData,    // implements psm.IStateData
		*GreetingEvent,   // implements psm.IEvent
		GreetingPSMEvent, // implements psm.IInnerEvent
		SE,               // Specific event type for the transition
	](cb)
}
func GreetingPSMDataHook[SE GreetingPSMEvent](cb func(context.Context, sqrlx.Transaction, *GreetingState, SE) error) psm.TransitionDataHook[
	*GreetingKeys,    // implements psm.IKeyset
	*GreetingState,   // implements psm.IState
	GreetingStatus,   // implements psm.IStatusEnum
	*GreetingData,    // implements psm.IStateData
	*GreetingEvent,   // implements psm.IEvent
	GreetingPSMEvent, // implements psm.IInnerEvent
	SE,               // Specific event type for the transition
] {
	return psm.TransitionDataHook[
		*GreetingKeys,    // implements psm.IKeyset
		*GreetingState,   // implements psm.IState
		GreetingStatus,   // implements psm.IStatusEnum
		*GreetingData,    // implements psm.IStateData
		*GreetingEvent,   // implements psm.IEvent
		GreetingPSMEvent, // implements psm.IInnerEvent
		SE,               // Specific event type for the transition
	](cb)
}
func GreetingPSMLinkHook[SE GreetingPSMEvent, DK psm.IKeyset, DIE psm.IInnerEvent](
	linkDestination psm.LinkDestination[DK, DIE],
	cb func(context.Context, *GreetingState, SE, func(DK, DIE)) error,
) psm.LinkHook[
	*GreetingKeys,    // implements psm.IKeyset
	*GreetingState,   // implements psm.IState
	GreetingStatus,   // implements psm.IStatusEnum
	*GreetingData,    // implements psm.IStateData
	*GreetingEvent,   // implements psm.IEvent
	GreetingPSMEvent, // implements psm.IInnerEvent
	SE,               // Specific event type for the transition
	DK,               // Destination Keys
	DIE,              // Destination Inner Event
] {
	return psm.LinkHook[
		*GreetingKeys,    // implements psm.IKeyset
		*GreetingState,   // implements psm.IState
		GreetingStatus,   // implements psm.IStatusEnum
		*GreetingData,    // implements psm.IStateData
		*GreetingEvent,   // implements psm.IEvent
		GreetingPSMEvent, // implements psm.IInnerEvent
		SE,               // Specific event type for the transition
		DK,               // Destination Keys
		DIE,              // Destination Inner Event
	]{
		Derive:      cb,
		Destination: linkDestination,
	}
}
func GreetingPSMGeneralLogicHook(cb func(context.Context, GreetingPSMHookBaton, *GreetingState, *GreetingEvent) error) psm.GeneralLogicHook[
	*GreetingKeys,    // implements psm.IKeyset
	*GreetingState,   // implements psm.IState
	GreetingStatus,   // implements psm.IStatusEnum
	*GreetingData,    // implements psm.IStateData
	*GreetingEvent,   // implements psm.IEvent
	GreetingPSMEvent, // implements psm.IInnerEvent
] {
	return psm.GeneralLogicHook[
		*GreetingKeys,    // implements psm.IKeyset
		*GreetingState,   // implements psm.IState
		GreetingStatus,   // implements psm.IStatusEnum
		*GreetingData,    // implements psm.IStateData
		*GreetingEvent,   // implements psm.IEvent
		GreetingPSMEvent, // implements psm.IInnerEvent
	](cb)
}
func GreetingPSMGeneralStateDataHook(cb func(context.Context, sqrlx.Transaction, *GreetingState) error) psm.GeneralStateDataHook[
	*GreetingKeys,    // implements psm.IKeyset
	*GreetingState,   // implements psm.IState
	GreetingStatus,   // implements psm.IStatusEnum
	*GreetingData,    // implements psm.IStateData
	*GreetingEvent,   // implements psm.IEvent
	GreetingPSMEvent, // implements psm.IInnerEvent
] {
	return psm.GeneralStateDataHook[
		*GreetingKeys,    // implements psm.IKeyset
		*GreetingState,   // implements psm.IState
		GreetingStatus,   // implements psm.IStatusEnum
		*GreetingData,    // implements psm.IStateData
		*GreetingEvent,   // implements psm.IEvent
		GreetingPSMEvent, // implements psm.IInnerEvent
	](cb)
}
func GreetingPSMGeneralEventDataHook(cb func(context.Context, sqrlx.Transaction, *GreetingState, *GreetingEvent) error) psm.GeneralEventDataHook[
	*GreetingKeys,    // implements psm.IKeyset
	*GreetingState,   // implements psm.IState
	GreetingStatus,   // implements psm.IStatusEnum
	*GreetingData,    // implements psm.IStateData
	*GreetingEvent,   // implements psm.IEvent
	GreetingPSMEvent, // implements psm.IInnerEvent
] {
	return psm.GeneralEventDataHook[
		*GreetingKeys,    // implements psm.IKeyset
		*GreetingState,   // implements psm.IState
		GreetingStatus,   // implements psm.IStatusEnum
		*GreetingData,    // implements psm.IStateData
		*GreetingEvent,   // implements psm.IEvent
		GreetingPSMEvent, // implements psm.IInnerEvent
	](cb)
}
func GreetingPSMEventPublishHook(cb func(context.Context, psm.Publisher, *GreetingState, *GreetingEvent) error) psm.EventPublishHook[
	*GreetingKeys,    // implements psm.IKeyset
	*GreetingState,   // implements psm.IState
	GreetingStatus,   // implements psm.IStatusEnum
	*GreetingData,    // implements psm.IStateData
	*GreetingEvent,   // implements psm.IEvent
	GreetingPSMEvent, // implements psm.IInnerEvent
] {
	return psm.EventPublishHook[
		*GreetingKeys,    // implements psm.IKeyset
		*GreetingState,   // implements psm.IState
		GreetingStatus,   // implements psm.IStatusEnum
		*GreetingData,    // implements psm.IStateData
		*GreetingEvent,   // implements psm.IEvent
		GreetingPSMEvent, // implements psm.IInnerEvent
	](cb)
}
func GreetingPSMUpsertPublishHook(cb func(context.Context, psm.Publisher, *GreetingState) error) psm.UpsertPublishHook[
	*GreetingKeys,  // implements psm.IKeyset
	*GreetingState, // implements psm.IState
	GreetingStatus, // implements psm.IStatusEnum
	*GreetingData,  // implements psm.IStateData
] {
	return psm.UpsertPublishHook[
		*GreetingKeys,  // implements psm.IKeyset
		*GreetingState, // implements psm.IState
		GreetingStatus, // implements psm.IStatusEnum
		*GreetingData,  // implements psm.IStateData
	](cb)
}

func (event *GreetingEvent) EventPublishMetadata() *psm_j5pb.EventPublishMetadata {
	tenantKeys := make([]*psm_j5pb.EventTenant, 0)
	return &psm_j5pb.EventPublishMetadata{
		EventId:   event.Metadata.EventId,
		Sequence:  event.Metadata.Sequence,
		Timestamp: event.Metadata.Timestamp,
		Cause:     event.Metadata.Cause,
		Auth: &psm_j5pb.PublishAuth{
			TenantKeys: tenantKeys,
		},
	}
}
