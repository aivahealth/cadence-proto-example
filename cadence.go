package main

import (
	"github.com/uber-go/tally"
	"go.uber.org/cadence/.gen/go/cadence/workflowserviceclient"
	"go.uber.org/cadence/client"
	"go.uber.org/cadence/worker"
	"go.uber.org/yarpc"
	"go.uber.org/yarpc/transport/tchannel"
	"go.uber.org/zap"
)

const (
	cadenceService = "cadence-frontend"
	clientName     = "example-worker"
	taskList       = "example-list"
)

func buildWorkflowServiceClient(endpoint string) workflowserviceclient.Interface {
	ch, err := tchannel.NewChannelTransport(tchannel.ServiceName(clientName))
	if err != nil {
		panic("Failed to setup Cadence tchannel")
	}
	dispatcher := yarpc.NewDispatcher(yarpc.Config{
		Name: clientName,
		Outbounds: yarpc.Outbounds{
			cadenceService: {Unary: ch.NewSingleOutbound(endpoint)},
		},
	})
	if err := dispatcher.Start(); err != nil {
		panic("Failed to start Cadence dispatcher")
	}

	return workflowserviceclient.New(dispatcher.ClientConfig(cadenceService))
}

func CadenceClient(endpoint, domain string) client.Client {
	clientOptions := &client.Options{
		DataConverter: &CustomDataConverter{},
	}
	return client.NewClient(
		buildWorkflowServiceClient(endpoint), domain, clientOptions,
	)
}

func StartCadenceWorker(endpoint, domain string) {
	zapLog, err := zap.NewDevelopment(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}

	workerOptions := worker.Options{
		MetricsScope:  tally.NewTestScope(taskList, map[string]string{}),
		Logger:        zapLog,
		DataConverter: &CustomDataConverter{},
	}

	w := worker.New(
		buildWorkflowServiceClient(endpoint),
		domain,
		taskList,
		workerOptions,
	)
	err = w.Start()
	if err != nil {
		panic("Failed to start Cadence worker")
	}
}
