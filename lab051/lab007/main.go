package main

import (
	"fmt"
	b3 "github.com/liguoqinjim/behavior3go"
	"github.com/liguoqinjim/behavior3go/config"
	"github.com/liguoqinjim/behavior3go/core"
	"github.com/liguoqinjim/behavior3go/loader"
	"lab051/lab007/fight"
	"lab051/lab007/mynode"
	"lab051/lab007/util"
	"log"
)

func main() {
	a1 := fight.NewArmy(0, util.SIDE_LEFT, 10, 5, 100, 10)
	a2 := fight.NewArmy(100, util.SIDE_RIGHt, 5, 5, 50, 5)
	a1.SetEnemy(a2)
	a2.SetEnemy(a1)

	//加载自定义节点
	extMap := b3.NewRegisterStructMaps()
	extMap.Register("Attack", &mynode.Attack{})
	extMap.Register("Move", &mynode.Move{})
	extMap.Register("GetEnemy", &mynode.GetEnemy{})
	extMap.Register("EnemyInRange", &mynode.EnemyInRange{})
	extMap.Register("HasEnemy", &mynode.HasEnemy{})

	//加载tree.json
	treeConfig, ok := config.LoadTreeCfg("tree.json")
	if !ok {
		log.Fatal("loadTree error")
	}

	//生成树
	tree := loader.CreateBevTreeFromConfig(treeConfig, extMap)
	tree.Print()

	//生成黑板
	board := core.NewBlackboard()
	fmt.Println("\n\n")

	//运行树
	fmt.Println("初始状态")
	fmt.Printf("%+v\n", a1)
	fmt.Printf("%+v\n", a2)

	for !a1.IsDead() && !a2.IsDead() {
		fmt.Println("\nTick")
		tree.Tick(a1, board)
		tree.Tick(a2, board)
		fmt.Printf("%+v\n", a1)
		fmt.Printf("%+v\n", a2)
	}
}
