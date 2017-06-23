package mynode

import (
	"fmt"
	b3 "github.com/liguoqinjim/behavior3go"
	. "github.com/liguoqinjim/behavior3go/core"
	"lab051/lab008/fight"
)

type Attack struct {
	Action
}

func (this *Attack) OnTick(tick *Tick) b3.Status {
	// todo
	fmt.Println("Attack节点")
	frame := tick.Blackboard.GetInt("frame", "", "")
	tick.Target.(*fight.Army).Attack(frame, frame+1)
	return b3.SUCCESS
}
