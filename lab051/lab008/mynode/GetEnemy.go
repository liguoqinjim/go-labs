package mynode

import (
	"fmt"
	b3 "github.com/liguoqinjim/behavior3go"
	. "github.com/liguoqinjim/behavior3go/core"
)

type GetEnemy struct {
	Action
}

func (this *GetEnemy) OnTick(tick *Tick) b3.Status {
	// todo
	fmt.Println("GetEnemy节点")
	return b3.SUCCESS
}
