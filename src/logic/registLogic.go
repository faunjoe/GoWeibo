package login

import (
	"GoWeibo/src/model"
	"GoWeibo/util"
	"errors"
	"fmt"
	"strings"
	"time"
	"unicode/utf8"
	"weibo/db"

	"github.com/gin-gonic/gin"
)

// UserLogic 用户登录逻辑
type UserLogic struct{}

// DefaultUser 空Struct
var DefaultUser = UserLogic{}

// CreateUser 创建用户
func (userLogic UserLogic) CreateUser(c *gin.Context) (errMsg string, err error) {
	email := c.Query("email")
	password := c.Query("password")
	u := model.User{Email: email, Password: password}

	if u.Email == "" {
		return "", errors.New("邮箱不能为空")
	}

	if u.Password == "" {
		return "", errors.New("密码不能为空")
	}

	if emailExist := Emailexists(u.Email); emailExist {
		return "邮箱已经存在", errors.New("邮箱已经存在")
	}

	// 计算MD5
	u.Password = util.GenerateStringMD5(u.Password)

	insert := db.MasterDB.Debug().Create(&u).RecordNotFound()
	if insert {
		return "", errors.New("注册用户失败")
	}
	return "", nil
}

// Emailexists 邮箱是否已经存在
func Emailexists(email string) bool {
	user := model.User{}
	db.MasterDB.Where(&model.User{Email: email}).First(&user)
	fmt.Println(user.Email)
	if utf8.RuneCountInString(user.Email) > 0 {
		fmt.Println("数据库已经存在email: ", email)
		return true
	}
	fmt.Println("数据库不存在email: ", email)
	return false
}

// UserLogin 用户登录
func (userLogic UserLogic) UserLogin(c *gin.Context) (data interface{}, err error) {
	userName := c.PostForm("username")
	passWord := c.PostForm("password")
	if userName == "" || passWord == "" {
		return "", errors.New("账号/密码不能为空")
	}

	if !Emailexists(userName) {
		return "", errors.New("账号/密码错误")
	}

	// 查数据库验证账号密码是否正确
	// 计算加密后的password
	md5Str := util.GenerateStringMD5(passWord)

	user := model.User{}
	db.MasterDB.Where(&model.User{Email: userName, Password: md5Str}).First(&user)
	if (strings.Compare(userName, user.Email) == 0) && (strings.Compare(md5Str, user.Password) == 0) {
		// 密码正确
		fmt.Println(user.Email, "验证成功")
		token := GenerateToken(&user)
		// token 写入数据库
		db.MasterDB.Model(&user).Update("Token", token)
		data := map[string]interface{}{"token": token}
		return data, nil
	}
	return "", errors.New("账号密码错误")
}

// GenerateToken 生成token
func GenerateToken(u *model.User) string {
	md5Pre := u.UID + u.Password + u.Email + string(time.Now().Unix())
	return util.GenerateStringMD5(md5Pre)
}
