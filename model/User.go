package model

// 用户模型

import (
	"encoding/base64"
	"log"

	"github.com/TobiasKruzhor/crawler/utils/errmsg"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/scrypt"
)

type User struct {
	gorm.Model
	Username string `gorm:"type: varchar(20);not null" json:"username"`
	Password string `gorm:"type: varchar(20);not null" json:"password"`
	Role     int    `gorm:"type: int" json:"role"`
}

// 查询用户是否存在
func CheckUser(name string) (code int) {
	var users User
	db.Select("id").Where("username = ?", name).First(&users)
	if users.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCSE
}

// 新增用户
func CreateUser(data *User) int {
	// data.Password = ScryptPw(data.Password)
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR //400
	}
	return errmsg.SUCCSE
}

// 查询用户列表
func GetUsers(pageSize int, pageNum int) []User {
	var users []User
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return users
}

// 编辑用户

// 删除用户

// 密码加密
func (u *User) BeforeSave() {
	u.Password = ScryptPw(u.Password)
}

func ScryptPw(password string) string {
	const KeyLen = 10
	salt := []byte{12, 32, 4, 6, 66, 22, 222, 11}

	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	fpw := base64.StdEncoding.EncodeToString(HashPw)
	return fpw
}
