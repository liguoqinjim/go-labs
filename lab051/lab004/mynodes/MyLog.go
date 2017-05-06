package mynodes

import (
	"fmt"
	b3 "github.com/liguoqinjim/behavior3go"
	. "github.com/liguoqinjim/behavior3go/core"
)

type MyLog struct {
	Action
}

func (this *MyLog) OnTick(tick *Tick) b3.Status {
	fmt.Printf("军队%d现在位置%d,血量%d\n", tick.Target.(*Army).Aid, tick.Target.(*Army).Apos, tick.Target.(*Army).Alife)
	return b3.SUCCESS
}
