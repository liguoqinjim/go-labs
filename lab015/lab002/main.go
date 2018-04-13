package main

import (
	"log"
	"os/exec"
)

func main() {
	//node E:\Workspace\js-labs\lab020\lab001\js\test1.js
	//E:\\Workspace\\js-labs\\lab020\\lab002\\test1.js
	//E:\Workspace\js-labs\lab022
	nodejsCmd := exec.Command("node", "E:\\Workspace\\js-labs\\lab022\\test.js", `{"rid":"R_SO_4_186016","offset":"0","total":"false","limit":"20","csrf_token":""}`)
	goEnvOut, err := nodejsCmd.Output()
	if err != nil {
		log.Fatalf("output error:%v", err)
	}

	log.Printf("\n%s", string(goEnvOut))
}
