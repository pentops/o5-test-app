package state

import (
	"context"

	"github.com/pentops/o5-messaging/o5msg"
	"github.com/pentops/o5-test-app/internal/gen/test/v1/test_pb"
	"github.com/pentops/o5-test-app/internal/gen/test/v1/test_tpb"
	"github.com/pentops/protostate/psm"
)

type batonSender struct {
	o5msg.TopicSet
}

func (bs *batonSender) Collect(tb test_pb.GreetingPSMHookBaton, msg o5msg.Message) {
	tb.SideEffect(msg)
}

func NewGreetingPSM() (*test_pb.GreetingPSM, error) {
	sm, err := test_pb.GreetingPSMBuilder().
		SystemActor(psm.MustSystemActor("216B6C2E-D996-492C-B80C-9AAD0CCFEEC4")).
		BuildStateMachine()

	if err != nil {
		return nil, err
	}

	bs := &batonSender{}
	greetingTopic := test_tpb.NewTestTopicCollector(bs)

	sm.From(0).
		OnEvent(test_pb.GreetingPSMEventInitiated).
		SetStatus(test_pb.GreetingStatus_INITIATED).
		Mutate(test_pb.GreetingPSMMutation(func(
			state *test_pb.GreetingData,
			event *test_pb.GreetingEventType_Initiated,
		) error {
			state.Name = event.Name
			state.AppVersion = event.AppVersion
			return nil
		})).
		LogicHook(test_pb.GreetingPSMLogicHook(func(
			ctx context.Context,
			tb test_pb.GreetingPSMHookBaton,
			state *test_pb.GreetingState,
			event *test_pb.GreetingEventType_Initiated,
		) error {
			msg := &test_tpb.GreetingMessage{
				GreetingId:  state.Keys.GreetingId,
				Name:        event.Name,
				WorkerError: event.WorkerError,
			}

			greetingTopic.Greeting(tb, msg)

			return nil
		}))

	sm.From(test_pb.GreetingStatus_INITIATED).
		OnEvent(test_pb.GreetingPSMEventReplied).
		SetStatus(test_pb.GreetingStatus_REPLIED).
		Mutate(test_pb.GreetingPSMMutation(func(
			state *test_pb.GreetingData,
			event *test_pb.GreetingEventType_Replied,
		) error {
			state.ReplyMessage = &event.ReplyMessage
			return nil
		}))

	return sm, nil
}
