package main

import (
	"fmt"
)

type TransLog interface {
	Commit()
	RollBack(*DataBase)
}

type DataBase struct {
	transLogs []TransLog
	player    map[int]*Player //key是无关逻辑的主键
}

type Player struct {
	Id   int
	Pid  int
	Gold int
}

func NewDataBase() *DataBase {
	return &DataBase{player: make(map[int]*Player)}
}
func (db *DataBase) Transaction(trans func()) {
	defer func() {
		if err := recover(); err != nil {
			for i := len(db.transLogs) - 1; i >= 0; i-- {
				db.transLogs[i].RollBack(db)
			}
		} else {
			for _, tl := range db.transLogs {
				tl.Commit()
			}
		}
		//清空translog
		db.transLogs = db.transLogs[0:0]
	}()

	trans()
}

func (db *DataBase) LookupPlayer(id int) *Player {
	return db.player[id]
}
func (db *DataBase) InsertPlayer(p *Player) {
	db.player[p.Id] = p

	//插入事务组
	db.transLogs = append(db.transLogs, &PlayerTransLog{Type: INSERT, New: p})
}
func (db *DataBase) DeletePlayer(p *Player) {
	old := db.player[p.Id]
	delete(db.player, p.Id)
	db.transLogs = append(db.transLogs, &PlayerTransLog{Type: DELETE, Old: old})
}
func (db *DataBase) UpdatePlayer(p *Player) {
	old := db.player[p.Id]
	db.player[p.Id] = p
	db.transLogs = append(db.transLogs, &PlayerTransLog{Type: UPDATE, Old: old, New: p})
}

type TransType int

const (
	INSERT TransType = iota
	DELETE
	UPDATE
)

type PlayerTransLog struct {
	Type TransType
	Old  *Player
	New  *Player
}

func (transLog *PlayerTransLog) Commit() {
	switch transLog.Type {
	case INSERT:
		fmt.Printf("insert into player(id,pid,gold) values(%d,%d,%d)\n", transLog.New.Id, transLog.New.Pid, transLog.New.Gold)
	case DELETE:
		fmt.Printf("delete from player where id = %d\n", transLog.Old.Id)
	case UPDATE:
		fmt.Printf("update player set pid=%d,gold=%d where id=%d\n", transLog.New.Pid, transLog.New.Gold, transLog.New.Id)
	}
}
func (transLog *PlayerTransLog) RollBack(db *DataBase) {
	fmt.Println("rollback")
	switch transLog.Type {
	case INSERT:
		delete(db.player, transLog.New.Id)
	case DELETE:
		db.player[transLog.Old.Id] = transLog.Old
	case UPDATE:
		db.player[transLog.Old.Id] = transLog.Old
	}
}

func main() {
	db := NewDataBase()

	//insert
	db.Transaction(func() {
		db.InsertPlayer(&Player{Id: 1, Pid: 10001, Gold: 1000})
		db.InsertPlayer(&Player{Id: 2, Pid: 10002, Gold: 2000})
	})

	fmt.Println("insert:", db.player)
	fmt.Println(db.player[1], db.player[2])

	//delete
	db.Transaction(func() {
		p := db.LookupPlayer(1)
		db.DeletePlayer(p)
	})
	fmt.Println("delete:", db.player)
	fmt.Println(db.player[2])

	//update
	db.Transaction(func() {
		p := db.LookupPlayer(2)
		np := &Player{Id: p.Id, Pid: p.Pid, Gold: p.Gold + 200}
		db.UpdatePlayer(np)
	})
	fmt.Println("update:", db.player)
	fmt.Println(db.player[2])

	//update rollback
	db.Transaction(func() {
		p := db.LookupPlayer(2)
		np := &Player{Id: p.Id, Pid: p.Pid, Gold: p.Gold + 200}
		db.UpdatePlayer(np)

		panic("force error")
	})
	fmt.Println("update rollback:", db.player)
	fmt.Println(db.player[2])
}
