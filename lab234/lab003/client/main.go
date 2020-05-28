package main

import (
	"context"
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	opentracing_log "github.com/opentracing/opentracing-go/log"
	"log"
	"time"
)

var (
	DefaultQueue = "machinery_tasks"
)

func main() {
	send()
}

func send() {
	//cleanup, err := tracers.SetupTracer("sender")
	//if err != nil {
	//	log.Fatalf("setupTracer error:%v", err)
	//}
	//defer cleanup()

	server, err := startServer()
	if err != nil {
		log.Fatalf("startServer error:%v", err)
	}

	/*
	 * Lets start a span representing this run of the `send` command and
	 * set a batch id as baggage so it can travel all the way into
	 * the worker functions.
	 */
	span, ctx := opentracing.StartSpanFromContext(context.Background(), "send")
	defer span.Finish()

	batchID := uuid.New().String()
	span.SetBaggageItem("batch.id", batchID)
	span.LogFields(opentracing_log.String("batch.id", batchID))

	log.Println("Starting batch:", batchID)

	/*
	 * First, let's try sending a single task
	 */

	log.Println("Single task:")

	testTask001 := tasks.Signature{
		Name: "test001",
		Args: []tasks.Arg{
			{
				Type:  "string",
				Value: "abc",
			},
			{
				Type:  "int",
				Value: 456,
			},
		},
	}

	asyncResult, err := server.SendTaskWithContext(ctx, &testTask001)
	if err != nil {
		log.Fatalf("sent task error:%v", err)
	} else {
		log.Printf("asyncResult:%+v", asyncResult)
	}

	results, err := asyncResult.Get(time.Duration(time.Millisecond * 5))
	if err != nil {
		log.Fatalf("result get error:%v", err)
	}
	log.Printf("1 + 1 = %v\n", tasks.HumanReadableResults(results))

	//连发test002
	//for i := 0; i < 10; i++ {
	//	testTask002 := tasks.Signature{
	//		Name: "test002",
	//		Args: []tasks.Arg{
	//			{
	//				Type:  "int",
	//				Value: 1,
	//			},
	//		},
	//	}
	//	if _, err := server.SendTask(&testTask002); err != nil {
	//		log.Fatalf("send task 002 error:%v", err)
	//	}
	//
	//	time.Sleep(time.Millisecond * 200)
	//}

	//实验retry
	for i := 0; i < 3; i++ {
		testTask003 := tasks.Signature{
			Name: "test003",
			Args: []tasks.Arg{
				{
					Type:  "int",
					Value: i,
				},
			},
			RetryCount: 5,
		}
		if _, err := server.SendTask(&testTask003); err != nil {
			log.Fatalf("send task 003 error:%v", err)
		}

		time.Sleep(time.Millisecond * 200)
	}
}

func startServer() (*machinery.Server, error) {
	var cnf = &config.Config{
		Broker:          "redis://localhost:6379/3",
		DefaultQueue:    DefaultQueue,
		ResultBackend:   "redis://localhost:6379/4",
		ResultsExpireIn: 3600,
	}

	// Create server instance
	server, err := machinery.NewServer(cnf)
	if err != nil {
		return nil, err
	}

	return server, nil
}
