package mynodes

import (
	"fmt"
	b3 "github.com/liguoqinjim/behavior3go"
	. "github.com/liguoqinjim/behavior3go/core"
)

type Move struct {
	Action
}

func (this *Move) OnTick(tick *Tick) b3.Status {
	tick.Target.(*Army).Move()

	fmt.Printf("军队%d移动\n", tick.Target.(*Army).Aid)
	return b3.SUCCESS
}
