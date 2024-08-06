package service

import (
	"github.com/pentops/go-grpc-helpers/protovalidatemw"
	"github.com/pentops/log.go/grpc_log"
	"github.com/pentops/log.go/log"
	"github.com/pentops/o5-auth/o5auth"
	"github.com/pentops/o5-test-app/internal/gen/test/v1/test_spb"
	"github.com/pentops/o5-test-app/internal/gen/test/v1/test_tpb"
	"github.com/pentops/o5-test-app/internal/state"
	"github.com/pentops/sqrlx.go/sqrlx"
	"google.golang.org/grpc"
)

func RegisterGRPC(server *grpc.Server, conn sqrlx.Connection, appVersion string) error {

	states, err := state.NewStateMachines()
	if err != nil {
		return err
	}

	commandService, err := NewGreetingCommandService(conn, appVersion, states)
	if err != nil {
		return err
	}
	test_spb.RegisterGreetingCommandServiceServer(server, commandService)

	testQueryService, err := NewGreetingQueryService(conn, states)
	if err != nil {
		return err
	}
	test_spb.RegisterGreetingQueryServiceServer(server, testQueryService)

	testWorker, err := NewTestWorker(conn, states)
	if err != nil {
		return err
	}
	test_tpb.RegisterTestTopicServer(server, testWorker)

	return nil
}

func GRPCMiddleware(version string) []grpc.UnaryServerInterceptor {
	return []grpc.UnaryServerInterceptor{
		grpc_log.UnaryServerInterceptor(log.DefaultContext, log.DefaultTrace, log.DefaultLogger),
		o5auth.GRPCMiddleware,
		protovalidatemw.UnaryServerInterceptor(),
	}
}
