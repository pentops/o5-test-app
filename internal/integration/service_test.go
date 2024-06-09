package integration

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/pentops/flowtest"
	"github.com/pentops/o5-test-app/internal/gen/test/v1/test_pb"
	"github.com/pentops/o5-test-app/internal/gen/test/v1/test_spb"
	"github.com/pentops/o5-test-app/internal/gen/test/v1/test_tpb"
)

func TestService(t *testing.T) {

	flow, uu := NewUniverse(t)
	defer flow.RunSteps(t)

	// Variables which cross step boundaries are declared
	var greetingID string
	var requestMessage *test_tpb.GreetingMessage

	flow.Step("Hello", func(ctx context.Context, t flowtest.Asserter) {
		greetingID = uuid.NewString()

		res, err := uu.GreetingCommand.Hello(ctx, &test_spb.HelloRequest{
			GreetingId: greetingID,
			Name:       "World",
		})
		t.NoError(err)
		t.Equal("World", res.Greeting.Data.Name)
		t.Equal(TestVersion, res.Greeting.Data.AppVersion)
		t.Equal(test_pb.GreetingStatus_INITIATED, res.Greeting.Status)

		requestMessage = &test_tpb.GreetingMessage{}
		uu.Outbox.PopMessage(t, requestMessage)

		t.Equal(greetingID, requestMessage.GreetingId)
	})

	flow.Step("Reply", func(ctx context.Context, t flowtest.Asserter) {
		res, err := uu.TestTopic.Greeting(ctx, requestMessage)
		t.NoError(err)
		if res == nil {
			t.Errorf("res is nil, shoud be empty")
		}
	})

	flow.Step("Check", func(ctx context.Context, t flowtest.Asserter) {
		greeting, err := uu.GreetingQuery.GetGreeting(ctx, &test_spb.GetGreetingRequest{
			GreetingId: greetingID,
		})
		t.NoError(err)
		t.Equal("Hello World", *greeting.Greeting.Data.ReplyMessage)
	})

}
