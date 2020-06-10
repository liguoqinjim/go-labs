package main

import (
	exampletasks "github.com/RichardKnop/machinery/example/tasks"
	"github.com/RichardKnop/machinery/example/tracers"
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/RichardKnop/machinery/v1/tasks"
	"log"
)

func main() {
	worker()
}

func worker() error {
	consumerTag := "machinery_worker01"

	cleanup, err := tracers.SetupTracer(consumerTag)
	if err != nil {
		log.Fatalf("setupTracer error:%v", err)
	}
	defer cleanup()

	server, err := startServer()
	if err != nil {
		log.Fatalf("startServer error:%v", err)
	}

	// The second argument is a consumer tag
	// Ideally, each worker should have a unique tag (worker1, worker2 etc)
	worker := server.NewWorker(consumerTag, 2)

	// Here we inject some custom code for error handling,
	// start and end of task hooks, useful for metrics for example.
	errorHandler := func(err error) {
		log.Printf("errorHandler err:%v", err)
	}

	preTaskHandler := func(signature *tasks.Signature) {
		log.Printf("preTaskHandler err:%v", err)
	}

	postTaskHandler := func(signature *tasks.Signature) {
		log.Printf("postTaskHandler err:%v", err)
	}

	worker.SetPostTaskHandler(postTaskHandler)
	worker.SetErrorHandler(errorHandler)
	worker.SetPreTaskHandler(preTaskHandler)

	return worker.Launch()
}

func loadConfig() (*config.Config, error) {
	return config.NewFromYaml("./config.yml", true)
}

func startServer() (*machinery.Server, error) {
	cnf, err := loadConfig()
	if err != nil {
		return nil, err
	}

	cnf.ResultsExpireIn = 3600 * 24

	// Create server instance
	server, err := machinery.NewServer(cnf)
	if err != nil {
		return nil, err
	}

	// Register tasks
	tasks := map[string]interface{}{
		"add":               exampletasks.Add,
		"multiply":          exampletasks.Multiply,
		//"sum_ints":          exampletasks.SumInts,
		//"sum_floats":        exampletasks.SumFloats,
		//"concat":            exampletasks.Concat,
		//"split":             exampletasks.Split,
		//"panic_task":        exampletasks.PanicTask,
		//"long_running_task": exampletasks.LongRunningTask,
	}

	return server, server.RegisterTasks(tasks)
}
