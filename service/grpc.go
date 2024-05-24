package service

import (
	"context"

	"github.com/pentops/go-grpc-helpers/grpcerror"
	"github.com/pentops/go-grpc-helpers/protovalidatemw"
	"github.com/pentops/log.go/grpc_log"
	"github.com/pentops/log.go/log"
	"github.com/pentops/o5-go/auth/v1/auth_pb"
	"github.com/pentops/o5-test-app/gen/test/v1/test_spb"
	"github.com/pentops/o5-test-app/gen/test/v1/test_tpb"
	"github.com/pentops/o5-test-app/state"
	"github.com/pentops/protostate/gen/state/v1/psm_pb"
	"github.com/pentops/sqrlx.go/sqrlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		grpcerror.UnaryServerInterceptor(log.DefaultLogger),
		PSMActionMiddleware(actorExtractor),
		protovalidatemw.UnaryServerInterceptor(),
	}
}

func actorExtractor(ctx context.Context) *auth_pb.Actor {
	return &auth_pb.Actor{
		Type: &auth_pb.Actor_Named{
			Named: &auth_pb.Actor_NamedActor{
				Name: "Unauthenticated Client",
			},
		},
	}
}

type actionContextKey struct{}

type PSMAction struct {
	Method string
	Actor  *auth_pb.Actor
}

// PSMCause is a gRPC middleware that injects the PSM cause into t he context.
func PSMActionMiddleware(actorExtractor func(context.Context) *auth_pb.Actor) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		actor := actorExtractor(ctx)
		cause := PSMAction{
			Method: info.FullMethod,
			Actor:  actor,
		}
		ctx = context.WithValue(ctx, actionContextKey{}, cause)
		return handler(ctx, req)
	}
}

func WithPSMAction(ctx context.Context, action PSMAction) context.Context {
	return context.WithValue(ctx, actionContextKey{}, action)
}

func CommandCause(ctx context.Context) (*psm_pb.Cause, error) {

	cause, ok := ctx.Value(actionContextKey{}).(PSMAction)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "no actor")
	}

	return &psm_pb.Cause{
		Type: &psm_pb.Cause_Command{
			Command: &psm_pb.CommandCause{
				MethodName: cause.Method,
				Actor:      cause.Actor,
			},
		},
	}, nil
}
