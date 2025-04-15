package test

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/pentops/flowtest"
	"github.com/pentops/o5-test-app/internal/genclient/test/v1/test"
	"github.com/pentops/o5-test-app/internal/test/universe"
)

func ReplayTests(flow flowtest.StepSetter, uu *universe.Universe) {
	greetingID := uuid.NewString()

	flow.Step("Hello", func(ctx context.Context, t flowtest.Asserter) {
		greetingClient := test.NewGreetingCommandService(uu.Client)

		reply, err := greetingClient.Hello(ctx, &test.HelloRequest{
			GreetingId: greetingID,
			Name:       "World",
		})
		t.NoError(err)

		t.NotNil(reply, reply.Greeting, reply.Greeting.Data)
		t.Equal(reply.Greeting.Data.Name, "World")
	})

	flow.Step("Wait Loop", func(ctx context.Context, t flowtest.Asserter) {
		queryClient := test.NewGreetingQueryService(uu.Client)

		for {
			reply, err := queryClient.GreetingGet(ctx, &test.GreetingGetRequest{
				GreetingId: greetingID,
			})
			t.NoError(err)

			t.NotNil(reply, reply.Greeting, reply.Greeting.Data)
			t.Equal(reply.Greeting.Data.Name, "World")

			if reply.Greeting.Data.ReplyMessage != nil {
				t.Equal(*reply.Greeting.Data.ReplyMessage, "Hello World")
				return
			}
			time.Sleep(1 * time.Second)
		}
	})
}
