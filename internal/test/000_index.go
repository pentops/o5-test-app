package test

import (
	"context"

	"github.com/pentops/flowtest/runner"
	"github.com/pentops/o5-test-app/internal/test/universe"
)

func Run(ctx context.Context, cfg *universe.APIConfig, tags []string) error {
	testSet := runner.TestSet{}
	testSet.Register(1, "Greeting", universe.UniverseWrapper(cfg, GreetingTests))

	testSet.Register(2, "HandlerError", universe.UniverseWrapper(cfg, HandlerErrorTests))

	testSet.Register(3, "WorkerErrorTests", universe.UniverseWrapper(cfg, WorkerErrorTests), "name=worker-error")
	testSet.Register(4, "ReplayTests", universe.UniverseWrapper(cfg, ReplayTests), "name=replay")

	return testSet.Run(ctx, tags)
}
