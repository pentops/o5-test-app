package state

import (
	"fmt"

	"github.com/pentops/o5-test-app/gen/test/v1/test_pb"
)

type StateMachines struct {
	Greeting *test_pb.GreetingPSM
}

func NewStateMachines() (*StateMachines, error) {
	greeting, err := NewGreetingPSM()
	if err != nil {
		return nil, fmt.Errorf("NewGreetingPSM: %w", err)
	}

	return &StateMachines{
		Greeting: greeting,
	}, nil

}
