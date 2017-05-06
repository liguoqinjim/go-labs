package main

import (
	"fmt"
	b3 "github.com/liguoqinjim/behavior3go"
	"github.com/liguoqinjim/behavior3go/config"
	"github.com/liguoqinjim/behavior3go/core"
	"github.com/liguoqinjim/behavior3go/loader"
	"lab051/lab004/mynodes"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//生成队伍,先两支队伍
	a1 := mynodes.NewArmy(10001, 1, 0, 100, 15, 15, 10)
	a2 := mynodes.NewArmy(10002, 2, 100, 120, 10, 10, 5)
	ag1 := mynodes.NewArmyGroup()
	ag1.Armys = append(ag1.Armys, a1)
	ag2 := mynodes.NewArmyGroup()
	ag2.Armys = append(ag2.Armys, a2)

	//extMap
	extMap := b3.NewRegisterStructMaps()
	//actions
	extMap.Register("Attack", &mynodes.Attack{})
	extMap.Register("HasEnemyInRange", &mynodes.HasEnemyInRange{})
	extMap.Register("IsDead", &mynodes.IsDead{})
	extMap.Register("Move", &mynodes.Move{})
	extMap.Register("MyLog", &mynodes.MyLog{})

	treeConfig, ok := config.LoadTreeCfg("tree.json")
	if ok {
		//tree := loader.CreateBevTreeFromConfig(treeConfig, nil)
		tree1 := loader.CreateBevTreeFromConfig(treeConfig, extMap)
		tree2 := loader.CreateBevTreeFromConfig(treeConfig, extMap)
		tree1.Print()
		fmt.Println("\n\n")

		board := core.NewBlackboard()
		board.Set("1", ag1, tree1.GetID(), "")
		board.Set("2", ag2, tree2.GetID(), "")
		tree1.Tick(a1, board)
		fmt.Println("\n\n")
		//tree2.Tick(a2, board)
	} else {
		fmt.Println("loadTreeCfg error")
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	fmt.Println("程序结束")
}
