package service

import (
	"github.com/pentops/o5-test-app/internal/gen/test/v1/test_spb"
	"github.com/pentops/o5-test-app/internal/state"
	"github.com/pentops/protostate/psm"
	"github.com/pentops/sqrlx.go/sqrlx"
	"google.golang.org/grpc"
)

type GreetingQueryService struct {
	db sqrlx.Transactor

	querySet *test_spb.GreetingPSMQuerySet
	test_spb.UnsafeGreetingQueryServiceServer

	*test_spb.GreetingQueryServiceImpl
}

var _ test_spb.GreetingQueryServiceServer = &GreetingQueryService{}

func NewGreetingQueryService(db sqrlx.Transactor, states *state.StateMachines) (*GreetingQueryService, error) {

	querySpec := test_spb.DefaultGreetingPSMQuerySpec(states.Greeting.StateTableSpec())
	querySet, err := test_spb.NewGreetingPSMQuerySet(querySpec, psm.StateQueryOptions{})
	if err != nil {
		return nil, err
	}

	impl := test_spb.NewGreetingQueryServiceImpl(db, querySet)

	return &GreetingQueryService{
		db:                       db,
		querySet:                 querySet,
		GreetingQueryServiceImpl: impl,
	}, nil
}

func (ds *GreetingQueryService) RegisterGRPC(server grpc.ServiceRegistrar) {
	test_spb.RegisterGreetingQueryServiceServer(server, ds)
}
