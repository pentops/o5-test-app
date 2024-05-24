package service

import (
	"context"

	sq "github.com/elgris/sqrl"
	"github.com/pentops/o5-test-app/gen/test/v1/test_spb"
	"github.com/pentops/o5-test-app/state"
	"github.com/pentops/protostate/psm"
	"github.com/pentops/sqrlx.go/sqrlx"
)

type GreetingQueryService struct {
	db *sqrlx.Wrapper

	querySet *test_spb.GreetingPSMQuerySet
	*test_spb.UnimplementedGreetingQueryServiceServer
}

func NewGreetingQueryService(conn sqrlx.Connection, states *state.StateMachines) (*GreetingQueryService, error) {
	db, err := sqrlx.New(conn, sq.Dollar)
	if err != nil {
		return nil, err

	}

	querySpec := test_spb.DefaultGreetingPSMQuerySpec(states.Greeting.StateTableSpec())
	querySet, err := test_spb.NewGreetingPSMQuerySet(querySpec, psm.StateQueryOptions{})
	if err != nil {
		return nil, err
	}

	return &GreetingQueryService{
		db:       db,
		querySet: querySet,
	}, nil
}

func (ds *GreetingQueryService) ListGreetingEvents(ctx context.Context, req *test_spb.ListGreetingEventsRequest) (*test_spb.ListGreetingEventsResponse, error) {
	res := &test_spb.ListGreetingEventsResponse{}

	return res, ds.querySet.ListEvents(ctx, ds.db, req, res)
}

func (ds *GreetingQueryService) GetGreeting(ctx context.Context, req *test_spb.GetGreetingRequest) (*test_spb.GetGreetingResponse, error) {
	res := &test_spb.GetGreetingResponse{}

	return res, ds.querySet.Get(ctx, ds.db, req, res)
}

func (ds *GreetingQueryService) ListGreetings(ctx context.Context, req *test_spb.ListGreetingsRequest) (*test_spb.ListGreetingsResponse, error) {
	res := &test_spb.ListGreetingsResponse{}

	return res, ds.querySet.List(ctx, ds.db, req, res)
}
