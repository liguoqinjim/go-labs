package main

import (
	b3 "github.com/liguoqinjim/behavior3go"
	//"github.com/liguoqinjim/behavior3go/config"
	//"github.com/liguoqinjim/behavior3go/core"
	//"github.com/liguoqinjim/behavior3go/loader"
	"fmt"
	"lab051/lab007/util"
)

type Army struct {
	x      int
	side   int //队伍的左右
	attack int //攻击力
	life   int //生命值
	speed  int //移动速读
}

func main() {
	a1 := &Army{0, util.SIDE_LEFT, 10, 100, 10}
	a2 := &Army{100, util.SIDE_RIGHt, 5, 50, 5}
	fmt.Println(a1, a2)

	extMap := b3.NewRegisterStructMaps()
}
