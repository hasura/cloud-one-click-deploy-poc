package app

import (
	"github.com/google/uuid"
)

// @@@SNIPSTART one-app-deployment-template-go-shared-task-queue
const OneClickDeploymentQueueName = "ONE_CLICK_DEPLOYMENT_TASK_QUEUE"

// @@@SNIPEND

// this will be used to share values across the activities
// @@@SNIPSTART one-app-deployment-template-go
type HasuraDeployment struct {
	HasuraRepo string
	HasuraDir  string
	UserID     uuid.UUID
}

// @@@SNIPEND
