package main

import (
	b3 "github.com/liguoqinjim/behavior3go"
	"github.com/liguoqinjim/behavior3go/config"
	"github.com/liguoqinjim/behavior3go/loader"
	"lab051/lab008/mynode"
	"log"
)

func main() {
	//加载自定义节点
	extMap := b3.NewRegisterStructMaps()
	extMap.Register("Attack", &mynode.Attack{})
	extMap.Register("Move", &mynode.Move{})
	extMap.Register("GetEnemy", &mynode.GetEnemy{})
	extMap.Register("CalDamage", &mynode.CalDamage{})
	extMap.Register("IsDead", &mynode.IsDead{})
	extMap.Register("IsEnemyInRange", &mynode.IsEnemyInRange{})
	extMap.Register("HasEnemy", &mynode.HasEnemy{})

	//加载treeConfig
	treeConfig, ok := config.LoadTreeCfg("tree.json")
	if !ok {
		log.Fatal("loadTree error")
	}

	//生成树
	tree := loader.CreateBevTreeFromConfig(treeConfig, extMap)
	tree.Print()
}
