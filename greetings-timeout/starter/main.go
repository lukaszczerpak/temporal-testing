package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/pborman/uuid"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"

	"github.com/temporalio/samples-go/greetings-timeout"
)

func main() {
	// The client is a heavyweight object that should be created once per process.
	c, err := client.Dial(client.Options{
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	workflowOptions := client.StartWorkflowOptions{
		ID:                 "greetings_" + uuid.New(),
		TaskQueue:          "greetings",
		WorkflowRunTimeout: 10 * time.Second,
	}

	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, greetings.GreetingSample)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
		return
	}
	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())
	var s string
	err = we.Get(context.Background(), &s)
	if err != nil {
		out := fmt.Sprintf("isTimeout=%v", temporal.IsTimeoutError(err))
		log.Println(out)
		log.Fatalln("Unable to get return value from workflow", err)
	}
}
