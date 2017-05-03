package main

import (
	"github.com/yuin/gopher-lua"
)

func main() {
	L := lua.NewState()
	defer L.Close()

	if err := L.DoString(`print("hello lab001")`); err != nil {
		panic(err)
	}

	//run file
	if err := L.DoFile("hello.lua"); err != nil {
		panic(err)
	}
}
