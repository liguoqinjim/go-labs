package mynodes

import (
	"fmt"
	b3 "github.com/liguoqinjim/behavior3go"
	. "github.com/liguoqinjim/behavior3go/core"
)

type IsBattleOver struct {
	Condition
}

func (this *IsBattleOver) OnTick(tick *Tick) b3.Status { //结束的时候返回success
	ag1 := tick.Blackboard.Get("1", tick.GetTree().GetID(), "").(*ArmyGroup)
	if ag1.IsDead() {
		fmt.Println("判断整个战斗是否结束success")
		return b3.SUCCESS
	}
	ag2 := tick.Blackboard.Get("2", tick.GetTree().GetID(), "").(*ArmyGroup)
	if ag2.IsDead() {
		fmt.Println("判断整个战斗是否结束success")
		return b3.SUCCESS
	}
	fmt.Println("判断整个战斗是否结束failure")
	return b3.FAILURE
}
