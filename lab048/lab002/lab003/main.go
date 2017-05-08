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

	if err := L.CallByParam(lua.P{
		Fn:      L.GetGlobal("concat"),
		NRet:    1,
		Protect: true,
	}, lua.LString("go"), lua.LString("lua")); err != nil {
		panic(err)
	}

	if str, ok := L.Get(-1).(lua.LString); ok {
		fmt.Println("返回值", str)
	}

	L.Pop(1)
}
