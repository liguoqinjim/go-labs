package main

import (
	"fmt"
	b3 "github.com/liguoqinjim/behavior3go"
	"github.com/liguoqinjim/behavior3go/config"
	"github.com/liguoqinjim/behavior3go/core"
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
	extMap.Register("HasDamage", &mynode.HasDamage{})

	//加载treeConfig
	treeConfig, ok := config.LoadTreeCfg("tree.json")
	if !ok {
		log.Fatal("loadTree error")
	}

	//生成树
	tree := loader.CreateBevTreeFromConfig(treeConfig, extMap)
	//tree.Print()

	//生成黑板
	board := core.NewBlackboard()
	_ = board

	//生成两支军队
	ag1 := &fight.ArmyGroup{1, 1, 1, make([]*fight.Army, 0)}
	ag2 := &fight.ArmyGroup{2, 2, 2, make([]*fight.Army, 0)}
	a1 := NewArmy(311001)
	a5 := NewArmy(311005)
	ag1.Armys = append(ag1.Armys, a1)
	ag2.Armys = append(ag2.Armys, a5)

	//修改targetFieldId
	SetArmyGroupId(ag1)
	SetArmyGroupId(ag2)

	fmt.Println("\n\n")
	//tick
	//在黑板中加入frame参数
	board.Set("frame", 0, "", "")
	//模拟有对手
	a1.TargetArmy = a5
	a5.TargetArmy = a1
	//打印两队初始状态
	fmt.Println("初始状态")
	for _, v := range ag1.Armys {
		fmt.Printf("军队id[%d],len(damages)=[%d]\n", v.TargetFieldId, len(v.Damages))
	}
	for _, v := range ag2.Armys {
		fmt.Printf("军队id[%d],len(damages)=[%d]\n", v.TargetFieldId, len(v.Damages))
	}
	fmt.Println("")
	for i := 0; i < 2; i++ {
		for _, v := range ag1.Armys {
			status := tree.Tick(v, board)
			fmt.Println("status=", status)
		}
		fmt.Println("")
		for _, v := range ag2.Armys {
			status := tree.Tick(v, board)
			fmt.Println("status=", status)
		}
		fmt.Println()
		frame := board.GetInt("frame", "", "") + 1
		board.Set("frame", frame, "", "")
	}
	//打印结束状态
	for _, v := range ag1.Armys {
		fmt.Printf("军队id[%d],len(damages)=[%d]\n", v.TargetFieldId, len(v.Damages))
	}
	fmt.Println("")
	for _, v := range ag2.Armys {
		fmt.Printf("军队id[%d],len(damages)=[%d]\n", v.TargetFieldId, len(v.Damages))
	}
}

func SetArmyGroupId(ag *fight.ArmyGroup) {
	for n, v := range ag.Armys {
		if ag.ArmyFormationSide == 1 {
			v.ArmyFieldId = n + 1
		} else {
			v.ArmyFieldId = n + 5
		}
	}
}

func NewArmy(heroId int) *fight.Army {
	a := new(fight.Army)

	hb := fight.HeroBase_map[heroId]
	if hb == nil {
		log.Fatal("heroBase id error ", heroId)
	}

	hero := new(fight.Hero)
	hero.HeroId = heroId
	hero.HeroTotalLife = hb.HeroLife
	hero.HeroLife = hero.HeroTotalLife

	soldier := new(fight.Soldier)
	soldier.SoldierId = hb.SoldierId
	soldier.SoldierTotalNum = hb.SoldierNum
	soldier.SoldierNum = hb.SoldierNum

	a.Hero = hero
	a.Soldier = soldier
	a.RangedAttack = hb.RangeAttack
	a.Damages = make(map[int]*fight.FrameDamageUnit)

	return a
}
