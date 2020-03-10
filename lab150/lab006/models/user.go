package models

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/iris-contrib/middleware/jwt"
	"github.com/jameskeane/bcrypt"
	"github.com/jinzhu/gorm"
	"time"
)

// 用户数据模型
type User struct {
	gorm.Model
	Name     string `gorm:"not null VARCHAR(191)"`
	Username string `gorm:"unique;VARCHAR(191)"`
	Password string `gorm:"not null VARCHAR(191)"`
}

// 根据用户名查询用户
func UserAdminCheckLogin(username string) *User {
	user := new(User)
	IsNotFound(Db.Where("username = ?", username).First(user).Error)
	return user
}

// 检查登陆用户，并生成登陆凭证 token
func CheckLogin(username, password string) (response Token, status bool, msg string) {
	user := UserAdminCheckLogin(username)
	if user.ID == 0 {
		msg = "用户不存在"
		return
	} else {
		if ok := bcrypt.Match(password, user.Password); ok {
			token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"exp": time.Now().Add(time.Hour * time.Duration(1)).Unix(),
				"iat": time.Now().Unix(),
			})
			tokenString, _ := token.SignedString([]byte("HS2JDFKhu7Y1av7b"))
			oauthToken := new(OauthToken)
			oauthToken.Token = tokenString
			oauthToken.UserId = user.ID
			oauthToken.Secret = "secret"
			oauthToken.Revoked = false
			oauthToken.ExpressIn = time.Now().Add(time.Hour * time.Duration(1)).Unix()
			oauthToken.CreatedAt = time.Now()
			response = oauthToken.OauthTokenCreate()
			status = true
			msg = "登陆成功"
			return
		} else {
			msg = "用户名或密码错误"
			return
		}
	}
}

// 作废token
func UpdateOauthTokenByUserId(userId uint) (ot *OauthToken) {
	Db.Model(ot).Where("revoked = ?", false).
		Where("user_id = ?", userId).
		Updates(map[string]interface{}{"revoked": true})
	return
}

// 登出用户
func UserAdminLogout(userId uint) bool {
	ot := UpdateOauthTokenByUserId(userId)
	return ot.Revoked
}

func CreateUser() (user *User) {
	salt, _ := bcrypt.Salt(10)
	hash, _ := bcrypt.Hash("password", salt)
	user = &User{
		Username: "username",
		Password: hash,
		Name:     "name",
	}
	if err := Db.Create(user).Error; err != nil {
		color.Red(fmt.Sprintf("CreateUserErr:%s \n ", err))
	}
	return
}
