package main

import (
	"context"
	"errors"
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
		addTask0, addTask1, addTask2                      tasks.Signature
		multiplyTask0, multiplyTask1                      tasks.Signature
		sumIntsTask, sumFloatsTask, concatTask, splitTask tasks.Signature
		panicTask                                         tasks.Signature
		longRunningTask                                   tasks.Signature
	)

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
		}

		addTask1 = tasks.Signature{
			Name: "add",
			Args: []tasks.Arg{
				{
					Type:  "int64",
					Value: 2,
				},
				{
					Type:  "int64",
					Value: 2,
				},
			},
		}

		addTask2 = tasks.Signature{
			Name: "add",
			Args: []tasks.Arg{
				{
					Type:  "int64",
					Value: 5,
				},
				{
					Type:  "int64",
					Value: 6,
				},
			},
		}

		multiplyTask0 = tasks.Signature{
			Name: "multiply",
			Args: []tasks.Arg{
				{
					Type:  "int64",
					Value: 4,
				},
			},
		}

		multiplyTask1 = tasks.Signature{
			Name: "multiply",
		}

		sumIntsTask = tasks.Signature{
			Name: "sum_ints",
			Args: []tasks.Arg{
				{
					Type:  "[]int64",
					Value: []int64{1, 2},
				},
			},
		}

		sumFloatsTask = tasks.Signature{
			Name: "sum_floats",
			Args: []tasks.Arg{
				{
					Type:  "[]float64",
					Value: []float64{1.5, 2.7},
				},
			},
		}

		concatTask = tasks.Signature{
			Name: "concat",
			Args: []tasks.Arg{
				{
					Type:  "[]string",
					Value: []string{"foo", "bar"},
				},
			},
		}

		splitTask = tasks.Signature{
			Name: "split",
			Args: []tasks.Arg{
				{
					Type:  "string",
					Value: "foo",
				},
			},
		}

		panicTask = tasks.Signature{
			Name: "panic_task",
		}

		longRunningTask = tasks.Signature{
			Name: "long_running_task",
		}
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

	// Now let's try a parallel execution
	initTasks()
	log.Println("Group of tasks (parallel execution):")

	group, err := tasks.NewGroup(&addTask0, &addTask1, &addTask2)
	if err != nil {
		return fmt.Errorf("Error creating group: %s", err.Error())
	} else {
		log.Printf("group:%+v", group)
	}

	asyncResults, err := server.SendGroupWithContext(ctx, group, 10)
	if err != nil {
		return fmt.Errorf("Could not send group: %s", err.Error())
	} else {
		log.Printf("asyncResult:%+v", asyncResults)
	}

	for n, asyncResult := range asyncResults {
		results, err = asyncResult.Get(time.Duration(time.Millisecond * 5))
		if err != nil {
			return fmt.Errorf("Getting task result failed with error: %s", err.Error())
		}
		log.Printf(
			"groupResult[%d]:%v + %v = %v\n", n,
			asyncResult.Signature.Args[0].Value,
			asyncResult.Signature.Args[1].Value,
			tasks.HumanReadableResults(results),
		)
	}

	// Now let's try a group with a chord
	initTasks()
	log.Println("Group of tasks with a callback (chord):")

	group, err = tasks.NewGroup(&addTask0, &addTask1, &addTask2)
	if err != nil {
		return fmt.Errorf("Error creating group: %s", err.Error())
	}

	chord, err := tasks.NewChord(group, &multiplyTask1)
	if err != nil {
		return fmt.Errorf("Error creating chord: %s", err)
	}

	chordAsyncResult, err := server.SendChordWithContext(ctx, chord, 10)
	if err != nil {
		return fmt.Errorf("Could not send chord: %s", err.Error())
	}

	results, err = chordAsyncResult.Get(time.Duration(time.Millisecond * 5))
	if err != nil {
		return fmt.Errorf("Getting chord result failed with error: %s", err.Error())
	}
	log.Printf("(1 + 1) * (2 + 2) * (5 + 6) = %v\n", tasks.HumanReadableResults(results))

	// Now let's try chaining task results
	initTasks()
	log.Println("Chain of tasks:")

	chain, err := tasks.NewChain(&addTask0, &addTask1, &addTask2, &multiplyTask0)
	if err != nil {
		return fmt.Errorf("Error creating chain: %s", err)
	}

	chainAsyncResult, err := server.SendChainWithContext(ctx, chain)
	if err != nil {
		return fmt.Errorf("Could not send chain: %s", err.Error())
	}

	results, err = chainAsyncResult.Get(time.Duration(time.Millisecond * 5))
	if err != nil {
		return fmt.Errorf("Getting chain result failed with error: %s", err.Error())
	}
	log.Printf("(((1 + 1) + (2 + 2)) + (5 + 6)) * 4 = %v\n", tasks.HumanReadableResults(results))

	// Let's try a task which throws panic to make sure stack trace is not lost
	initTasks()
	asyncResult, err = server.SendTaskWithContext(ctx, &panicTask)
	if err != nil {
		return fmt.Errorf("Could not send task: %s", err.Error())
	}

	_, err = asyncResult.Get(time.Duration(time.Millisecond * 5))
	if err == nil {
		return errors.New("Error should not be nil if task panicked")
	}
	log.Printf("Task panicked and returned error = %v\n", err.Error())

	// Let's try a long running task
	initTasks()
	asyncResult, err = server.SendTaskWithContext(ctx, &longRunningTask)
	if err != nil {
		return fmt.Errorf("Could not send task: %s", err.Error())
	}

	results, err = asyncResult.Get(time.Duration(time.Millisecond * 5))
	if err != nil {
		return fmt.Errorf("Getting long running task result failed with error: %s", err.Error())
	}
	log.Printf("Long running task returned = %v\n", tasks.HumanReadableResults(results))

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
