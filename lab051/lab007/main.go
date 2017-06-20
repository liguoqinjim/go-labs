package main

import (
	b3 "github.com/liguoqinjim/behavior3go"
	//"github.com/liguoqinjim/behavior3go/config"
	//"github.com/liguoqinjim/behavior3go/core"
	//"github.com/liguoqinjim/behavior3go/loader"
	"fmt"
	"lab051/lab007/fight"
	"lab051/lab007/util"
)

func main() {
	a1 := &fight.Army{0, util.SIDE_LEFT, 10, 100, 10}
	a2 := &fight.Army{100, util.SIDE_RIGHt, 5, 50, 5}
	fmt.Println(a1, a2)

	extMap := b3.NewRegisterStructMaps()
}
