package service

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/pentops/j5/gen/j5/state/v1/psm_j5pb"
	"github.com/pentops/o5-test-app/internal/gen/test/v1/test_pb"
	"github.com/pentops/o5-test-app/internal/gen/test/v1/test_tpb"
	"github.com/pentops/o5-test-app/internal/state"
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

var replyNamespace = uuid.MustParse("7B4D4FB7-28BA-4848-9EE3-4C3B0B2263E6")

func replyID(greetingID string) string {
	return uuid.NewSHA1(replyNamespace, []byte(greetingID)).String()
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

	// TODO: Greeting should reply to a reply topic with the reply, but for now
	// we are just going directly to the state machine.

	evt := &test_pb.GreetingPSMEventSpec{
		Keys: &test_pb.GreetingKeys{
			GreetingId: req.GreetingId,
		},
		EventID:   replyID(req.GreetingId),
		Timestamp: time.Now(),
		Cause: &psm_j5pb.Cause{
			Type: &psm_j5pb.Cause_ExternalEvent{
				ExternalEvent: &psm_j5pb.ExternalEventCause{
					SystemName: "test",
					EventName:  "greeting",
				},
			},
		},

		Event: &test_pb.GreetingEventType_Replied{
			ReplyMessage: fmt.Sprintf("Hello %s", req.Name),
		},
	}

	_, err := ww.stateMachines.Greeting.Transition(ctx, ww.db, evt)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
