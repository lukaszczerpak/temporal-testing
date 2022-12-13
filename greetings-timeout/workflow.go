package greetings

import (
	"fmt"
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

// GreetingSample workflow definition.
// This greetings sample workflow executes 3 activities in sequential.
// It gets greeting and name from 2 different activities,
// and then pass greeting and name as input to a 3rd activity to generate final greetings.
func GreetingSample(ctx workflow.Context) (wret string, werr error) {
	logger := workflow.GetLogger(ctx)

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    1 * time.Second,
			BackoffCoefficient: 1.0,
			MaximumAttempts:    0,
		},
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	// @@@SNIPSTART samples-go-dependency-sharing-workflow
	var a *Activities // use a nil struct pointer to call activities that are part of a structure

	var greetResult string
	err := workflow.ExecuteActivity(ctx, a.GetGreeting).Get(ctx, &greetResult)
	if err != nil {
		logger.Error("Get greeting failed.", "Error", err)
		return "", err
	}
	// @@@SNIPEND

	// Get Name.
	var nameResult string
	err = workflow.ExecuteActivity(ctx, a.GetName).Get(ctx, &nameResult)
	if err != nil {
		logger.Error("Get name failed.", "Error", err)
		return "", err
	}

	// Say Greeting.
	var sayResult string
	err = workflow.ExecuteActivity(ctx, a.SayGreeting, greetResult, nameResult).Get(ctx, &sayResult)
	if err != nil {
		out := fmt.Sprintf("isTimeout=%v,ctx.Err=%v", temporal.IsTimeoutError(err), ctx.Err())
		logger.Info(out)
		logger.Error("Marshalling failed with error.", "Error", err)
		return "", err
	}

	logger.Info("GreetingSample completed.", "Result", sayResult)
	return sayResult, nil
}
