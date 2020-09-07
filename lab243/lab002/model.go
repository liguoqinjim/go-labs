package main

import "time"

type AuthCode struct {
	Id          int       `gorm:"column:id;type:int(10) unsigned;primary_key;AUTO_INCREMENT;not null" json:"id"`
	AuthCode    string    `gorm:"column:auth_code;type:varchar(32);size:32;not null" json:"authCode"`
	AuthType    int       `gorm:"column:auth_type;type:int(11);not null" json:"authType"`
	UserId      int       `gorm:"column:user_id;type:int(11);not null" json:"userId"`
	Invalid     int       `gorm:"column:invalid;type:tinyint(4);not null;default:0" json:"invalid"`
	GeneratorId int       `gorm:"column:generator_id;type:int(11);not null;default:0" json:"generatorId"`
	Price       int       `gorm:"column:price;type:int(11);not null;default:0" json:"price"`
	OrderId     int       `gorm:"column:order_id;type:int(11);not null;default:0" json:"orderId"`
	GenTime     time.Time `gorm:"column:gen_time;type:datetime;not null" json:"genTime"`
	ActiveTime  time.Time `gorm:"column:active_time;type:datetime;not null" json:"activeTime"`
	ExpireTime  time.Time `gorm:"column:expire_time;type:datetime;not null" json:"expireTime"`
	UpdateTime  time.Time `gorm:"column:update_time;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"updateTime"`
	Deleted     int       `gorm:"column:deleted;type:tinyint(4);not null;default:0" json:"deleted"`
}

type Team struct {
	Id           int       `gorm:"column:id;type:int(10) unsigned;primary_key;AUTO_INCREMENT;not null" json:"id"`
	TeamCode     string    `gorm:"column:team_code;type:varchar(32);size:32;not null" json:"teamCode"`
	TeamName     string    `gorm:"column:team_name;type:varchar(32);size:32;not null" json:"teamName"`
	TeamAvatar   string    `gorm:"column:team_avatar;type:varchar(128);size:128" json:"teamAvatar"`
	LeaderId     int       `gorm:"column:leader_id;type:int(11);not null" json:"leaderId"`
	TeamParentId int       `gorm:"column:team_parent_id;type:int(11);not null;default:0" json:"teamParentId"`
	CreateTime   time.Time `gorm:"column:create_time;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"createTime"`
	UpdateTime   time.Time `gorm:"column:update_time;type:datetime" json:"updateTime"`
}

func (Team) TableName() string {
	return "t_team"
}
