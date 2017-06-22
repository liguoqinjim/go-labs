package fight

import (
	"fmt"
	"lab051/lab008/util"
	"strconv"
	"strings"
)

//xml映射struct start
var FightFactorsXml FightFactorsXmlStruct

type FightFactorsXmlStruct struct {
	FightFactor          FightFactorXmlStruct          `xml:"FightFactor"`
	CloseCombatPosFactor CloseCombatPosFactorXmlStruct `xml:"CloseCombatPosFactor"`
	MoveSpeedFactor      MoveSpeedXmlStruct            `xml:"MoveSpeedFactor"`
}

//战斗参数
type FightFactorXmlStruct struct {
	FieldLength                string `xml:"FieldLength,attr"`
	PauseTime                  string `xml:"PauseTime,attr"`
	SputteringRange            string `xml:"SputteringRange,attr"`
	RangedAttackDetourDistance string `xml:"RangedAttackDetourDistance,attr"`
}

//近战攻击位置偏移参数
type CloseCombatPosFactorXmlStruct struct {
	PosOffset string `xml:"PosOffset,attr"`
}

//移动速度参数
type MoveSpeedXmlStruct struct {
	MoveSpeed string `xml:"MoveSpeed,attr"`
}

//集团军参数
type ArmyGroupXml struct {
	ArmyGroupId          string `xml:"ArmyGroupId,attr"`
	ArmyGroupFormationId string `xml:"ArmyGroupFormationId,attr"`
	ArmyGroup            string `xml:"ArmyGroup,attr"`
}
type ArmyGroupXmls struct {
	Nodes []ArmyGroupXml `xml:"ArmyGroupXml"`
}

//阵容
type ArmyFormationXml struct {
	ArmyFormationId string `xml:"ArmyFormationId,attr"`
	Formations      string `xml:"Formations,attr"`
}
type ArmyFormationXmls struct {
	Nodes []ArmyFormationXml `xml:"ArmyFormationXml"`
}

//英雄
type HeroXml struct {
	HeroId               string `xml:"HeroId,attr"`
	HeroAtk              string `xml:"HeroAtk,attr"`
	HeroDef              string `xml:"HeroDef,attr"`
	HeroLife             string `xml:"HeroLife,attr"`
	HeroStrategy         string `xml:"HeroStrategy,attr"`
	HeroFirstAtkInterval string `xml:"HeroFirstAtkInterval,attr"`
	HeroAttackInterval   string `xml:"HeroAttackInterval,attr"`
}
type HeroXmls struct {
	Nodes []HeroXml `xml:"HeroXml"`
}

//兵种表
type SoldierXml struct {
	SoldierId               string `xml:"SoldierId,attr"`
	SoldierTypeId           string `xml:"SoldierTypeId,attr"`
	SoldierAtk              string `xml:"SoldierAtk,attr"`
	SoldierDef              string `xml:"SoldierDef,attr"`
	SoldierLife             string `xml:"SoldierLife,attr"`
	SoldierAtkType          string `xml:"SoldierAtkType,attr"`
	SoldierFirstAtkInterval string `xml:"SoldierFirstAtkInterval,attr"`
	SoldierAttackInterval   string `xml:"SoldierAttackInterval,attr"`
	SoldierAttackRange      string `xml:"SoldierAttackRange,attr"`
}
type SoldierXmls struct {
	Nodes []SoldierXml `xml:"SoldierXml"`
}

//兵系表
type SoldierTypeXml struct {
	SoldierTypeId    string `xml:"SoldierTypeId,attr"`
	AttackRange      string `xml:"AttackRange,attr"`
	MoveSpeed        string `xml:"MoveSpeed,attr"`
	FirstAtkInterval string `xml:"FirstAtkInterval,attr"`
	AttackInterval   string `xml:"AttackInterval,attr"`
}
type SoldierTypeXmls struct {
	Nodes []SoldierTypeXml `xml:"SoldierTypeXml"`
}

type PosBaseXml struct {
	PosId string `xml:"PosId,attr"`
	PosLX string `xml:"PosLX,attr"`
	PosLY string `xml:"PosLY,attr"`
	PosRX string `xml:"PosRX,attr"`
	PosRY string `xml:"PosRY,attr"`
}

type PosBaseXmls struct {
	Nodes []*PosBaseXml `xml:"PosBaseXml"`
}

//xml映射struct end

//逻辑映射struct start
type FightFactorStruct struct {
	FieldLength                int
	PauseTime                  int
	SputteringRange            int
	RangedAttackDetourDistance int
}

var FightFactor FightFactorStruct

type PosOffset struct {
	x float64
	y float64
}

type CloseCombatPosFactorStruct struct {
	PosOffsets [8]PosOffset
}

var CloseCombatPosFactor CloseCombatPosFactorStruct
var CloseCombatPosFactorSpec float64

var MoveSpeedFactor int

var ArmyGroupBase_map map[int]*ArmyGroupBase

type ArmyGroupBase struct {
	ArmyGroupId          int
	ArmyGroupFormationId int
	Armys                [4]*ArmyBase
}
type ArmyBase struct {
	HeroId     int
	SoldierId  int
	SoldierNum int
}

var ArmyFormationBase_map map[int]*ArmyFormationBase

type ArmyFormationBase struct {
	ArmyFormationId int
	Formations      [8]PosBase
}
type PosBase struct {
	PosX int
	PosY int
	Lx   float64
	Ly   float64
	Rx   float64
	Ry   float64
}

var PosBase_map map[int]*PosBase

var HeroBase_map map[int]*HeroBase

type HeroBase struct {
	HeroId               int
	HeroAtk              int
	HeroDef              int
	HeroLife             int
	HeroStrategy         int
	HeroFirstAtkInterval int
	HeroAttackInterval   int
}

var SoldierBase_map map[int]*SoldierBase

type SoldierBase struct {
	SoldierId               int
	SoldierTypeId           int
	SoldierAtk              int
	SoldierDef              int
	SoldierLife             int
	SoldierAtkType          int
	SoldierFirstAtkInterval int
	SoldierAttackInterval   int
	SoldierAttackRange      int
}

var SoldierTypeBase_map map[int]*SoldierTypeBase

type SoldierTypeBase struct {
	SoldierTypeId    int
	AttackRange      int
	MoveSpeed        int
	FirstAtkInterval int
	AttackInterval   int
}

//逻辑映射struct end

func init() {
	resinit()
}

func resinit() {
	initFightFactor()
	initArmyGroupXml()
	initArmyFormationXml()
	initHeroXml()
	initSoldierXml()
	initSoldierTypeXml()
	initPosBaseXml()
}

func initFightFactor() {
	util.ReadXml(&FightFactorsXml, "fightDemo/FightFactor.xml")

	FightFactor.FieldLength, _ = strconv.Atoi(FightFactorsXml.FightFactor.FieldLength)
	FightFactor.PauseTime, _ = strconv.Atoi(FightFactorsXml.FightFactor.PauseTime)
	FightFactor.SputteringRange, _ = strconv.Atoi(FightFactorsXml.FightFactor.SputteringRange)
	FightFactor.RangedAttackDetourDistance, _ = strconv.Atoi(FightFactorsXml.FightFactor.RangedAttackDetourDistance)

	posOffsets := strings.Split(FightFactorsXml.CloseCombatPosFactor.PosOffset, ";")
	if len(posOffsets) > 0 {
		for n, v := range posOffsets {
			offset := strings.Split(v, ",")
			if len(offset) == 2 {
				x, _ := strconv.Atoi(offset[0])
				y, _ := strconv.Atoi(offset[1])
				CloseCombatPosFactor.PosOffsets[n] = PosOffset{x: float64(x), y: float64(y)}
			}
		}
	}
	CloseCombatPosFactorSpec = CloseCombatPosFactor.PosOffsets[0].x
	if CloseCombatPosFactorSpec < 0 {
		CloseCombatPosFactorSpec = -CloseCombatPosFactorSpec
	}

	MoveSpeedFactor, _ = strconv.Atoi(FightFactorsXml.MoveSpeedFactor.MoveSpeed)
	fmt.Println("speed=", MoveSpeedFactor)
}

func initArmyGroupXml() {
	ArmyGroupBase_map = make(map[int]*ArmyGroupBase)
	armyGroupXmls := ArmyGroupXmls{}
	util.ReadXml(&armyGroupXmls, "fightDemo/ArmyGroup.xml")

	nodes := armyGroupXmls.Nodes
	for _, v := range nodes {
		agb := new(ArmyGroupBase)
		agb.ArmyGroupId, _ = strconv.Atoi(v.ArmyGroupId)
		agb.ArmyGroupFormationId, _ = strconv.Atoi(v.ArmyGroupFormationId)
		armys := strings.Split(v.ArmyGroup, "|")
		for n2, v2 := range armys {
			values := strings.Split(v2, ";")
			if len(values) != 3 {
				fmt.Println("initArmyGroupXml error")
			}
			ab := new(ArmyBase)
			heroId, _ := strconv.Atoi(values[0])
			soldierId, _ := strconv.Atoi(values[1])
			if heroId == 0 && soldierId == 0 { //全等于0的情况认为这个位置没有军队
				continue
			}
			ab.HeroId = heroId
			ab.SoldierId = soldierId
			ab.SoldierNum, _ = strconv.Atoi(values[2])
			agb.Armys[n2] = ab
		}
		ArmyGroupBase_map[agb.ArmyGroupId] = agb
	}
}

func initArmyFormationXml() {
	ArmyFormationBase_map = make(map[int]*ArmyFormationBase)
	armyFormationXmls := ArmyFormationXmls{}
	util.ReadXml(&armyFormationXmls, "fightDemo/ArmyFormation.xml")

	nodes := armyFormationXmls.Nodes
	for _, v := range nodes {
		afb := new(ArmyFormationBase)
		afb.ArmyFormationId, _ = strconv.Atoi(v.ArmyFormationId)
		values := strings.Split(v.Formations, ";")
		for n2, v2 := range values {
			poses := strings.Split(v2, ",")
			if len(poses) != 2 {
				fmt.Println("ArmyFormation.xml error")
			} else {
				posX, _ := strconv.Atoi(poses[0])
				posY, _ := strconv.Atoi(poses[1])
				posBase := PosBase{PosX: posX, PosY: posY}
				afb.Formations[n2] = posBase
			}
		}
		ArmyFormationBase_map[afb.ArmyFormationId] = afb
	}
}

func initHeroXml() {
	HeroBase_map = make(map[int]*HeroBase)
	heroXmls := HeroXmls{}
	util.ReadXml(&heroXmls, "fightDemo/Hero.xml")

	nodes := heroXmls.Nodes
	for _, v := range nodes {
		hb := new(HeroBase)
		hb.HeroId, _ = strconv.Atoi(v.HeroId)
		hb.HeroAtk, _ = strconv.Atoi(v.HeroAtk)
		hb.HeroDef, _ = strconv.Atoi(v.HeroDef)
		hb.HeroLife, _ = strconv.Atoi(v.HeroLife)
		hb.HeroStrategy, _ = strconv.Atoi(v.HeroStrategy)
		hb.HeroFirstAtkInterval, _ = strconv.Atoi(v.HeroFirstAtkInterval)
		hb.HeroAttackInterval, _ = strconv.Atoi(v.HeroAttackInterval)
		HeroBase_map[hb.HeroId] = hb
	}
}

func initSoldierXml() {
	SoldierBase_map = make(map[int]*SoldierBase)
	soldierXmls := SoldierXmls{}
	util.ReadXml(&soldierXmls, "fightDemo/Soldier.xml")

	nodes := soldierXmls.Nodes
	for _, v := range nodes {
		sb := new(SoldierBase)
		sb.SoldierId, _ = strconv.Atoi(v.SoldierId)
		sb.SoldierTypeId, _ = strconv.Atoi(v.SoldierTypeId)
		sb.SoldierAtk, _ = strconv.Atoi(v.SoldierAtk)
		sb.SoldierDef, _ = strconv.Atoi(v.SoldierDef)
		sb.SoldierLife, _ = strconv.Atoi(v.SoldierLife)
		sb.SoldierAtkType, _ = strconv.Atoi(v.SoldierAtkType)
		sb.SoldierFirstAtkInterval, _ = strconv.Atoi(v.SoldierFirstAtkInterval)
		sb.SoldierAttackInterval, _ = strconv.Atoi(v.SoldierAttackInterval)
		sb.SoldierAttackRange, _ = strconv.Atoi(v.SoldierAttackRange)
		SoldierBase_map[sb.SoldierId] = sb
	}
}

func initSoldierTypeXml() {
	SoldierTypeBase_map = make(map[int]*SoldierTypeBase)
	soldierTypeXmls := SoldierTypeXmls{}
	util.ReadXml(&soldierTypeXmls, "fightDemo/SoldierType.xml")

	nodes := soldierTypeXmls.Nodes
	for _, v := range nodes {
		stb := new(SoldierTypeBase)
		stb.SoldierTypeId, _ = strconv.Atoi(v.SoldierTypeId)
		stb.AttackRange, _ = strconv.Atoi(v.AttackRange)
		stb.MoveSpeed, _ = strconv.Atoi(v.MoveSpeed)
		stb.FirstAtkInterval, _ = strconv.Atoi(v.FirstAtkInterval)
		stb.AttackInterval, _ = strconv.Atoi(v.AttackInterval)
		SoldierTypeBase_map[stb.SoldierTypeId] = stb
	}
}

func initPosBaseXml() {
	PosBase_map = make(map[int]*PosBase)
	posBaseXmls := PosBaseXmls{}
	util.ReadXml(&posBaseXmls, "fightDemo/PosBase.xml")
	nodes := posBaseXmls.Nodes

	for _, v := range nodes {
		if v != nil {
			pos := &PosBase{}

			pos_id, _ := strconv.Atoi(v.PosId)

			pos.Lx, _ = strconv.ParseFloat(v.PosLX, 32)
			pos.Ly, _ = strconv.ParseFloat(v.PosLY, 32)
			pos.Rx, _ = strconv.ParseFloat(v.PosRX, 32)
			pos.Ry, _ = strconv.ParseFloat(v.PosRY, 32)

			PosBase_map[pos_id] = pos
		}
	}
}
