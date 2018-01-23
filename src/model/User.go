package model

// User 用户登陆信息
type User struct {
	UID           string `gorm:"column:id"`
	Email         string `gorm:"column:email"`
	Password      string `gorm:"column:password"`
	Status        int    `gorm:"column:status"`
	LastloginTime uint64 `gorm:"column:lastlogintime"`
	Token         string `gorm:"column:rememberToken"`
}

// TableName 返回ORM表名
func (U *User) TableName() string {
	return "users"
}
