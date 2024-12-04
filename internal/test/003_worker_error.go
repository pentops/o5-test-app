package test

import (
	"context"

	"github.com/google/uuid"
	"github.com/pentops/flowtest"
	"github.com/pentops/o5-test-app/internal/genclient/o5/dante/v1/dante"
	"github.com/pentops/o5-test-app/internal/genclient/test/v1/test"
	"github.com/pentops/o5-test-app/internal/test/universe"
	"google.golang.org/grpc/codes"
)

func WorkerErrorTests(flow flowtest.StepSetter, uu *universe.Universe) {

	greetingID := uuid.NewString()

	flow.Step("Hello", func(ctx context.Context, t flowtest.Asserter) {

		greetingClient := test.NewGreetingCommandService(uu.Client)

		_, err := greetingClient.Hello(ctx, &test.HelloRequest{
			GreetingId: greetingID,
			Name:       "World",
			WorkerError: &test.TestError{
				Code:    uint32(codes.InvalidArgument),
				Message: "Test Error",
			},
		})
		if err != nil {
			t.Fatal(err)
		}

	})

	flow.Step("Dante Dead", func(ctx context.Context, t flowtest.Asserter) {

		danteClient := dante.NewDeadMessageQueryService(uu.MetaClient)

		msgs, err := danteClient.ListDeadMessages(ctx, &dante.ListDeadMessagesRequest{})
		t.NoError(err)

		t.Log(msgs)

	})

}
