package mynode

import (
	b3 "github.com/liguoqinjim/behavior3go"
	. "github.com/liguoqinjim/behavior3go/core"
)

type Move struct {
	Action
}

func (this *Move) OnTick(tick *Tick) b3.Status {
	// todo

	return b3.SUCCESS
}
