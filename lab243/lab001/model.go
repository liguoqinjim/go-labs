package main

import (
	"gorm.io/gorm"
	"time"
)

type Message struct {
	Id      int    `gorm:"column:id;type:int(11);primaryKey;autoIncrement;not null" json:"id"`
	UserId  int    `gorm:"column:user_id;type:int(11);not null" json:"userId"`
	Message string `gorm:"column:message;type:varchar(32);size:32" json:"message"`
}

//TableName 不支持动态变化，它会被缓存下来以便后续使用。想要使用动态表名，你可以使用 Scopes
func (Message) TableName() string {
	return "t_message"
}

func MessageTable(message *Message) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if message.UserId%2 == 0 {
			return db.Table("t_message_1")
		} else {
			return db.Table("t_message_2")
		}
	}
}

type User struct {
	Id         int       `gorm:"column:id;type:int(11);primaryKey;autoIncrement;not null" json:"id"`
	Username   string    `gorm:"column:username;type:varchar(32);size:32;not null" json:"username"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"createTime"`
}

func (User) TableName() string {
	return "t_user"
}
