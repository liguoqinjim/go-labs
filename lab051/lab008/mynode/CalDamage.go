package mynode

import (
	"fmt"
	b3 "github.com/liguoqinjim/behavior3go"
	. "github.com/liguoqinjim/behavior3go/core"
	"lab051/lab008/fight"
)

type CalDamage struct {
	Action
}

func (this *CalDamage) OnTick(tick *Tick) b3.Status {
	fmt.Println("CalDamage节点")
	// todo
	//模拟军队死亡
	//tick.Target.(*fight.Army).SimDead()
	_ = tick.Target.(*fight.Army)

	return b3.SUCCESS
}
