package integration

import (
	"context"
	"testing"

	"github.com/pentops/flowtest"
	"github.com/pentops/o5-test-app/gen/test/v1/test_spb"
	"github.com/pentops/o5-test-app/service"
	"gopkg.daemonl.com/log"
)

type Universe struct {
	//Outbox *outboxtest.OutboxAsserter
	*flowtest.Stepper[*testing.T]

	TestService test_spb.TestServiceClient
}

func NewUniverse(t *testing.T) *Universe {
	name := t.Name()
	stepper := flowtest.NewStepper[*testing.T](name)
	uu := &Universe{
		Stepper: stepper,
	}
	return uu
}

const TestVersion = "test-app-version"

func (uu *Universe) RunSteps(t *testing.T) {
	t.Helper()

	ctx := context.Background()

	log.DefaultLogger = log.NewCallbackLogger(uu.Stepper.Log)

	grpcPair := flowtest.NewGRPCPair(t, service.GRPCMiddleware(TestVersion)...)

	if err := service.RegisterGRPC(grpcPair.Server); err != nil {
		t.Fatal(err.Error())
	}

	uu.TestService = test_spb.NewTestServiceClient(grpcPair.Client)

	grpcPair.ServeUntilDone(t, ctx)

	uu.Stepper.RunSteps(t)
}
