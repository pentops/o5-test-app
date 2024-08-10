package service

import (
	"context"
	"fmt"
	"time"

	sq "github.com/elgris/sqrl"
	"github.com/google/uuid"
	"github.com/pentops/o5-test-app/internal/gen/test/v1/test_pb"
	"github.com/pentops/o5-test-app/internal/gen/test/v1/test_spb"
	"github.com/pentops/o5-test-app/internal/state"
	"github.com/pentops/realms/j5auth"
	"github.com/pentops/sqrlx.go/sqrlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

	action, err := j5auth.GetAuthenticatedAction(ctx)
	if err != nil {
		return nil, err
	}

	if req.ThrowError != nil {
		if req.ThrowError.Code == 0 {
			// while 0 means OK, if it is being set in an error that's not
			// useful, so we are using it for properly un-handled errors
			return nil, fmt.Errorf("TestError:%s", req.ThrowError.Message)
		}
		return nil, status.Error(codes.Code(req.ThrowError.Code), req.ThrowError.Message)
	}

	evt := &test_pb.GreetingPSMEventSpec{
		Keys: &test_pb.GreetingKeys{
			GreetingId: req.GreetingId,
		},
		EventID:   uuid.NewString(),
		Timestamp: time.Now(),
		Action:    action,
		Event: &test_pb.GreetingEventType_Initiated{
			Name:        req.Name,
			AppVersion:  ss.appVersion,
			WorkerError: req.WorkerError,
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
