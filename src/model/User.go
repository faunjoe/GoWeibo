package model

import (
	"GoWeibo/src/db"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"unicode/utf8"
)

// User 用户信息
type User struct {
	Email         string `gorm:"column:email"`
	Password      string `gorm:"column:password"`
	Status        int    `gorm:"column:status"`
	LastloginTime uint64 `gorm:"column:lastlogintime"`
	Token         string `gorm:"column:rememberToken"`
}

// TableName 返回ORM表名
func (User) TableName() string {
	return "users"
}

// Emailexists 邮箱是否已经存在
func Emailexists(email string) bool {
	user := User{}
	db.MasterDB.Where(&User{Email: email}).First(&user)
	fmt.Println(user.Email)
	if utf8.RuneCountInString(user.Email) > 0 {
		fmt.Println("数据库已经存在email: ", email)
		return true
	}
	fmt.Println("数据库不存在email: ", email)
	return false
}

// InsertUser 插入用户
func InsertUser(u *User) error {
	b := Emailexists(u.Email)

	if u.Email == "" {
		return errors.New("邮箱不能为空")
	}

	if b {
		return errors.New("用户已经存在")
	}

	if u.Password == "" {
		return errors.New("密码不能为空")
	}

	// 计算MD5
	md5Pre := u.Password
	md5 := md5.New()
	io.WriteString(md5, md5Pre)
	md5Later := md5.Sum(nil)
	u.Password = hex.EncodeToString(md5Later)
	fmt.Println("md5值:", u.Password)
	insert := db.MasterDB.Debug().Create(&u).RecordNotFound()
	if insert {
		return errors.New("注册用户失败")
	}
	return nil
}
