package integration

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/pentops/flowtest"
	"github.com/pentops/golib/gl"
	"github.com/pentops/o5-test-app/internal/gen/test/v1/test_pb"
	"github.com/pentops/o5-test-app/internal/gen/test/v1/test_spb"
	"github.com/pentops/o5-test-app/internal/gen/test/v1/test_tpb"
	"github.com/pentops/realms/authtest"
)

func TestEventLog(t *testing.T) {
	flow, uu := NewUniverse(t)
	defer flow.RunSteps(t)

	// Variables which cross step boundaries are declared
	var greetingID string
	var createdEvent *test_tpb.GreetingEventMessage

	var requestMessage *test_tpb.GreetingMessage

	flow.Step("Hello", func(ctx context.Context, t flowtest.Asserter) {
		ctx = authtest.JWTContext(ctx)

		greetingID = uuid.NewString()

		res, err := uu.GreetingCommand.Hello(ctx, &test_spb.HelloRequest{
			GreetingId: greetingID,
			Name:       "World",
		})
		t.NoError(err)
		t.Equal("World", res.Greeting.Data.Name)
		t.Equal(TestVersion, res.Greeting.Data.AppVersion)
		t.Equal(test_pb.GreetingStatus_INITIATED, res.Greeting.Status)

		requestMessage = uu.PopGreeting(t)
		t.Equal(greetingID, requestMessage.GreetingId)

		createdEvent = uu.PopGreetingEvent(t)
	})

	flow.Step("Receive First Greeting Event", func(ctx context.Context, t flowtest.Asserter) {
		t.MustMessage(uu.GreetingPublish.GreetingEvent(ctx, createdEvent))
	})

	flow.Step("Check Messages", func(ctx context.Context, t flowtest.Asserter) {
		ctx = authtest.JWTContext(ctx)
		messages, err := uu.EventLog.GetMessages(ctx, &test_spb.GetMessagesRequest{
			GreetingId: gl.Ptr(greetingID),
			EventId:    gl.Ptr(createdEvent.Metadata.EventId),
		})
		t.NoError(err)
		t.Equal(1, len(messages.Messages))
		t.Equal(greetingID, messages.Messages[0].GreetingId)
	})

	flow.Step("Receive Replay Greeting Event", func(ctx context.Context, t flowtest.Asserter) {
		t.MustMessage(uu.GreetingPublish.GreetingEvent(ctx, createdEvent))
	})

	flow.Step("Check Messages", func(ctx context.Context, t flowtest.Asserter) {
		ctx = authtest.JWTContext(ctx)
		messages, err := uu.EventLog.GetMessages(ctx, &test_spb.GetMessagesRequest{
			GreetingId: gl.Ptr(greetingID),
			EventId:    gl.Ptr(createdEvent.Metadata.EventId),
		})
		t.NoError(err)
		t.Equal(2, len(messages.Messages))
		t.Equal(greetingID, messages.Messages[0].GreetingId)
		if messages.Messages[0].MessageId == messages.Messages[1].MessageId {
			t.Errorf("MessageId should be different")
		}
	})
}
