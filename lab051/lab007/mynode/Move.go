package mynode

import (
	b3 "github.com/liguoqinjim/behavior3go"
	. "github.com/liguoqinjim/behavior3go/core"
	"lab051/lab007/fight"
)

type Move struct {
	Action
}

func (this *Move) OnTick(tick *Tick) b3.Status {
	tick.Target.(*fight.Army).Move()
	return b3.SUCCESS
}
