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
		//Message: "Message from the client which cause a failure, since it does not match the expected message.",
		// Sucess:
		Message: "Expected Message",
	}

	temporalClient, err := client.NewClient(client.Options{
		Namespace: "default",
		HostPort:  "127.0.0.1:7233",
	})
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
