package service

import (
	"github.com/pentops/grpc.go/protovalidatemw"
	"github.com/pentops/log.go/grpc_log"
	"github.com/pentops/log.go/log"
	"github.com/pentops/o5-test-app/internal/state"
	"github.com/pentops/realms/j5auth"
	"github.com/pentops/sqrlx.go/sqrlx"
	"google.golang.org/grpc"
)

type App struct {
	GreetingCommand *GreetingCommandService
	GreetingQuery   *GreetingQueryService
	TestAccess      *TestAccess
	TestWorker      *TestWorker
	PublishWorker   *PublishWorker
}

func (aa *App) RegisterGRPC(server grpc.ServiceRegistrar) {
	aa.GreetingCommand.RegisterGRPC(server)
	aa.GreetingQuery.RegisterGRPC(server)
	aa.TestWorker.RegisterGRPC(server)
	aa.TestAccess.RegisterGRPC(server)
	aa.PublishWorker.RegisterGRPC(server)

}

func NewApp(db sqrlx.Transactor, appVersion string) (*App, error) {
	states, err := state.NewStateMachines()
	if err != nil {
		return nil, err
	}

	commandService, err := NewGreetingCommandService(db, appVersion, states)
	if err != nil {
		return nil, err
	}

	testQueryService, err := NewGreetingQueryService(db, states)
	if err != nil {
		return nil, err
	}

	testWorker, err := NewTestWorker(db, states)
	if err != nil {
		return nil, err
	}

	testAccess, err := NewTestAccess(db)
	if err != nil {
		return nil, err
	}

	publishWorker, err := NewPublishWorker(db)
	if err != nil {
		return nil, err
	}

	return &App{
		GreetingCommand: commandService,
		GreetingQuery:   testQueryService,
		TestWorker:      testWorker,
		TestAccess:      testAccess,
		PublishWorker:   publishWorker,
	}, nil
}

func GRPCMiddleware(version string) []grpc.UnaryServerInterceptor {
	return []grpc.UnaryServerInterceptor{
		grpc_log.UnaryServerInterceptor(log.DefaultContext, log.DefaultTrace, log.DefaultLogger),
		j5auth.GRPCMiddleware,
		protovalidatemw.UnaryServerInterceptor(),
	}
}
