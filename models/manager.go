package models

//管理员表

import (
	"github.com/alexedwards/argon2id"
)

type Manager struct { // 结构体首字母大写, 和数据库表名对应, 默认访问数据表users, 可以设置访问数据表的方法
	Id       int
	Username string
	Password string
	Mobile   string
	Email    string
	Status   int
	RoleId   int
	AddTime  int
	IsSuper  int
	Role     Role `gorm:"foreignKey:RoleId;references:Id"` // 配置关联关系
}

// 配置数据库操作的表名称
func (Manager) TableName() string {
	return "manager"
}

// AdminSession is the non-sensitive projection stored in the admin session.
type AdminSession struct {
	Id       int
	Username string
	RoleId   int
	IsSuper  int
}

func HashPassword(password string) (string, error) {
	return argon2id.CreateHash(password, argon2id.DefaultParams)
}

func VerifyPassword(password, hash string) bool {
	// Keep compatibility with the legacy 32-character MD5 hashes.
	if len(hash) == 32 && Md5(password) == hash {
		return true
	}

	ok, err := argon2id.ComparePasswordAndHash(password, hash)
	return err == nil && ok
}
