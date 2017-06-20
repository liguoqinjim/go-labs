package mynode

import (
	"fmt"
	b3 "github.com/liguoqinjim/behavior3go"
	. "github.com/liguoqinjim/behavior3go/config"
	. "github.com/liguoqinjim/behavior3go/core"
)

type IsMax struct {
	Condition
	max int
}

func (this *IsMax) Initialize(params *BTNodeCfg) {
	this.max = params.GetPropertyAsInt("max")
}

func (this *IsMax) OnOpen(tick *Tick) {
	fmt.Println("IsMax节点 OnOpen")
}

func (this *IsMax) OnTick(tick *Tick) b3.Status {
	now := tick.Blackboard.Get("now", "", "").(int)
	fmt.Println("IsMax节点 now=", now)
	if now < this.max {
		return b3.SUCCESS
	} else {
		return b3.FAILURE
	}
}
