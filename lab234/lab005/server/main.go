package main

import (
	"github.com/RichardKnop/machinery/example/tracers"
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/RichardKnop/machinery/v1/tasks"
	"log"
	"time"
)

func main() {
	worker()
}

func worker() error {
	consumerTag := "machinery_worker"

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
		"add": Add,
	}

	return server, server.RegisterTasks(tasks)
}

func Add(args ...int64) (int64, error) {
	log.Println("当前时间:", time.Now())

	sum := int64(0)
	for _, arg := range args {
		sum += arg
	}
	return sum, nil
}
