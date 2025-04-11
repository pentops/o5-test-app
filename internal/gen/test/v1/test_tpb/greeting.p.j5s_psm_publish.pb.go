// Code generated by protoc-gen-go-psm. DO NOT EDIT.

package test_tpb

import (
	context "context"
	test_pb "github.com/pentops/o5-test-app/internal/gen/test/v1/test_pb"
	psm "github.com/pentops/protostate/psm"
)

// Publish Toipc for test.v1.Greeting
func PublishGreeting() psm.EventPublishHook[
	*test_pb.GreetingKeys,    // implements psm.IKeyset
	*test_pb.GreetingState,   // implements psm.IState
	test_pb.GreetingStatus,   // implements psm.IStatusEnum
	*test_pb.GreetingData,    // implements psm.IStateData
	*test_pb.GreetingEvent,   // implements psm.IEvent
	test_pb.GreetingPSMEvent, // implements psm.IInnerEvent
] {
	return test_pb.GreetingPSMEventPublishHook(func(
		ctx context.Context,
		publisher psm.Publisher,
		state *test_pb.GreetingState,
		event *test_pb.GreetingEvent,
	) error {
		publisher.Publish(&GreetingEventMessage{
			Metadata: event.EventPublishMetadata(),
			Keys:     event.Keys,
			Event:    event.Event,
			Data:     state.Data,
			Status:   state.Status,
		})
		return nil
	})
}
