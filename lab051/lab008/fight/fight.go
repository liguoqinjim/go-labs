package fight

import (
	"fmt"
	"lab051/lab008/util"
)

type ArmyGroup struct {
	ArmyGroupId       int
	ArmyFormationId   int
	ArmyFormationSide int //集团军初始位置在战场的哪边(constsUtil里面有常量)
	Armys             []*Army
}

type Army struct {
	Hero                   *Hero
	Soldier                *Soldier
	ArmyFieldId            int   //在整个战场中的id
	TargetFieldId          int   //对手在整个战场的id
	TargetArmy             *Army //有了这个值就不用通过TargetFieldId找了
	TargetPos              int   //近战攻击位置(1-8)
	EnemyPos               uint8 //对手占用的位置
	PosId                  int
	PosX, PosY             float64
	NextHeroAttackFrame    int                      //下一个英雄攻击帧
	NextSoldierAttackFrame int                      //下一个士兵攻击帧
	RangedAttack           bool                     //是否是远程攻击
	State                  int                      //军团状态，moving,attacking,这个参数在远程绕路的时候用
	PosXDetour, PosYDetour float64                  //远程绕路点
	Damages                map[int]*FrameDamageUnit //key是frame
}

type Hero struct {
	HeroId        int //英雄id
	HeroLife      int //英雄血量
	HeroTotalLife int //英雄总血量
}

type Soldier struct {
	SoldierId       int //兵种id
	SoldierNum      int //兵种数量
	SoldierTotalNum int //兵种总数量
}

//伤害单元
type DamageUnit struct {
	//最基本伤害单元
	FromId    int //伤害来自[数值范围1-8]
	FromType  int //1的时候是士兵，2的时候是英雄
	FromFrame int //来自第几帧的伤害
	// ToIds    []int //伤害去向
	ToId int //伤害去向，现在改为就对一个目标产生伤害
}
type FrameDamageUnit struct {
	//每帧伤害单元
	DamageUnits []*DamageUnit
}

//ArmyGroup
func (ag *ArmyGroup) IsDead() bool {
	for _, v := range ag.Armys {
		if !v.IsDead() {
			return false
		}
	}
	return true
}

//Army
func (a *Army) IsDead() bool {
	if a.Soldier.IsDead() && a.Hero.IsDead() {
		return true
	} else {
		return false
	}
}
func (a *Army) SimDead() {
	a.Hero.HeroLife = 0
	a.Soldier.SoldierNum = 0
}
func (a *Army) HasEnemy() bool {
	if a.TargetArmy != nil {
		return true
	} else {
		return false
	}
}
func (a *Army) HasDamage(nowFrame int) bool {
	fmt.Printf("HasDamage,armyId[%d]frame=[%d]\n", a.ArmyFieldId, nowFrame)
	if _, ok := a.Damages[nowFrame]; !ok {
		return false
	} else {
		return true
	}
}
func (a *Army) IsEnemyInRange() bool {
	// todo
	return true
}
func (a *Army) Attack(nowFrame, frame int) {
	unit := new(DamageUnit)
	unit.FromFrame = nowFrame
	unit.FromType = util.ATTACK_FROM_HERO
	unit.FromId = a.ArmyFieldId
	unit.ToId = a.TargetArmy.ArmyFieldId
	a.TargetArmy.AddDamageUnit(frame, unit)
}
func (a *Army) AddDamageUnit(frame int, unit *DamageUnit) {
	v, ok := a.Damages[frame]
	if !ok {
		fd := new(FrameDamageUnit)
		a.Damages[frame] = fd
		fd.DamageUnits = make([]*DamageUnit, 1)
		fd.DamageUnits[0] = unit
	} else {
		v.DamageUnits = append(v.DamageUnits, unit)
	}
}
func (a *Army) GetAtk(atk_type int) int {
	if atk_type == util.ATTACK_FROM_HERO {
		return a.Hero.GetAtk()
	} else {
		return a.Soldier.GetAtk()
	}
}
func (a *Army) BeHurt(damage int) {
	if !a.Soldier.IsDead() {
		restDamage := a.Soldier.BeHurt(damage)
		if restDamage > 0 {
			a.Hero.BeHurt(restDamage)
		}
	} else { //直接扣英雄的血
		a.Hero.BeHurt(damage)
	}
}
func (a *Army) CalDamage(nowFrame int, enemyAg *ArmyGroup) {
	fmt.Println("CalDamage", nowFrame)
	fd := a.Damages[nowFrame] //这里不判断是否为空，为空在上一个节点已经判断过了
	for _, v := range fd.DamageUnits {
		enemy := GetArmyByFieldId(v.FromId, enemyAg)
		if enemy.IsDead() {
			continue
		}
		damage := enemy.GetAtk(v.FromType)
		a.BeHurt(damage)
	}
	delete(a.Damages, nowFrame)
}
func GetArmyByFieldId(fieldId int, ag *ArmyGroup) *Army {
	if fieldId > 4 {
		fieldId -= 4
	}
	return ag.Armys[fieldId-1]
}

//Hero
func (h *Hero) IsDead() bool {
	if h.HeroLife <= 0 {
		return true
	} else {
		return false
	}
}
func (h *Hero) GetAtk() int {
	return HeroBase_map[h.HeroId].HeroAtk
}
func (h *Hero) BeHurt(damage int) {
	h.HeroLife -= damage
	if h.HeroLife < 0 {
		h.HeroLife = 0
	}
}

//Soldier
func (s *Soldier) IsDead() bool {
	if s.SoldierNum == 0 {
		return true
	} else {
		return false
	}
}
func (s *Soldier) GetAtk() int {
	return SoldierBase_map[s.SoldierId].SoldierAtk * s.SoldierNum
}
func (s *Soldier) GetLife() int {
	return SoldierBase_map[s.SoldierId].SoldierLife
}
func (s *Soldier) BeHurt(damage int) (restDamage int) { //返回一个剩余伤害
	restDamage = 0
	deadSoldier := damage / s.GetLife()
	if deadSoldier >= s.SoldierNum {
		restDamage = damage - s.GetLife()*s.SoldierNum
		s.SoldierNum = 0
	} else {
		s.SoldierNum -= deadSoldier
	}
	return restDamage
}
