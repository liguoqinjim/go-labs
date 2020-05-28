package main

import (
	"fmt"
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/RichardKnop/machinery/v1/tasks"
	"log"
	"time"
)

var (
	consumerTag  = "machinery_worker"
	DefaultQueue = "machinery_tasks"
)

func main() {
	worker()
}

func worker() error {
	//cleanup, err := tracers.SetupTracer(consumerTag)
	//if err != nil {
	//	log.Fatalf("setupTracer error:%v", err)
	//}
	//defer cleanup()

	server, err := startServer()
	if err != nil {
		log.Fatalf("startServer error:%v", err)
	}

	// The second argument is a consumer tag
	// Ideally, each worker should have a unique tag (worker1, worker2 etc)
	worker := server.NewWorker("001号worker", 1)
	worker2 := server.NewWorker("002worker", 1)

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

	go worker2.Launch()

	return worker.Launch()
}

func startServer() (*machinery.Server, error) {
	var cnf = &config.Config{
		Broker:          "redis://localhost:6379/3",
		DefaultQueue:    DefaultQueue,
		ResultBackend:   "redis://localhost:6379/4",
		ResultsExpireIn: 3600,
	}

	server, err := machinery.NewServer(cnf)
	if err != nil {
		return nil, err
	}

	// Register tasks
	tasks := map[string]interface{}{
		"test001": Test001,
		"test002": Test002,
		"test003": Test003,
	}

	return server, server.RegisterTasks(tasks)
}

// 测试参数
func Test001(str string, id int) error {
	log.Printf("Task[Test001] [%s,%d]", str, id)

	return nil
}

//实验worker
func Test002(id int) error {
	log.Printf("Task[Test002] [%d]", id)
	time.Sleep(time.Second * 2)

	return nil
}

//实验retry
func Test003(id int) error {
	log.Printf("Task[Test003] id:%d", id)

	if id == 2 {
		return fmt.Errorf("special id error")
	}

	return nil
}
