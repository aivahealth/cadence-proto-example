package prototest

import (
	"context"
	"fmt"
	"go.uber.org/cadence/activity"
	"time"

	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
)

func init() {
	workflow.Register(ExampleWorkflow)
	activity.Register(ExampleActivity)
}

func ExampleWorkflow(
	ctx workflow.Context,
	arg *ExampleMsg,
) error {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Second * 60,
		StartToCloseTimeout:    time.Second * 60,
	})

	log := workflow.GetLogger(ctx)
	log.Info("ExampleWorkflow", zap.Any("arg", arg))

	var result ExampleMsg
	err := workflow.ExecuteActivity(
		ctx,
		ExampleActivity,
		arg,
	).Get(ctx, &result)
	if err != nil {
		return err
	}

	log.Info("Got result from activity", zap.Any("result", result))

	return nil
}

func ExampleActivity(
	ctx context.Context,
	arg *ExampleMsg,
) (*ExampleMsg, error) {
	fmt.Printf("ExampleActivity arg: %#v\n", arg)

	result := ExampleMsg{
		SimpleString: "Different object!",
		ComplexField: &ExampleMsg_SomeString{
			SomeString: "Different concrete type!",
		},
	}

	return &result, nil
}
