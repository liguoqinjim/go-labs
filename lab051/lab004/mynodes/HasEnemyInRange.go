package mynodes

import (
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
		return tick.Blackboard.Get("2", tick.GetTree().GetID(), "").([]*Army)
	} else {
		return tick.Blackboard.Get("1", tick.GetTree().GetID(), "").([]*Army)
	}
}

func (this *HasEnemyInRange) Tick(tick *Tick) b3.Status {
	if tick.Target.(*Army).HasEnemyInRange(this.GetEnemiesFromBlackboard(tick)) { //有敌人在射程内
		return b3.SUCCESS
	} else {
		return b3.FAILURE
	}
}
