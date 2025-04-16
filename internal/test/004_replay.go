package test

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/pentops/flowtest"
	"github.com/pentops/golib/gl"
	"github.com/pentops/o5-test-app/internal/genclient/o5/ges/v1/ges"
	"github.com/pentops/o5-test-app/internal/genclient/test/v1/test"
	"github.com/pentops/o5-test-app/internal/test/universe"
)

func ReplayTests(flow flowtest.StepSetter, uu *universe.Universe) {
	greetingID := uuid.NewString()

	var logClient *test.EventLogService
	var greetingClient *test.GreetingCommandService
	var gesClient *ges.CombinedClient

	flow.Setup(func(ctx context.Context, t flowtest.Asserter) error {
		logClient = test.NewEventLogService(uu.Client)
		greetingClient = test.NewGreetingCommandService(uu.Client)
		gesClient = ges.NewCombinedClient(uu.Client)
		return nil
	})

	flow.Step("Hello", func(ctx context.Context, t flowtest.Asserter) {

		reply, err := greetingClient.Hello(ctx, &test.HelloRequest{
			GreetingId: greetingID,
			Name:       "World",
		})
		t.NoError(err)

		t.NotNil(reply, reply.Greeting, reply.Greeting.Data)
		t.Equal(reply.Greeting.Data.Name, "World")
	})

	flow.Step("Messages in Test", func(ctx context.Context, t flowtest.Asserter) {

		for ii := 0; ii < 10; ii++ {
			reply, err := logClient.GetMessages(ctx, &test.GetMessagesRequest{
				GreetingId: gl.Ptr(greetingID),
			})
			t.NoError(err)

			for _, message := range reply.Messages {
				t.Logf("Event %s", message.EventId)
			}
			if len(reply.Messages) == 2 {
				return
			}

			time.Sleep(1 * time.Second)
		}
		t.Error("Messages not found")
	})

	flow.Step("Replay Messages in Test", func(ctx context.Context, t flowtest.Asserter) {
		for ii := 0; ii < 10; ii++ {
			// re-trigger replay on each iteration, as this is a race
			_, err := gesClient.ReplayEvents(ctx, &ges.ReplayEventsRequest{
				QueueUrl:    uu.ReplayQueueURL,
				GrpcService: "test.v1.topic.GreetingPublishTopic",
				GrpcMethod:  "GreetingEvent",
			})
			t.NoError(err)
			reply, err := logClient.GetMessages(ctx, &test.GetMessagesRequest{
				GreetingId: gl.Ptr(greetingID),
			})
			t.NoError(err)

			for _, message := range reply.Messages {
				t.Logf("Event %s", message.EventId)
			}
			if len(reply.Messages) > 2 {
				return
			}

			time.Sleep(1 * time.Second)
		}
		t.Error("Replayed messages not found")

	})

}
