package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"unicode"
)

func convert(input string) string {
	var buf bytes.Buffer
	for _, r := range input {
		if unicode.IsControl(r) {
			fmt.Fprintf(&buf, "\\u%04X", r)
		} else {
			fmt.Fprintf(&buf, "%c", r)
		}
	}
	return buf.String()
}

func main() {
	dateCmd := exec.Command("dir")

	dateOut, err := dateCmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(dateOut))

	goEnvCmd := exec.Command("go", "env")
	goEnvOut, _ := goEnvCmd.Output()
	fmt.Println(string(goEnvOut))
}
