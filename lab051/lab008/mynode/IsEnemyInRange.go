package mynode

import (
	"fmt"
	b3 "github.com/liguoqinjim/behavior3go"
	. "github.com/liguoqinjim/behavior3go/core"
	"lab051/lab008/fight"
)

type IsEnemyInRange struct {
	Action
}

func (this *IsEnemyInRange) OnTick(tick *Tick) b3.Status {
	fmt.Println("IsEnemyInRange节点")
	if tick.Target.(*fight.Army).IsEnemyInRange() {
		return b3.SUCCESS
	} else {
		return b3.FAILURE
	}
}
