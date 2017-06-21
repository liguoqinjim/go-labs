package mynode

import (
	b3 "github.com/liguoqinjim/behavior3go"
	. "github.com/liguoqinjim/behavior3go/core"
)

type GetEnemy struct {
	Action
}

func (this *GetEnemy) Tick(tick *Tick) b3.Status {
	return b3.SUCCESS
}
