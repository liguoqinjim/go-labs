package mynodes

import (
	"fmt"
	b3 "github.com/liguoqinjim/behavior3go"
	. "github.com/liguoqinjim/behavior3go/core"
)

type Attack struct {
	Action
}

func (this *Attack) GetEnemiesFromBlackboard(tick *Tick) []*Army {
	side := tick.Target.(*Army).Aside

	//tick.Blackboard.Set("runningChild", 0, tick.GetTree().GetID(), this.GetID())
	if side == 1 {
		return tick.Blackboard.Get("2", tick.GetTree().GetID(), "").(*ArmyGroup).Armys
	} else {
		return tick.Blackboard.Get("1", tick.GetTree().GetID(), "").(*ArmyGroup).Armys
	}
}

func (this *Attack) OnTick(tick *Tick) b3.Status {
	enemies := this.GetEnemiesFromBlackboard(tick)
	tick.Target.(*Army).Attack(enemies)

	fmt.Printf("军队%d进攻\n", tick.Target.(*Army).Aid)
	return b3.SUCCESS
}
