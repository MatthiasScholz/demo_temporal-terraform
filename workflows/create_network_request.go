package workflows

import (
	// "fmt"

	"go.temporal.io/sdk/workflow"
)

type CreateNetworkRequestWorkflowInput struct {
	NetworkName string
}

type CreateNetworkRequestWorkflowResult struct {
	NetworkName string
	NetworkCIDR string
}

func CreateNetworkRequestWorkflow(ctx workflow.Context, input *CreateNetworkRequestWorkflowInput) (*CreateNetworkRequestWorkflowResult, error) {
	return nil, nil
}
