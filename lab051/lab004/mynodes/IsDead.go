package mynodes

import (
	"fmt"
	b3 "github.com/liguoqinjim/behavior3go"
	. "github.com/liguoqinjim/behavior3go/core"
)

type IsDead struct {
	Condition
}

func (this *IsDead) OnTick(tick *Tick) b3.Status {
	if tick.Target.(*Army).IsDead() { //死亡
		fmt.Println("判断是否死亡success")
		return b3.SUCCESS
	} else {
		fmt.Println("判断是否死亡failure")
		return b3.FAILURE
	}
}
