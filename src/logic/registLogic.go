package login

import (
	"GoWeibo/src/model"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"unicode/utf8"
	"weibo/db"

	"github.com/gin-gonic/gin"
)

type UserLogic struct{}

var DefaultUser = UserLogic{}

// CreateUser 创建用户
func (self UserLogic) CreateUser(e *gin.Engine) (errMsg string, err error) {
	email := e.Query("email")
	password := e.Query("password")
	u := model.User{Email: email, Password: password}

	if u.Email == "" {
		return nil, errors.New("邮箱不能为空")
	}

	if u.Password == "" {
		return nil, errors.New("密码不能为空")
	}

	if emailExist := Emailexists(u.Email); emailExist {
		return "邮箱已经存在", errors.New("邮箱已经存在")
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
