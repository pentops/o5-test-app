package integration

import (
	"context"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/pentops/flowtest"
	"github.com/pentops/o5-test-app/internal/gen/test/v1/test_pb"
	"github.com/pentops/o5-test-app/internal/gen/test/v1/test_spb"
	"github.com/pentops/o5-test-app/internal/gen/test/v1/test_tpb"
	"github.com/pentops/realms/authtest"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestService(t *testing.T) {
	flow, uu := NewUniverse(t)
	defer flow.RunSteps(t)

	// Variables which cross step boundaries are declared
	var greetingID string

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

		uu.PopGreetingEvent(t)
	})

	flow.Step("Reply", func(ctx context.Context, t flowtest.Asserter) {
		res, err := uu.TestTopic.Greeting(ctx, requestMessage)
		t.NoError(err)
		if res == nil {
			t.Errorf("res is nil, shoud be empty")
		}

		uu.PopGreetingEvent(t)
	})

	flow.Step("Check", func(ctx context.Context, t flowtest.Asserter) {
		ctx = authtest.JWTContext(ctx)
		greeting, err := uu.GreetingQuery.GreetingGet(ctx, &test_spb.GreetingGetRequest{
			GreetingId: greetingID,
		})
		t.NoError(err)
		t.Equal("Hello World", *greeting.Greeting.Data.ReplyMessage)
	})

}

func TestThrowError(t *testing.T) {
	flow, uu := NewUniverse(t)
	defer flow.RunSteps(t)

	flow.Step("Throw Unknown", func(ctx context.Context, t flowtest.Asserter) {
		ctx = authtest.JWTContext(ctx)
		_, err := uu.GreetingCommand.Hello(ctx, &test_spb.HelloRequest{
			GreetingId: uuid.NewString(),
			Name:       "World",

			ThrowError: &test_pb.TestError{
				Code:    0,
				Message: "Test Error",
			},
		})
		if err == nil {
			t.Fatal("expected error")
		}

		codeErr, ok := status.FromError(err)
		if !ok {
			t.Fatalf("expected status error, got: %v", err)
		}

		if codeErr.Code() != codes.Unknown {
			t.Fatalf("expected Unknown, got: %v", codeErr.Code())
		}
		if !strings.Contains(err.Error(), "Test Error") {
			t.Errorf("unexpected error: %v", err)
		}
	})

	flow.Step("Throw Known", func(ctx context.Context, t flowtest.Asserter) {
		ctx = authtest.JWTContext(ctx)
		_, err := uu.GreetingCommand.Hello(ctx, &test_spb.HelloRequest{
			GreetingId: uuid.NewString(),
			Name:       "World",

			ThrowError: &test_pb.TestError{
				Code:    uint32(codes.InvalidArgument),
				Message: "Test Error",
			},
		})
		if err == nil {
			t.Fatal("expected error")
		}

		codeErr, ok := status.FromError(err)
		if !ok {
			t.Fatalf("expected status error, got: %v", err)
		}

		if codeErr.Code() != codes.InvalidArgument {
			t.Fatalf("expected InvalidArgument, got: %v", codeErr.Code())
		}
		if !strings.Contains(err.Error(), "Test Error") {
			t.Errorf("unexpected error: %v", err)
		}
	})
}

func TestMessageError(t *testing.T) {
	flow, uu := NewUniverse(t)
	defer flow.RunSteps(t)

	var greetingID string
	var requestMessage *test_tpb.GreetingMessage
	flow.Step("Throw Passthrough", func(ctx context.Context, t flowtest.Asserter) {
		ctx = authtest.JWTContext(ctx)
		greetingID = uuid.NewString()
		_, err := uu.GreetingCommand.Hello(ctx, &test_spb.HelloRequest{
			GreetingId: greetingID,
			Name:       "World",

			WorkerError: &test_pb.TestError{
				Code:    uint32(codes.InvalidArgument),
				Message: "Test Error",
			},
		})
		if err != nil {
			t.Fatalf("expected no error here, should be passed to the worker: %s", err)
		}

		requestMessage = uu.PopGreeting(t)

		t.Equal(greetingID, requestMessage.GreetingId)
		if requestMessage.WorkerError == nil {
			t.Errorf("expected worker error to be set")
		}

		uu.PopGreetingEvent(t)
	})

	flow.Step("Reply", func(ctx context.Context, t flowtest.Asserter) {
		_, err := uu.TestTopic.Greeting(ctx, requestMessage)
		if err == nil {
			t.Fatalf("expected error here, should be passed to the worker")
		}

		codeErr, ok := status.FromError(err)
		if !ok {
			t.Fatalf("expected status error, got: %v", err)
		}

		if codeErr.Code() != codes.InvalidArgument {
			t.Fatalf("expected InvalidArgument, got: %v", codeErr.Code())
		}
	})
}
