package service

import (
	"context"
	"fmt"

	"github.com/pentops/o5-test-app/internal/gen/test/v1/test_pb"
	"github.com/pentops/o5-test-app/internal/gen/test/v1/test_tpb"
	"github.com/pentops/o5-test-app/internal/state"
	"github.com/pentops/realms/j5auth"
	"github.com/pentops/sqrlx.go/sqrlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type TestWorker struct {
	db            sqrlx.Transactor
	stateMachines *state.StateMachines
	test_tpb.UnsafeTestTopicServer
}

var _ test_tpb.TestTopicServer = &TestWorker{}

func NewTestWorker(db sqrlx.Transactor, ss *state.StateMachines) (*TestWorker, error) {
	return &TestWorker{
		db:            db,
		stateMachines: ss,
	}, nil
}

func (ww *TestWorker) RegisterGRPC(server grpc.ServiceRegistrar) {
	test_tpb.RegisterTestTopicServer(server, ww)
}

func (ww *TestWorker) Greeting(ctx context.Context, req *test_tpb.GreetingMessage) (*emptypb.Empty, error) {

	if req.WorkerError != nil {
		if req.WorkerError.Code == 0 {
			// while 0 means OK, if it is being set in an error that's not
			// useful, so we are using it for properly un-handled errors
			return nil, fmt.Errorf("TestError:%s", req.WorkerError.Message)
		}
		return nil, status.Error(codes.Code(req.WorkerError.Code), req.WorkerError.Message)
	}

	message, err := j5auth.GetMessageCause(ctx)
	if err != nil {
		return nil, err
	}

	evt := &test_pb.GreetingPSMEventSpec{
		Keys: &test_pb.GreetingKeys{
			GreetingId: req.GreetingId,
		},
		Message: message,
		Event: &test_pb.GreetingEventType_Replied{
			ReplyMessage: fmt.Sprintf("Hello %s", req.Name),
		},
	}

	_, err = ww.stateMachines.Greeting.Transition(ctx, ww.db, evt)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
