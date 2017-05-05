package main

import (
	"fmt"
	"github.com/magicsea/behavior3go/config"
	"github.com/magicsea/behavior3go/core"
	"github.com/magicsea/behavior3go/loader"
)

func main() {
	treeConfig, ok := config.LoadTreeCfg("tree.json")
	if ok {
		tree := loader.CreateBevTreeFromConfig(treeConfig, nil)
		tree.Print()
		fmt.Println("\n\n")

		board := core.NewBlackboard()
		for i := 0; i < 5; i++ {
			tree.Tick(i, board)
		}
	} else {
		fmt.Println("LoadTreeCfg error")
	}
}
