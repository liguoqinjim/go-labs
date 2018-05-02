package main

import (
	"log"
	"os"
	"os/exec"
)

//因为是在windows里面写的，就不用`sh -c`了，直接测试echo命令
func Call(cmd, value string) (int, string) {
	//bytes, err := exec.Command("sh", "-c", cmd).CombinedOutput()
	bytes, err := exec.Command(cmd, value).CombinedOutput()
	output := string(bytes)
	if err != nil {
		return 1, reportExecFailed(output)
	}

	return 0, output
}

func reportExecFailed(msg string) string {
	os.Exit(1)
	return msg
}

func main() {
	r, output := Call("echo", "123")
	log.Println("r=", r)
	log.Println("output=", output)
}
