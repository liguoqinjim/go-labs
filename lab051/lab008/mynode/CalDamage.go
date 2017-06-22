package mynode

import (
	b3 "github.com/liguoqinjim/behavior3go"
	. "github.com/liguoqinjim/behavior3go/core"
)

type CalDamage struct {
	Action
}

func (this *CalDamage) OnTick(tick *Tick) b3.Status {
	// todo
	return b3.SUCCESS
}
