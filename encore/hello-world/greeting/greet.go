package greeting

import (
	"context"
	"log"

	"encore.app/greeting/workflow"
	"go.temporal.io/sdk/client"
)

type GreetingResponse struct {
	Greeting string
}

//encore:api public path=/greet/:name
func (s *Service) Greet(ctx context.Context, name string) (*GreetingResponse, error) {
	options := client.StartWorkflowOptions{
		ID: "greeting-workflow",
		TaskQueue: greetingTaskQueue,
	}

	we, err := s.client.ExecuteWorkflow(ctx, options, workflow.Greeting, name)
	if err != nil {
		return nil, err
	}
	log.Println("started workflow", "id", we.GetID(), "run_id", we.GetRunID())

	var greeting string
	err = we.Get(ctx, &greeting)
	if err != nil {
		return nil, err
	}
	return &GreetingResponse{
		Greeting: greeting,
	}, nil
}
