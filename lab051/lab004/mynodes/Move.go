package mynodes

import (
	b3 "github.com/liguoqinjim/behavior3go"
	. "github.com/liguoqinjim/behavior3go/core"
)

type Move struct {
	Action
}

func (this *Move) OnTick(tick *Tick) b3.Status {
	tick.Target.(*Army).Move()

	return b3.SUCCESS
}
