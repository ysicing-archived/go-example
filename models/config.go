// Copyright (c) 2022 ysicing All rights reserved.
// Use of this source code is governed by WTFPL LICENSE
// license that can be found in the LICENSE file.

package models

import (
	"fmt"
	"os"
	"time"

	"github.com/ergoapi/util/color"
	"github.com/ergoapi/util/exhash"
	"github.com/ergoapi/util/rand"
	"github.com/ergoapi/zlog"
	"gorm.io/gorm"
)

// Configs 配置
type Configs struct {
	Ckey string
	Cval string
}

func init() {
	Migrate(Configs{})
}

func (c *Configs) Save() error {
	return GDB.Where("ckey = ?", c.Ckey).Save(c).Error
}

// ConfigsGet 获取配置
func ConfigsGet(ckey string) (string, error) {
	var obj Configs
	has := GDB.Model(Configs{}).Where("ckey=?", ckey).Last(&obj)
	if has.Error != nil && has.Error != gorm.ErrRecordNotFound {
		return "", has.Error
	}

	if has.RowsAffected == 0 {
		return "", nil
	}

	return obj.Cval, nil
}

// ConfigsSet 添加配置
func ConfigsSet(ckey, cval string) error {
	var obj Configs
	has := GDB.Model(Configs{}).Where("ckey=?", ckey).Last(&obj)
	if has.Error != nil && has.Error != gorm.ErrRecordNotFound {
		return has.Error
	}
	var err error
	if has.RowsAffected == 0 {
		err = GDB.Create(&Configs{
			Ckey: ckey,
			Cval: cval,
		}).Error
	} else {
		obj.Cval = cval
		err = obj.Save()
	}
	return err
}

// InitSalt gen salt
func InitSalt() {
	val, err := ConfigsGet("salt")
	if err != nil {
		zlog.Fatal("cannot query salt", err)
	}
	if val != "" {
		zlog.Info(color.SGreen("exist salt %v success...", val))
		return
	}
	content := fmt.Sprintf("%s%d%d%s", rand.RandLetters(6), os.Getpid(), time.Now().UnixNano(), rand.RandLetters(6))
	salt := exhash.MD5(content)
	err = ConfigsSet("salt", salt)
	if err != nil {
		zlog.Fatal("init salt in mysql", err)
	}
	zlog.Info(color.SGreen("create salt %v success...", salt))
}
