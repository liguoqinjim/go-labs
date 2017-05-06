package main

import (
	"fmt"
	b3 "github.com/liguoqinjim/behavior3go"
	"github.com/liguoqinjim/behavior3go/config"
	"github.com/liguoqinjim/behavior3go/core"
	"github.com/liguoqinjim/behavior3go/loader"
	"lab051/lab003/mynodes"
)

func main() {
	//extMap
	extMap := b3.NewRegisterStructMaps()
	//actions
	extMap.Register("Move", &mynodes.Move{})
	extMap.Register("MyLog", &mynodes.MyLog{})

	treeConfig, ok := config.LoadTreeCfg("tree2.json")
	if ok {
		//tree := loader.CreateBevTreeFromConfig(treeConfig, nil)
		tree := loader.CreateBevTreeFromConfig(treeConfig, extMap)
		tree.Print()

		board := core.NewBlackboard()
		p := &mynodes.Player{Pid: 10001, Px: 0}
		tree.Tick(p, board)
	} else {
		fmt.Println("loadTreeCfg error")
	}
}
