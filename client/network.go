package main

import (
	"context"
	"log"

	"github.com/dynajoe/temporal-terraform-demo/workflows"
	"github.com/pborman/uuid"
	"go.temporal.io/sdk/client"
)

func main() {
	workflowID := "CreateNetworkRequestWorkflow_" + uuid.New()
	workflowOptions := client.StartWorkflowOptions{
		ID:        workflowID,
		TaskQueue: workflows.CreateNetworkRequestTaskQueue,
	}
	signal := workflows.CreateNetworkRequestSignal{
		// Failure:
		Message: "Message from the client",
		// Sucess:
		//Message: "SOME VALUE",
	}

	temporalClient, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
		return
	}
	input := workflows.CreateNetworkRequestWorkflowInput{NetworkName: "signal_network"}
	// NOTE: Sending a signal to an ALREADY running workflow
	// err := temporalClient.SignalWorkflow(
	// 	context.Background(),
	// 	"your-workflow-id",
	// 	"",
	// 	workflows.CreateNetworkRequestSignalChannelName,
	// 	signal)
	//  NOTE: Start a workflow and provide a signal
	run, err := temporalClient.SignalWithStartWorkflow(
		context.Background(),
		workflowID,
		workflows.CreateNetworkRequestSignalChannelName,
		signal,
		workflowOptions,
		workflows.CreateNetworkRequestWorkflow,
		&input,
	)
	if err != nil {
		log.Fatalln("Error sending the signal", err)
		return
	}

	log.Println("RunID: ", run.GetRunID())
	return
}