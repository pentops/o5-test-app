package service

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/pentops/o5-test-app/internal/gen/test/v1/test_pb"
	"github.com/pentops/o5-test-app/internal/gen/test/v1/test_spb"
	"github.com/pentops/o5-test-app/internal/state"
	"github.com/pentops/realms/j5auth"
	"github.com/pentops/sqrlx.go/sqrlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GreetingCommandService struct {
	db         sqrlx.Transactor
	appVersion string

	stateMachines *state.StateMachines
	test_spb.UnsafeGreetingCommandServiceServer
}

var _ test_spb.GreetingCommandServiceServer = &GreetingCommandService{}

func NewGreetingCommandService(db sqrlx.Transactor, version string, sm *state.StateMachines) (*GreetingCommandService, error) {

	return &GreetingCommandService{
		db:            db,
		appVersion:    version,
		stateMachines: sm,
	}, nil
}

func (ss *GreetingCommandService) RegisterGRPC(server grpc.ServiceRegistrar) {
	test_spb.RegisterGreetingCommandServiceServer(server, ss)
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
