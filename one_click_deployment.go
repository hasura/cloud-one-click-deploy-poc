package app

import (
	"context"

	"github.com/google/uuid"
)

type OneClickDeploymentHasura struct {
	dir          string
	projectID    uuid.UUID
	commit_id    string
	status       string
	err          error
	hasuraClient *HasuraClient
}

func (client *OneClickDeploymentHasura) CreateProject(ctx context.Context, data HasuraDeployment) (OneClickDeploymentOutput, error) {

	var mutation struct {
		CreateTenant struct {
			Tenant struct {
				Project struct {
					ID uuid.UUID `json:"id" graphql:"id"`
				} `json:"project" graphql:"project"`
			} `json:"tenant" graphql:"tenant"`
		} `graphql:"createTenant(cloud: $cloud, region: $region, plan: $plan)"`
	}

	type String string

	variable := map[string]interface{}{
		"cloud":  String("aws"),
		"region": String("us-east-2"),
		"plan":   String("cloud_free"),
	}

	err := client.hasuraClient.HasuraClient.NamedMutate(ctx, "CreateTenant", &mutation, variable)
	if err != nil {
		return OneClickDeploymentOutput{
			Status: "failed",
		}, err
	}

	client.projectID = mutation.CreateTenant.Tenant.Project.ID

	return OneClickDeploymentOutput{
		Status: "success",
	}, nil
}

func (client *OneClickDeploymentHasura) CloneRepo(ctx context.Context, data HasuraDeployment) (OneClickDeploymentOutput, error) {
	return OneClickDeploymentOutput{
		Status: "success",
	}, nil
}
