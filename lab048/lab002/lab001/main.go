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
	L.

	lv := L.Get(-1)
	fmt.Println(lv)
	if i, ok := lv.(lua.LNumber); ok {
		fmt.Println("i=", i)
	} else {
		fmt.Println("error")
	}

}
