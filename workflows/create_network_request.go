package workflows

import (
	// "fmt"

	"errors"

	"go.temporal.io/sdk/workflow"
)

type CreateNetworkRequestWorkflowInput struct {
	NetworkName string
}

type CreateNetworkRequestWorkflowResult struct {
	NetworkName string
	NetworkCIDR string
}

const CreateNetworkRequestTaskQueue = "CreateNetworkRequestTaskQueue"
const CreateNetworkRequestSignalChannelName = "CreateNetworkRequestSignalChannelName"

type CreateNetworkRequestSignal struct {
	Message string
}

func CreateNetworkRequestWorkflow(ctx workflow.Context, input *CreateNetworkRequestWorkflowInput) (*CreateNetworkRequestWorkflowResult, error) {
	var signal CreateNetworkRequestSignal
	signalChan := workflow.GetSignalChannel(ctx, CreateNetworkRequestSignalChannelName)
	selector := workflow.NewSelector(ctx)
	selector.AddReceive(signalChan, func(channel workflow.ReceiveChannel, more bool) {
		channel.Receive(ctx, &signal)
	})
	selector.Select(ctx)

	logger := workflow.GetLogger(ctx)
	logger.Info("Recieved message: ", signal.Message)

	// Some dummy message condition
	if len(signal.Message) > 0 && signal.Message != "SOME VALUE" {
		logger.Error("Invalid value for message: ", signal.Message)
		return nil, errors.New("signal")
	}

	logger.Info("Workflow successfully processed: ", signal.Message)
	return nil, nil
}
