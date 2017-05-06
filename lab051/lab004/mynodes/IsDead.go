package mynodes

import (
	"fmt"
	b3 "github.com/liguoqinjim/behavior3go"
	. "github.com/liguoqinjim/behavior3go/core"
)

type IsDead struct {
	Condition
}

func (this *IsDead) Tick(tick *Tick) b3.Status {
	fmt.Println("判断是否死亡")
	if tick.Target.(*Army).IsDead() { //死亡
		return b3.SUCCESS
	} else {
		return b3.FAILURE
	}
}
