package integration

import (
	"context"
	"testing"

	"github.com/pentops/flowtest"
	"github.com/pentops/o5-test-app/gen/test/v1/test_spb"
)

func TestService(t *testing.T) {

	flow := NewUniverse(t)
	defer flow.RunSteps(t)

	flow.StepC("Hello", func(ctx context.Context, t flowtest.Asserter) {
		res, err := flow.TestService.Hello(ctx, &test_spb.HelloRequest{Name: "World"})
		t.NoError(err)
		t.Equal("Hello, World", res.Message)
		t.Equal(TestVersion, res.AppVersion)
	})

}
