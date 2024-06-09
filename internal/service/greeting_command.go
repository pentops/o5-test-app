package service

import (
	"context"
	"time"

	sq "github.com/elgris/sqrl"
	"github.com/google/uuid"
	"github.com/pentops/o5-test-app/internal/gen/test/v1/test_pb"
	"github.com/pentops/o5-test-app/internal/gen/test/v1/test_spb"
	"github.com/pentops/o5-test-app/internal/state"
	"github.com/pentops/sqrlx.go/sqrlx"
)

type GreetingCommandService struct {
	db         *sqrlx.Wrapper
	appVersion string

	stateMachines *state.StateMachines
	*test_spb.UnimplementedGreetingCommandServiceServer
}

func NewGreetingCommandService(conn sqrlx.Connection, version string, sm *state.StateMachines) (*GreetingCommandService, error) {
	db, err := sqrlx.New(conn, sq.Dollar)
	if err != nil {
		return nil, err
	}

	return &GreetingCommandService{
		db:            db,
		appVersion:    version,
		stateMachines: sm,
	}, nil
}

func (ss *GreetingCommandService) Hello(ctx context.Context, req *test_spb.HelloRequest) (*test_spb.HelloResponse, error) {

	cause, err := CommandCause(ctx)
	if err != nil {
		return nil, err
	}

	evt := &test_pb.GreetingPSMEventSpec{
		Keys: &test_pb.GreetingKeys{
			GreetingId: req.GreetingId,
		},
		EventID:   uuid.NewString(),
		Timestamp: time.Now(),
		Cause:     cause,
		Event: &test_pb.GreetingEventType_Initiated{
			Name:       req.Name,
			AppVersion: ss.appVersion,
		},
	}

	newState, err := ss.stateMachines.Greeting.Transition(ctx, ss.db, evt)
	if err != nil {
		return nil, err
	}

	return &test_spb.HelloResponse{
		Greeting: newState,
	}, nil
}
