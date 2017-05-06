package mynodes

import (
	"fmt"
	b3 "github.com/liguoqinjim/behavior3go"
	. "github.com/liguoqinjim/behavior3go/core"
)

type IsAlive struct {
	Condition
}

func (this *IsAlive) OnTick(tick *Tick) b3.Status {
	if !tick.Target.(*Army).IsDead() { //活着
		fmt.Println("判断是否活着success")
		return b3.SUCCESS
	} else {
		fmt.Println("判断是否活着failure")
		return b3.FAILURE
	}
}
