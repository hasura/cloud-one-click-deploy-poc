package main

import (
	"fmt"
	"log"
	"net/http"

	"temporal-poc/app"

	gql "github.com/hasura/go-graphql-client"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

const (
	DataAdminSecret          = "randomsecret"
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
	r.Header.Set("Content-Type", "application/json")
	return t.rt.RoundTrip(r)
}

func initClients() *gql.Client {

	httpClient := &http.Client{}
	httpClient.Transport = &hasuraTransport{apiKey: "randomsecret", rt: http.DefaultTransport}
	actionDataClient := gql.NewClient(fmt.Sprintf("%s/v1/graphql", DataHost), httpClient)
	return actionDataClient
}

// @@@SNIPSTART one-click-deployment-template-go-worker
func main() {

	hasuraClient := app.HasuraClient{
		HasuraClient: initClients(),
	}

	// Create the client object just once per process
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	w := worker.New(c, app.OneClickDeploymentQueueName, worker.Options{})

	// This worker hosts both Workflow and Activity functions
	w.RegisterWorkflow(hasuraClient.OneClickDeployment)
	w.RegisterActivity(hasuraClient.CreateProject)
	w.RegisterActivity(hasuraClient.HealthCheck)
	w.RegisterActivity(hasuraClient.CloneRepo)
	w.RegisterActivity(hasuraClient.ApplyMetadata)
	w.RegisterActivity(hasuraClient.ApplyMigration)

	// Start listening to the Task Queue
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}

// @@@SNIPEND
