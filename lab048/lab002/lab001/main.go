package main

import (
	"fmt"
	"github.com/yuin/gopher-lua"
)

func main() {
	L := lua.NewState()
	defer L.Close()
	if err := L.DoFile("test.lua"); err != nil {
		panic(err)
	}

	a := L.GetGlobal("a")
	fmt.Printf("a=%v\n", a)

	b := L.GetGlobal("b").(*lua.LTable)
	b.ForEach(func(key, value lua.LValue) {
		fmt.Printf("b.%v=%v\n", key, value)
	})
}
