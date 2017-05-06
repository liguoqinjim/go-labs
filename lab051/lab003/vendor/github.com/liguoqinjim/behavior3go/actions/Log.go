package actions

import (
	b3 "github.com/liguoqinjim/behavior3go"
	. "github.com/liguoqinjim/behavior3go/config"
	. "github.com/liguoqinjim/behavior3go/core"
)

type Log struct {
	Action
	info string
}

func (this *Log) Initialize(setting *BTNodeCfg) {
	this.Action.Initialize(setting)
	this.info = setting.GetPropertyAsString("info")
}

func (this *Log) OnTick(tick *Tick) b3.Status {

	return b3.SUCCESS
}
