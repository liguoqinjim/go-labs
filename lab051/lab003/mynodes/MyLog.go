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
	posX := tick.Target.(*Player).Px
	pid := tick.Target.(*Player).Pid
	fmt.Printf("玩家%d现在位置:%d\n", pid, posX)

	return b3.SUCCESS
}
