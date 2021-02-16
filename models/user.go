// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package models

import (
	"github.com/ysicing/ext/logger/zlog"
	"gorm.io/gorm"
)

type User struct {
	Model
	Username string `gorm:"column:username" json:"username"`
	Password string `gorm:"column:password" json:"password"`
	Email    string `gorm:"column:email" json:"email"`
	Banned   bool   `gorm:"column:banned" json:"banned"`
	Token    string `gorm:"column:token" json:"token"`
}

func (User) TableName() string {
	return "system_user"
}

func init() {
	Migrate(User{})
}

func (u *User) New() *User {
	tx := GDB.Save(&u)
	if tx.Error != nil {
		zlog.Debug("添加人员save: %v, err: %v", u.Username, tx.Error.Error())
		return nil
	}
	var nw User
	if err := tx.Row().Scan(&nw); err != nil {
		zlog.Debug("添加人员scan: %v, err: %v", u.Username, err.Error())
		return nil
	}
	return &nw
}

func (u *User) Exist() bool {
	tx := GDB.Model(User{}).Where("username = ?", u.Username).Or("email = ?", u.Email).Find(&User{})
	if tx.Error != nil && tx.Error != gorm.ErrRecordNotFound {
		return false
	}
	if tx.RowsAffected > 0 {
		return true
	}
	return false
}
