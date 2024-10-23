package app

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

func TestWorkflow(ctx workflow.Context, name string) (string, error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}

	ctx = workflow.WithActivityOptions(ctx, options)

	var result string
	err := workflow.ExecuteActivity(ctx, ComposeTest, name).Get(ctx, &result)

	return result, err
}
