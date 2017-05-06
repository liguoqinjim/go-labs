package mynodes

import (
	"fmt"
	b3 "github.com/liguoqinjim/behavior3go"
	. "github.com/liguoqinjim/behavior3go/core"
)

type HasEnemyInRange struct {
	Condition
}

func (this *HasEnemyInRange) GetEnemiesFromBlackboard(tick *Tick) []*Army { //现在默认黑板里面key1的时候1这边的armygroup
	side := tick.Target.(*Army).Aside

	//tick.Blackboard.Set("runningChild", 0, tick.GetTree().GetID(), this.GetID())
	if side == 1 {
		return tick.Blackboard.Get("2", tick.GetTree().GetID(), "").(*ArmyGroup).Armys
	} else {
		return tick.Blackboard.Get("1", tick.GetTree().GetID(), "").(*ArmyGroup).Armys
	}
}

func (this *HasEnemyInRange) OnTick(tick *Tick) b3.Status {
	if tick.Target.(*Army).HasEnemyInRange(this.GetEnemiesFromBlackboard(tick)) { //有敌人在射程内
		fmt.Println("判断射程内是否有敌人success")
		return b3.SUCCESS
	} else {
		fmt.Println("判断射程内是否有敌人failure")
		return b3.FAILURE
	}
}
