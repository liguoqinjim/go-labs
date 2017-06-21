package mynode

import (
	"fmt"
	b3 "github.com/liguoqinjim/behavior3go"
	. "github.com/liguoqinjim/behavior3go/core"
	"lab051/lab007/fight"
)

type Attack struct {
	Action
}

func (this *Attack) OnTick(tick *Tick) b3.Status {
	fmt.Println("Attack节点")
	tick.Target.(*fight.Army).Attack()
	return b3.SUCCESS
}
