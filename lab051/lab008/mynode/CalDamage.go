package mynode

import (
	"fmt"
	b3 "github.com/liguoqinjim/behavior3go"
	. "github.com/liguoqinjim/behavior3go/core"
	"lab051/lab008/fight"
	"lab051/lab008/util"
)

type CalDamage struct {
	Action
}

func (this *CalDamage) OnTick(tick *Tick) b3.Status {
	fmt.Println("CalDamage节点")
	// todo
	//模拟军队死亡
	//tick.Target.(*fight.Army).SimDead()
	var enemyAg *fight.ArmyGroup
	nowFrame := tick.Blackboard.GetInt(util.BOARD_KEY_FRAME, "", "")
	if tick.Target.(*fight.Army).ArmyFieldId > 4 {
		enemyAg = tick.Blackboard.Get(util.BOARD_KEY_AG1, "", "").(*fight.ArmyGroup)
	} else {
		enemyAg = tick.Blackboard.Get(util.BOARD_KEY_AG2, "", "").(*fight.ArmyGroup)
	}
	tick.Target.(*fight.Army).CalDamage(nowFrame, enemyAg)

	return b3.SUCCESS
}
