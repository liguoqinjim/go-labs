package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	fmt.Println("go fmt test.go")

	goFmtCmd := exec.Command("go", "fmt", "test.go")
	cmdOut, err := goFmtCmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(cmdOut))
}
