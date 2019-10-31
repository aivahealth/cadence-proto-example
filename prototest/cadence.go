package prototest

import (
	"github.com/uber-go/tally"
	"go.uber.org/cadence/.gen/go/cadence/workflowserviceclient"
	"go.uber.org/cadence/client"
	"go.uber.org/cadence/encoded"
	"go.uber.org/cadence/worker"
	"go.uber.org/yarpc"
	"go.uber.org/yarpc/transport/tchannel"
	"go.uber.org/zap"
)

const (
	CadenceService = "cadence-frontend"
	ClientName     = "example-worker"
	TaskList       = "example-list"
)

func buildWorkflowServiceClient(endpoint string) workflowserviceclient.Interface {
	ch, err := tchannel.NewChannelTransport(tchannel.ServiceName(ClientName))
	if err != nil {
		panic("Failed to setup Cadence tchannel")
	}
	dispatcher := yarpc.NewDispatcher(yarpc.Config{
		Name: ClientName,
		Outbounds: yarpc.Outbounds{
			CadenceService: {Unary: ch.NewSingleOutbound(endpoint)},
		},
	})
	if err := dispatcher.Start(); err != nil {
		panic("Failed to start Cadence dispatcher")
	}

	return workflowserviceclient.New(dispatcher.ClientConfig(CadenceService))
}

func CadenceClient(endpoint, domain string) client.Client {
	clientOptions := &client.Options{
		DataConverter: NewProtoDataConverter(encoded.GetDefaultDataConverter()),
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
		MetricsScope:  tally.NewTestScope(TaskList, map[string]string{}),
		Logger:        zapLog,
		DataConverter: NewProtoDataConverter(encoded.GetDefaultDataConverter()),
	}

	w := worker.New(
		buildWorkflowServiceClient(endpoint),
		domain,
		TaskList,
		workerOptions,
	)
	err = w.Start()
	if err != nil {
		panic("Failed to start Cadence worker")
	}
}
