package models

import "github.com/jinzhu/gorm"

// token 数据模型
type OauthToken struct {
	gorm.Model
	Token     string `gorm:"not null default '' comment('Token') VARCHAR(191)"`
	UserId    uint   `gorm:"not null default '' comment('UserId') VARCHAR(191)"`
	Secret    string `gorm:"not null default '' comment('Secret') VARCHAR(191)"`
	ExpressIn int64  `gorm:"not null default 0 comment('是否是标准库') BIGINT(20)"`
	Revoked   bool
}

// 创建 token
func (ot *OauthToken) OauthTokenCreate() (response Token) {
	Db.Create(ot)
	response = Token{ot.Token}
	return
}

type Token struct {
	Token string `json:"access_token"`
}

// 获取 access_token 信息
func GetOauthTokenByToken(token string) (ot *OauthToken) {
	ot = new(OauthToken)
	Db.Where("token =  ?", token).First(&ot)
	return
}
