package integration

import (
	"context"
	"testing"

	"github.com/pentops/flowtest"
	"github.com/pentops/log.go/log"
	"github.com/pentops/o5-test-app/internal/gen/test/v1/test_spb"
	"github.com/pentops/o5-test-app/internal/gen/test/v1/test_tpb"
	"github.com/pentops/o5-test-app/internal/service"
	"github.com/pentops/outbox.pg.go/outboxtest"
	"github.com/pentops/pgtest.go/pgtest"
)

type Universe struct {
	Outbox          *outboxtest.OutboxAsserter
	GreetingCommand test_spb.GreetingCommandServiceClient
	GreetingQuery   test_spb.GreetingQueryServiceClient
	TestTopic       test_tpb.TestTopicClient
}

func NewUniverse(t *testing.T) (*flowtest.Stepper[*testing.T], *Universe) {
	name := t.Name()
	stepper := flowtest.NewStepper[*testing.T](name)
	uu := &Universe{}

	stepper.Setup(func(ctx context.Context, t flowtest.Asserter) error {
		log.DefaultLogger = log.NewCallbackLogger(stepper.LevelLog)
		setupUniverse(ctx, t, uu)
		return nil
	})

	stepper.PostStepHook(func(ctx context.Context, t flowtest.Asserter) error {
		t.Log("POST HOOK")
		uu.Outbox.AssertNoMessages(t)
		return nil
	})

	return stepper, uu
}

const TestVersion = "test-app-version"

// setupUniverse should only be called from the Setup callback, it is effectively
// a method but shouldn't show up there.
func setupUniverse(ctx context.Context, t flowtest.Asserter, uu *Universe) {
	t.Helper()

	conn := pgtest.GetTestDB(t, pgtest.WithDir("../../ext/db"))

	uu.Outbox = outboxtest.NewOutboxAsserter(t, conn)

	grpcPair := flowtest.NewGRPCPair(t, service.GRPCMiddleware(TestVersion)...)

	if err := service.RegisterGRPC(grpcPair.Server, conn, TestVersion); err != nil {
		t.Fatal(err.Error())
	}

	uu.GreetingCommand = test_spb.NewGreetingCommandServiceClient(grpcPair.Client)
	uu.GreetingQuery = test_spb.NewGreetingQueryServiceClient(grpcPair.Client)
	uu.TestTopic = test_tpb.NewTestTopicClient(grpcPair.Client)

	grpcPair.ServeUntilDone(t, ctx)
}
