package runner

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/pentops/flowtest"
	"github.com/pentops/flowtest/runner"
	"github.com/pentops/flowtest/runner/testclient"
	"github.com/pentops/o5-test-app/internal/genclient/o5/dante/v1/dante"
	"github.com/pentops/o5-test-app/internal/genclient/test/v1/test"
	"google.golang.org/grpc/codes"
)

type APIConfig struct {
	APIRoot     string
	MetaRoot    string
	BearerToken string
}

type GreetingRequest struct {
	Name       string `json:"name"`
	GreetingID string `json:"greetingId"`
}

type GreetingResponse struct {
	Greeting GreetingState `json:"greeting"`
}

type GreetingState struct {
	Status string       `json:"status"`
	Data   GreetingData `json:"data"`
	Keys   GreetingKeys `json:"keys"`
}

type GreetingData struct {
	Name         string  `json:"name"`
	ReplyMessage *string `json:"replyMessage"`
}

type GreetingKeys struct {
	GreetingID string `json:"greetingId"`
}

func LogAPIRequest(flow flowtest.StepSetter) func(req *testclient.RequestLog) {
	return func(req *testclient.RequestLog) {
		flow.Log(req)
	}
}

func UniverseWrapper(cfg *APIConfig, callback func(flow flowtest.StepSetter, uu *Universe)) func(flow flowtest.StepSetter) {
	return func(flow flowtest.StepSetter) {

		universe := &Universe{}

		flow.Setup(func(ctx context.Context, t flowtest.Asserter) error {
			t.Log("SETUP")
			client, err := testclient.NewAPI(cfg.APIRoot)
			if err != nil {
				return err
			}

			client.Auth = testclient.BearerToken(cfg.BearerToken)

			client.Logger = LogAPIRequest(flow)
			universe.Client = client

			metaClient, err := testclient.NewAPI(cfg.MetaRoot)
			if err != nil {
				return err
			}

			metaClient.Auth = testclient.BearerToken(cfg.BearerToken)

			metaClient.Logger = LogAPIRequest(flow)
			universe.MetaClient = metaClient

			return nil
		})

		callback(flow, universe)
	}
}

func Run(ctx context.Context, cfg *APIConfig, tags []string) error {
	testSet := runner.TestSet{}

	testSet.Register(1, "Greeting", UniverseWrapper(cfg, GreetingTests))
	testSet.Register(2, "HandlerError", UniverseWrapper(cfg, HandlerErrorTests))
	testSet.Register(3, "WorkerErrorTests", UniverseWrapper(cfg, WorkerErrorTests), "name=worker-error")

	return testSet.Run(ctx, tags)
}

type Universe struct {
	Client     *testclient.API
	MetaClient *testclient.API
}

func GreetingTests(flow flowtest.StepSetter, uu *Universe) {

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
			reply, err := queryClient.GetGreeting(ctx, &test.GetGreetingRequest{
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

func HandlerErrorTests(flow flowtest.StepSetter, uu *Universe) {

	greetingID := uuid.NewString()

	flow.Step("Hello", func(ctx context.Context, t flowtest.Asserter) {

		greetingClient := test.NewGreetingCommandService(uu.Client)

		_, err := greetingClient.Hello(ctx, &test.HelloRequest{
			GreetingId: greetingID,
			Name:       "World",
			ThrowError: &test.TestError{
				Code:    int64(codes.InvalidArgument),
				Message: "Test Error",
			},
		})
		if err == nil {
			t.Fatal("expected error")
		}

		t.Log(err.Error())

	})

}

func WorkerErrorTests(flow flowtest.StepSetter, uu *Universe) {

	greetingID := uuid.NewString()

	flow.Step("Hello", func(ctx context.Context, t flowtest.Asserter) {

		greetingClient := test.NewGreetingCommandService(uu.Client)

		_, err := greetingClient.Hello(ctx, &test.HelloRequest{
			GreetingId: greetingID,
			Name:       "World",
			WorkerError: &test.TestError{
				Code:    int64(codes.InvalidArgument),
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
