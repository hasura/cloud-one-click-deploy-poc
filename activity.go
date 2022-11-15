package app

import (
	"context"
	"log"

	gql "github.com/hasura/go-graphql-client"
)

type HasuraClient struct {
	HasuraClient *gql.Client
}

// @@@SNIPSTART one-click-deployment-go-create-project
func (hasuraClient *HasuraClient) CreateProject(ctx context.Context, data HasuraDeployment) (OneClickDeploymentOutput, error) {
	log.Printf(
		"create Hasura project for sample app repo: $%s directory %s\n\n",
		data.HasuraRepo,
		data.HasuraDir,
	)
	OneClickDeployment := &OneClickDeploymentHasura{
		hasuraClient: hasuraClient,
	}
	status, err := OneClickDeployment.CreateProject(ctx, data)
	return status, err
}

// @@@SNIPEND

// @@@SNIPSTART one-click-deployment-go-health-check-repo
func (hasuraClient *HasuraClient) HealthCheck(ctx context.Context, data HasuraDeployment) (OneClickDeploymentOutput, error) {

	//log health check
	log.Print("checking project health")
	status := OneClickDeploymentOutput{
		Status: "success",
	}
	return status, nil
}

// @@@SNIPEND

// @@@SNIPSTART one-click-deployment-go-clone-repo
func (hasuraClient *HasuraClient) CloneRepo(ctx context.Context, data HasuraDeployment) (OneClickDeploymentOutput, error) {

	OneClickDeployment := &OneClickDeploymentHasura{
		hasuraClient: hasuraClient,
	}
	status, err := OneClickDeployment.CloneRepo(ctx, data)
	return status, err
}

// @@@SNIPEND

// @@@SNIPSTART
func (hasuraClient *HasuraClient) ApplyMetadata(ctx context.Context, data HasuraDeployment) (OneClickDeploymentOutput, error) {

	//log health check
	log.Print("Applying metadata")
	status := OneClickDeploymentOutput{
		Status: "success",
	}
	return status, nil
}

// @@@SNIPEND

// @@@SNIPSTART
func (hasuraClient *HasuraClient) ApplyMigration(ctx context.Context, data HasuraDeployment) (OneClickDeploymentOutput, error) {

	//log health check
	log.Print("Applying migration")
	status := OneClickDeploymentOutput{
		Status: "success",
	}
	return status, nil
}

// @@@SNIPEND
