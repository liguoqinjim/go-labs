package mynode

import (
	"fmt"
	b3 "github.com/liguoqinjim/behavior3go"
	. "github.com/liguoqinjim/behavior3go/core"
)

type Move struct {
	Action
}

func (this *Move) OnTick(tick *Tick) b3.Status {
	// todo
	fmt.Println("Move节点")
	return b3.SUCCESS
}
