package greeting

import (
	"context"
	"fmt"

	"encore.app/greeting/workflow"
	"encore.dev"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

var (
	envName           = encore.Meta().Environment.Name
	greetingTaskQueue = envName + "-greeting"
)

//encore:service
type Service struct {
	client client.Client
	worker worker.Worker
}

func initService() (*Service, error) {
	c, err := client.Dial(client.Options{
		HostPort: "127.0.0.1:7233",
		// ConnectionOptions: client.ConnectionOptions{
		// 	DialTimeout: 10 * time.Second,
		// },
	})
	if err != nil {
		return nil, fmt.Errorf("create temporal client: %v", err)
	}

	w := worker.New(c, greetingTaskQueue, worker.Options{})

	err = w.Start()
	if err != nil {
		c.Close()
		return nil, fmt.Errorf("start temporal worker: %v", err)
	}

	w.RegisterWorkflow(workflow.Greeting)
	w.RegisterActivity(workflow.ComposeGreeting)

	return &Service{client: c, worker: w}, nil
}

func (s *Service) Shutdown(force context.Context) {
	s.client.Close()
	s.worker.Stop()
}
