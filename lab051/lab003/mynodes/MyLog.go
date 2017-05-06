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
	fmt.Println("玩家现在位置:", posX)

	return b3.SUCCESS
}
