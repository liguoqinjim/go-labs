package main

import (
	"github.com/Knetic/govaluate"
	"log"
)

func main() {
	//evaluate(普通表达式)
	expression, err := govaluate.NewEvaluableExpression("10 > 0")
	if err != nil {
		log.Fatalf("NewValuableExpression error:%v", err)
	}
	result, err := expression.Evaluate(nil)
	if err != nil {
		log.Fatalf("Evaluate error:%v", err)
	}
	log.Println("result=", result)

	//with parameters(传入参数)
	expression2, err := govaluate.NewEvaluableExpression("foo > 0")
	if err != nil {
		log.Fatalf("NewEvaluableExpression error:%v", err)
	}
	parameters := make(map[string]interface{})
	parameters["foo"] = -1
	result, err = expression2.Evaluate(parameters)
	if err != nil {
		log.Fatalf("Evaluate error:%v", err)
	}
	log.Println("result=", result)

	//a complex use(传入多个参数)
	expression3, err := govaluate.NewEvaluableExpression("(requests_made * requests_succeeded / 100) >= 90")
	if err != nil {
		log.Fatalf("NewEvaluableExpression error:%v", err)
	}
	parameters2 := make(map[string]interface{})
	parameters2["requests_made"] = 100
	parameters2["requests_succeeded"] = 80
	result, err = expression3.Evaluate(parameters2)
	if err != nil {
		log.Fatalf("Evaluate error:%v", err)
	}
	log.Println("result=", result)

	//string(比较字符串)
	expression4, err := govaluate.NewEvaluableExpression("http_response_body == 'service is ok'")
	parameters3 := make(map[string]interface{})
	parameters3["http_response_body"] = "service is ok"
	result, err = expression4.Evaluate(parameters3)
	if err != nil {
		log.Fatalf("Evaluate error:%v", err)
	}
	log.Println("result=", result)

	//numeric(返回数字)
	expression5, err := govaluate.NewEvaluableExpression("(mem_used / total_mem) * 100")
	if err != nil {
		log.Fatalf("NewEvaluableExpression error:%v", err)
	}
	parameters4 := make(map[string]interface{})
	parameters4["total_mem"] = 1024
	parameters4["mem_used"] = 512
	result, err = expression5.Evaluate(parameters4)
	if err != nil {
		log.Fatalf("Evaluate error:%v", err)
	}
	log.Println("result=", result)
}
