package main

import (
	"fmt"
	"github.com/liguoqinjim/behavior3go"
	b3 "github.com/liguoqinjim/behavior3go"
	"github.com/liguoqinjim/behavior3go/config"
	"github.com/liguoqinjim/behavior3go/core"
	"github.com/liguoqinjim/behavior3go/loader"
	"lab051/lab006/mynode"
)

func main() {
	treeConfig, _ := config.LoadTreeCfg("tree_memsequence.json")
	extMap := b3.NewRegisterStructMaps()
	extMap.Register("IsMax", &mynode.IsMax{})

	tree := loader.CreateBevTreeFromConfig(treeConfig, extMap)
	tree.Print()
	fmt.Println("\n\n")

	board := core.NewBlackboard()
	i := 0

	nowFrame := 1
	board.Set("now", nowFrame, "", "")
	status := tree.Tick(i, board)
	fmt.Println("status=", status)
	for status == behavior3go.RUNNING {
		nowFrame += 1
		board.Set("now", nowFrame, "", "")
		status = tree.Tick(i, board)
		fmt.Println("status=", status)
	}
	fmt.Println("\n\n")
	status = tree.Tick(i, board)
	fmt.Println("end status=", status)
}
