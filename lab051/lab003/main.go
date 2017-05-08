package main

import (
	"fmt"
	b3 "github.com/liguoqinjim/behavior3go"
	"github.com/liguoqinjim/behavior3go/config"
	"github.com/liguoqinjim/behavior3go/core"
	"github.com/liguoqinjim/behavior3go/loader"
	"lab051/lab003/mynodes"
	"os"
	"os/signal"
	"syscall"
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
		p1 := &mynodes.Player{Pid: 10001, Px: 0}
		p2 := &mynodes.Player{Pid: 10002, Px: 100}
		p3 := &mynodes.Player{Pid: 10003, Px: 200}
		p4 := &mynodes.Player{Pid: 10004, Px: 300}
		p5 := &mynodes.Player{Pid: 10005, Px: 400}

		go tree.Tick(p1, board)
		go tree.Tick(p2, board)
		go tree.Tick(p3, board)
		go tree.Tick(p4, board) //5个一起的时候就会报错
		go tree.Tick(p5, board)
	} else {
		fmt.Println("loadTreeCfg error")
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	fmt.Println("程序结束")
}
