package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	shell01()
	shell02()
}

//脚本参数
func shell01() {
	cmd := exec.Command("/bin/bash", "lab001.sh", "tom")
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("cmd.Output error:%v", err)
	}
	log.Printf("lab001.sh output:%s", output)
}

//可以捕捉脚本错误
func shell02() {
	cmd := exec.Command("/bin/bash", "lab002.sh")
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("cmd.Output error:%v", err)
	}
	log.Printf("lab002.sh output:%s", output)
}

func shell03() {
	cmd := exec.Command("find", "/", "-maxdepth", "1", "-exec", "wc", "-c", "{}", "\\")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}
	fmt.Println("Result: " + out.String())
}
