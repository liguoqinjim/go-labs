package main

import (
	"fmt"
	"github.com/robertkrimen/otto"
	"log"
	"time"
)

func main() {
	startTime := time.Now()

	vm := otto.New()
	script, err := vm.Compile("js/test1.js", nil)
	if err != nil {
		log.Fatalf("vm error:%+v", err)
	}
	t1 := time.Since(startTime)
	log.Println("t1:", t1)

	_, err = vm.Run(script)
	if err != nil {
		log.Fatalf("vm run error:%+v", err)
	}

	t2 := time.Since(startTime)
	log.Println("t2:", t2)

	//开始拼接出加密需要的参数
	//186016;R_SO_4_186016;2132178;4;1
	commentId := "R_SO_4_186016"
	pageNum := 1
	commentParams := GetCommentParam(commentId, pageNum)
	log.Printf("commentParams=%s", commentParams)
	funcName := fmt.Sprintf("myTest(%s)", commentParams)
	value, _ := vm.Run(funcName)
	if !value.IsObject() {
		log.Fatalf("js的返回不是对象类型")
	}

	t3 := time.Since(startTime)
	log.Println("t3:", t3)

	obj := value.Object()
	encTextValue, _ := obj.Get("encText")
	encText := encTextValue.String()
	encSecKeyValue, _ := obj.Get("encSecKey")
	encSecKey := encSecKeyValue.String()
	log.Printf("encText:%s", encText)
	log.Printf("encSecKey:%s", encSecKey)

	t4 := time.Since(startTime)
	log.Println("t4:", t4)
}

func GetCommentParam(commentId string, pageNum int) string {
	//var d = '{"rid":"R_SO_4_186016","offset":"100","total":"false","limit":"20","csrf_token":""}';
	pageNum = 20 * (pageNum - 1)
	commentParam := fmt.Sprintf(`'{"rid":"%s","offset":"%d","total":"false","limit":"20","csrf_token":""}'`, commentId, pageNum)

	return commentParam
}
