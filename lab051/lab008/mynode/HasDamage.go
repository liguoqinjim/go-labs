package mynode

import (
	"fmt"
	b3 "github.com/liguoqinjim/behavior3go"
	. "github.com/liguoqinjim/behavior3go/core"
	"lab051/lab008/fight"
	"lab051/lab008/util"
)

type HasDamage struct {
	Condition
}

func (this *HasDamage) OnTick(tick *Tick) b3.Status {
	fmt.Println("HasDamage节点")
	nowFrame := tick.Blackboard.GetInt(util.BOARD_KEY_FRAME, "", "")
	if tick.Target.(*fight.Army).HasDamage(nowFrame) {
		fmt.Println("HasDamage节点1,", nowFrame)
		return b3.SUCCESS
	} else {
		fmt.Println("HasDamage节点2,", nowFrame)
		return b3.FAILURE
	}
}
