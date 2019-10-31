package main

import (
	"context"
	"github.com/aivahealth/cadence-proto-example/prototest"
	"log"
	"os"
	"time"

	"go.uber.org/cadence/client"
)

var (
	cadenceClient client.Client
)

func main() {
	ctx := context.Background()

	cadenceEndpoint := os.Getenv("CADENCE_CLI_ADDRESS")
	if cadenceEndpoint == "" {
		log.Fatal("Please set CADENCE_CLI_ADDRESS")
	}

	cadenceDomain := os.Getenv("CADENCE_CLI_DOMAIN")
	if cadenceDomain == "" {
		log.Fatal("Please set CADENCE_CLI_DOMAIN")
	}

	prototest.StartCadenceWorker(cadenceEndpoint, cadenceDomain)
	cadenceClient = prototest.CadenceClient(cadenceEndpoint, cadenceDomain)

	argIn := prototest.ExampleMsg{
		SimpleString: "Hello, world!",
		ComplexField: &prototest.ExampleMsg_SomeNumber{
			SomeNumber: 111222333444555,
		},
	}

	opts := client.StartWorkflowOptions{
		ExecutionStartToCloseTimeout: time.Second * 60,
		TaskList:                     prototest.TaskList,
	}
	run, err := cadenceClient.ExecuteWorkflow(
		ctx,
		opts,
		prototest.ExampleWorkflow,
		&argIn,
	)
	if err != nil {
		panic(err)
	}

	log.Printf("workflow=%q, run=%q", run.GetID(), run.GetRunID())

	err = run.Get(ctx, nil)
	if err != nil {
		log.Printf("Failure: %v", err.Error())
		//panic(err)
	}

	log.Println("done")
}
