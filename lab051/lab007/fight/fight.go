package fight

import "lab051/lab007/util"

type Army struct {
	x           int
	side        int //队伍的左右
	attack      int //攻击力
	attackRange int //攻击范围
	life        int //生命值
	speed       int //移动速读
	enemy       *Army
}

func (army *Army) Move() {
	if army.side == util.SIDE_LEFT {
		army.x += army.speed
	} else {
		army.x -= army.speed
	}
}

func (army *Army) Attack(enemy *Army) {
	enemy.life -= army.attack
	if enemy.life < 0 {
		enemy.life = 0
	}
}

func (army *Army) IsDead() bool {
	if army.life <= 0 {
		return true
	} else {
		return false
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
