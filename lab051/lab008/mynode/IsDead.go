package mynode

import (
	"fmt"
	b3 "github.com/liguoqinjim/behavior3go"
	. "github.com/liguoqinjim/behavior3go/core"
	"lab051/lab008/fight"
)

type IsDead struct {
	Condition
}

func (this *IsDead) OnTick(tick *Tick) b3.Status {
	fmt.Println("IsDead节点")
	if tick.Target.(*fight.Army).IsDead() {
		return b3.SUCCESS
	} else {
		return b3.FAILURE
	}
}
