package mynode

import (
	b3 "github.com/liguoqinjim/behavior3go"
	. "github.com/liguoqinjim/behavior3go/core"
	"lab051/lab007/fight"
)

type HasEnemy struct {
	Condition
}

func (this *HasEnemy) OnTick(tick *Tick) b3.Status {
	if tick.Target.(*fight.Army).HasEnemy() {
		return b3.SUCCESS
	} else {
		return b3.FAILURE
	}
}
