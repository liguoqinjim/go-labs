package mynodes

import (
	"fmt"
	"math/rand"
)

type Army struct {
	Aid     int
	Aside   int
	Apos    int
	Alife   int //生命值
	Arange  int //攻击范围
	Aattack int //攻击力
	Aspeed  int //移动速度
}

type ArmyGroup struct {
	Armys []*Army
}

func NewArmyGroup() *ArmyGroup {
	return &ArmyGroup{Armys: make([]*Army, 0)}
}

func NewArmy(id, side, pos, life, arange, attack, speed int) *Army {
	return &Army{Aid: id, Aside: side, Apos: pos, Alife: life, Arange: arange, Aattack: attack, Aspeed: speed}
}

func (a *Army) IsDead() bool { //返回true的队伍死亡
	if a.Alife <= 0 {
		return true
	} else {
		return false
	}
}

func (a *Army) HasEnemyInRange(enemies []*Army) bool { //有敌人返回true
	for _, v := range enemies {
		if v.Alife <= 0 {
			continue
		}
		dis := a.Apos - v.Apos
		if dis < 0 {
			dis = -dis
		}

		if dis <= a.Arange {
			return true
		}
	}
	return false
}

func (a *Army) Attack(enemies []*Army) {
	for _, v := range enemies {
		if v.Alife <= 0 {
			continue
		}

		damage := rand.Intn(a.Aattack) + 1
		v.Alife -= damage
	}
	fmt.Printf("Army%d开始攻击\n", a.Aid)
}

func (a *Army) Move() {
	if a.Aside == 1 { //+
		a.Apos += a.Aspeed
	} else { //-
		a.Apos -= a.Aspeed
	}
	fmt.Printf("Army%d移动，移动到%d\n", a.Aid, a.Apos)
}
