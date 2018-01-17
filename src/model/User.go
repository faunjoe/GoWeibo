package model

import (
	"GoWeibo/src/db"
	"fmt"

	"github.com/jinzhu/gorm"
)

// User 用户信息
type User struct {
	gorm.Model
	Email         string `gorm:"email"`
	Password      string `gorm:"password"`
	Status        int    `gorm:"status"`
	LastloginTime uint64 `gorm:"lastlogintime"`
}

// Emailexists 邮箱是否已经存在
func (u *User) Emailexists(email string) bool {
	b := db.MasterDB.Not("email", email).First(&u)
	fmt.Println(b)
	return true
}
