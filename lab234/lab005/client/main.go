package main

import (
	"context"
	"fmt"
	exampletasks "github.com/RichardKnop/machinery/example/tasks"
	"github.com/RichardKnop/machinery/example/tracers"
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	opentracing_log "github.com/opentracing/opentracing-go/log"
	"log"
	"time"
)

func main() {
	send()
}

func send() error {
	cleanup, err := tracers.SetupTracer("sender")
	if err != nil {
		log.Fatalf("setupTracer error:%v", err)
	}
	defer cleanup()

	server, err := startServer()
	if err != nil {
		log.Fatalf("startServer error:%v", err)
	}

	var (
		addTask0 tasks.Signature
	)

	eta := time.Now().Add(time.Second * 10)
	var initTasks = func() {
		addTask0 = tasks.Signature{
			Name: "add",
			Args: []tasks.Arg{
				{
					Type:  "int64",
					Value: 1,
				},
				{
					Type:  "int64",
					Value: 1,
				},
			},
			ETA: &eta,
		}
	}

	span, ctx := opentracing.StartSpanFromContext(context.Background(), "send")
	defer span.Finish()

	batchID := uuid.New().String()
	span.SetBaggageItem("batch.id", batchID)
	span.LogFields(opentracing_log.String("batch.id", batchID))

	log.Println("Starting batch:", batchID)

	/*
	 * First, let's try sending a single task
	 */
	initTasks()
	log.Println("Single task:")

	asyncResult, err := server.SendTaskWithContext(ctx, &addTask0)
	if err != nil {
		return fmt.Errorf("Could not send task: %s", err.Error())
	} else {
		log.Printf("asyncResult:%+v", asyncResult)
	}

	results, err := asyncResult.Get(time.Duration(time.Millisecond * 5))
	if err != nil {
		return fmt.Errorf("Getting task result failed with error: %s", err.Error())
	}
	log.Printf("1 + 1 = %v\n", tasks.HumanReadableResults(results))

	return nil
}

func loadConfig() (*config.Config, error) {
	return config.NewFromYaml("./config.yml", true)
}

func startServer() (*machinery.Server, error) {
	cnf, err := loadConfig()
	if err != nil {
		return nil, err
	}

	// Create server instance
	server, err := machinery.NewServer(cnf)
	if err != nil {
		return nil, err
	}

	// Register tasks
	tasks := map[string]interface{}{
		"add":               exampletasks.Add,
		"multiply":          exampletasks.Multiply,
		"sum_ints":          exampletasks.SumInts,
		"sum_floats":        exampletasks.SumFloats,
		"concat":            exampletasks.Concat,
		"split":             exampletasks.Split,
		"panic_task":        exampletasks.PanicTask,
		"long_running_task": exampletasks.LongRunningTask,
	}

	return server, server.RegisterTasks(tasks)
}
