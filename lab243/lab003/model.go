package main

import (
	"time"
)

type User struct {
	Id         int       `gorm:"column:id;type:int(11);primaryKey;autoIncrement;not null" json:"id"`
	Username   string    `gorm:"column:username;type:varchar(32);size:32;not null" json:"username"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"createTime"`
}

func (User) TableName() string {
	return "t_user"
}
