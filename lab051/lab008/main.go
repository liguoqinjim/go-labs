package main

import (
	"fmt"
	b3 "github.com/liguoqinjim/behavior3go"
	"github.com/liguoqinjim/behavior3go/config"
	"github.com/liguoqinjim/behavior3go/loader"
	"lab051/lab008/fight"
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
	//tree.Print()
	_ = tree

	//生成两支军队
	ag1 := &fight.ArmyGroup{1, 1, 1, make([]*fight.Army, 0)}
	ag2 := &fight.ArmyGroup{2, 2, 2, make([]*fight.Army, 0)}

}

func NewArmy(heroId int) *fight.Army {
	a := new(fight.Army)

	hero := new(fight.Hero)
	hero.HeroId = heroId
	hero.HeroTotalLife = fight.HeroBase_map[heroId].HeroLife
	hero.HeroLife = hero.HeroTotalLife

	soldier := new(fight.Soldier)

	return a
}
