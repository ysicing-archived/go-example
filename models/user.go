// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package models

import (
	"github.com/ergoapi/util/color"
	"github.com/ergoapi/util/exhash"
	"github.com/ergoapi/zlog"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type User struct {
	Model
	Username string `gorm:"column:username" json:"username"`
	Password string `gorm:"column:password" json:"password"`
	Email    string `gorm:"column:email" json:"email"`
	Banned   bool   `gorm:"column:banned" json:"banned"`
	Token    string `gorm:"column:token" json:"token"`
	Role string `json:"role"`
}

func (User) TableName() string {
	return "system_user"
}

func init() {
	Migrate(User{})
}

func (u *User) Save() error  {
	var uu User
	err := GDB.Model(User{}).Where("username = ?").Last(&uu).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if uu.ID > 0 {
		uu.Role = u.Role
		uu.Banned = u.Banned
		uu.Email = u.Email
		return GDB.Save(&uu).Error
	}
	return GDB.Create(&u).Error
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
	tx := GDB.Model(User{}).Where("username = ?", u.Username).Find(&User{})
	if tx.Error != nil && tx.Error != gorm.ErrRecordNotFound {
		return false
	}
	if tx.RowsAffected > 0 {
		return true
	}
	return false
}

// InitAdmin init
func InitAdmin() {
	val, err := ConfigsGet("initadmin")
	if err != nil {
		zlog.Fatal("cannot query initadmin", err)
	}
	if val != "" {
		zlog.Info(color.SGreen("exist initadmin %v success...", val))
		return
	}
	user := viper.GetString("server.admin.user")
	adminuser := User{
		Username: user,
		Password: exhash.MD5(viper.GetString("server.admin.pass")),
		Email:    viper.GetString("server.admin.mail"),
		Banned:   false,
		Token:    exhash.GenUUIDForUser(user),
	}
	if err := adminuser.Save(); err != nil {
		zlog.Fatal("init admin in mysql", err)
	}
	err = ConfigsSet("initadmin", "done")
	if err != nil {
		zlog.Fatal("init initadmin in mysql", err)
	}
	zlog.Info(color.SGreen("init  admin %v success...", user))
}