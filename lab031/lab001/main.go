package main

import (
	"fmt"
	"github.com/shirou/gopsutil/mem"
)

func main() {
	v, _ := mem.VirtualMemory()

	// almost every return value is a struct
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

	// convert to JSON. String() is also implemented
	fmt.Println(v)

	v2, _ := mem.SwapMemory()
	fmt.Println(v2)

	fmt.Println("进程相关\n\n\n\n")
}
