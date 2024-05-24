package service

import (
	"context"
	"fmt"
	"time"

	sq "github.com/elgris/sqrl"
	"github.com/google/uuid"
	"github.com/pentops/o5-test-app/gen/test/v1/test_pb"
	"github.com/pentops/o5-test-app/gen/test/v1/test_tpb"
	"github.com/pentops/o5-test-app/state"
	"github.com/pentops/protostate/gen/state/v1/psm_pb"
	"github.com/pentops/sqrlx.go/sqrlx"
	"google.golang.org/protobuf/types/known/emptypb"
)

type TestWorker struct {
	db            *sqrlx.Wrapper
	stateMachines *state.StateMachines
	*test_tpb.UnimplementedTestTopicServer
}

func NewTestWorker(conn sqrlx.Connection, ss *state.StateMachines) (*TestWorker, error) {

	db, err := sqrlx.New(conn, sq.Dollar)
	if err != nil {
		return nil, err
	}

	return &TestWorker{
		db:            db,
		stateMachines: ss,
	}, nil
}

func (ww *TestWorker) Greeting(ctx context.Context, req *test_tpb.GreetingMessage) (*emptypb.Empty, error) {

	// TODO: Greeting should reply to a reply topic with the reply, but for now
	// we are just going directly to the state machine.

	evt := &test_pb.GreetingPSMEventSpec{
		Keys: &test_pb.GreetingKeys{
			GreetingId: req.GreetingId,
		},
		EventID:   uuid.NewString(),
		Timestamp: time.Now(),
		Cause: &psm_pb.Cause{
			Type: &psm_pb.Cause_ExternalEvent{
				ExternalEvent: &psm_pb.ExternalEventCause{
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
