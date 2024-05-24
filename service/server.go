package service

import (
	"context"

	"github.com/pentops/o5-test-app/gen/test/v1/test_spb"
)

type TestService struct {
	*test_spb.UnimplementedTestServiceServer
}

func NewTestService() (*TestService, error) {
	return &TestService{}, nil
}

func (s *TestService) Hello(ctx context.Context, req *test_spb.HelloRequest) (*test_spb.HelloResponse, error) {
	return &test_spb.HelloResponse{
		Message:    "Hello, " + req.Name,
		AppVersion: "test-app-version",
	}, nil
}
