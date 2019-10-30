package main

import (
	"context"
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

	StartCadenceWorker(cadenceEndpoint, cadenceDomain)
	cadenceClient = CadenceClient(cadenceEndpoint, cadenceDomain)

	argIn := ExampleMsg{
		SimpleString: "Hello, world!",
		ComplexField: &ExampleMsg_SomeNumber{
			SomeNumber: 111222333444555,
		},
	}

	opts := client.StartWorkflowOptions{
		ExecutionStartToCloseTimeout: time.Second * 60,
		TaskList:                     taskList,
	}
	run, err := cadenceClient.ExecuteWorkflow(
		ctx,
		opts,
		ExampleWorkflow,
		&argIn,
	)
	if err != nil {
		panic(err)
	}

	log.Printf("workflow=%q, run=%q", run.GetID(), run.GetRunID())

	err = run.Get(ctx, nil)
	if err != nil {
		panic(err)
	}

	log.Println("done")
}
