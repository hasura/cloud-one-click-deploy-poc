package app

import (
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

type OneClickDeploymentOutput struct {
	Status string `json:"status"`
}

// @@@SNIPSTART one-click-deployment-template-go-workflow
func (hasuraClient *HasuraClient) OneClickDeployment(ctx workflow.Context, input HasuraDeployment) (string, error) {
	// RetryPolicy specifies how to automatically handle retries if an Activity fails.
	retrypolicy := &temporal.RetryPolicy{
		InitialInterval:    time.Second,
		BackoffCoefficient: 2.0,
		MaximumInterval:    time.Minute,
		MaximumAttempts:    500,
	}
	options := workflow.ActivityOptions{
		// Timeout options specify when to automatically timeout Activity functions.
		StartToCloseTimeout: time.Minute,
		// Optionally provide a customized RetryPolicy.
		// Temporal retries failures by default, this is just an example.
		RetryPolicy: retrypolicy,
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	var output OneClickDeploymentOutput
	err := workflow.ExecuteActivity(ctx, hasuraClient.CreateProject, input).Get(ctx, &output)
	if err != nil {
		return "", err
	}

	err = workflow.ExecuteActivity(ctx, hasuraClient.HealthCheck, input).Get(ctx, &output)
	if err != nil {
		return "", err
	}

	err = workflow.ExecuteActivity(ctx, hasuraClient.CloneRepo, input).Get(ctx, &output)
	if err != nil {
		return "", err
	}

	err = workflow.ExecuteActivity(ctx, hasuraClient.ApplyMetadata, input).Get(ctx, &output)
	if err != nil {
		return "", err
	}

	err = workflow.ExecuteActivity(ctx, hasuraClient.ApplyMigration, input).Get(ctx, &output)
	if err != nil {
		return "", err
	}

	result := "deployment completed"
	return result, nil
}

// @@@SNIPEND
