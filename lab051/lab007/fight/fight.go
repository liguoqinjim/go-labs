package fight

import (
	"lab051/lab007/util"
)

type Army struct {
	x           int
	side        int //队伍的左右
	attack      int //攻击力
	attackRange int //攻击范围
	life        int //生命值
	speed       int //移动速读
	enemy       *Army
}

func NewArmy(x, side, attack, attackRange, life, speed int) *Army {
	return &Army{x: x, side: side, attack: attack, attackRange: attackRange, life: life, speed: speed}
}

func (army *Army) Move() {
	if army.side == util.SIDE_LEFT {
		army.x += army.speed
	} else {
		army.x -= army.speed
	}
}

func (army *Army) Attack() {
	if army.enemy != nil {
		if !army.enemy.IsDead() {
			army.enemy.life -= army.attack
			if army.enemy.life < 0 {
				army.enemy.life = 0
			}
		}
	}
}

func (army *Army) SetEnemy(enemy *Army) {
	army.enemy = enemy
}

func (army *Army) IsDead() bool {
	if army.life <= 0 {
		return true
	} else {
		return false
	}
}

func (army *Army) HasEnemy() bool {
	if army.enemy == nil {
		return false
	} else {
		return true
	}
}

func (army *Army) IsEnemyInRange() bool {
	distance := army.x - army.enemy.x
	if distance < 0 {
		distance = -distance
	}
	if distance <= army.attackRange {
		return true
	} else {
		return false
	}
}
