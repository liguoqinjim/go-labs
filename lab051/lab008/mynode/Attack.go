package mynode

import (
	"fmt"
	b3 "github.com/liguoqinjim/behavior3go"
	. "github.com/liguoqinjim/behavior3go/core"
)

type Attack struct {
	Action
}

func (this *Attack) OnTick(tick *Tick) b3.Status {
	// todo
	fmt.Println("Attack节点")
	return b3.SUCCESS
}
