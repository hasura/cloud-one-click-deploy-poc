package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"temporal-poc/app"

	temporalClient "go.temporal.io/sdk/client"

	"github.com/google/uuid"
	gql "github.com/hasura/go-graphql-client"
)

const (
	DataAdminSecret = "randomsecret"
	// Authorization   = "hasura-collaborator-token"
	// AuthorizationToken        = ""
	userRole                 = "admin"
	DataHost                 = "http://data.lux-dev.hasura.me"
	adminSecretHeader string = "x-hasura-admin-secret"
	XHasuraUserRole          = "x-hasura-role"
)

type hasuraTransport struct {
	apiKey string
	rt     http.RoundTripper
}

func (t *hasuraTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Set(XHasuraUserRole, userRole)
	r.Header.Set(adminSecretHeader, t.apiKey)
	// r.Header.Set("Hasura-Client-Name", "hasura-console")
	r.Header.Set("Content-Type", "application/json")
	return t.rt.RoundTrip(r)
}

func initClients() *gql.Client {

	httpClient := &http.Client{}
	httpClient.Transport = &hasuraTransport{apiKey: "randomsecret", rt: http.DefaultTransport}
	actionDataClient := gql.NewClient(fmt.Sprintf("%s/v1/graphql", DataHost), httpClient)
	return actionDataClient
}

// @@@SNIPSTART one-click-deployment-template-go-start-workflow
func main() {

	// Create the client object just once per process
	c, err := temporalClient.Dial(temporalClient.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	options := temporalClient.StartWorkflowOptions{
		ID:        "one-click-deployment",
		TaskQueue: app.OneClickDeploymentQueueName,
	}
	hasuraClient := app.HasuraClient{
		HasuraClient: initClients(),
	}

	input := app.HasuraDeployment{
		HasuraRepo: "https://github.com/rikinsk/hasura-one-click-deploy-sample-app.git",
		HasuraDir:  "hasura",
		UserID:     uuid.Must(uuid.Parse("4b58297f-14ca-4b53-a73e-72cbf15d3b11")),
	}

	log.Printf("Repo: %s Directory %s", input.HasuraRepo, input.HasuraDir)

	we, err := c.ExecuteWorkflow(context.Background(), options, hasuraClient.OneClickDeployment, input)
	if err != nil {
		log.Fatalln("unable to start the Workflow", err)
	}

	log.Printf("WorkflowID: %s RunID: %s\n", we.GetID(), we.GetRunID())

	var result string
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("unable to get Workflow result", err)
	}
	log.Println(result)
}

// @@@SNIPEND
