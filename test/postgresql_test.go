package test

import (
	"fmt"
	"myproject3/config"
	"myproject3/model"
	"testing"
)

func TestConn(t *testing.T) {
	db := config.InitDB()
	defer db.Close()
	var user model.UserInfo
	db.SingularTable(true)
	if err := db.Where("uid=?", 1).First(&user).Error; err != nil {
		fmt.Print("fail")
		return
	} else {
		fmt.Print("success")
		return
	}
}
func TestJoin(t *testing.T) {
	type Result struct {
		Uname    string `json:"uname"`
		Isreal   int    `json:"isreal"`
		Sex      int    `json:"sex"`
		Portrait string `json:"portrait"`
		Uid      int    ` json:"uid"`
		Descr    string ` json:"descr"`
		Label    string ` json:"label"`
		Fileinfo string ` json:"fileinfo"`
		Ct       string `json:"ct"`
	}

	db := config.GetDB()

	var result Result
	sql := "select i.uname as uname,i.isreal_name as isreal,i.sex as sex,i.portrait as portrait,i.uid as uid,f.descr as descr,f.label as label,f.fileinfo as fileinfo ,f.create_time as ct from user_info as i left join user_feedback as f on i.uid=f.uid where i.uid=?"
	db.Raw(sql, 1).Find(&result)

}

func TestScan(t *testing.T) {
	type Result struct {
		Uname      string `gorm:"uname"`
		IsrealName int64  `gorm:"isreal_name"`
		Sex        int64  `gorm:"sex"`
		Portrait   string `gorm:"portrait"`
		Uid        int64  ` gorm:"uid"`
		Descr      string ` gorm:"descr"`
		Label      string ` gorm:"label"`
		Fileinfo   string ` gorm:"fileinfo"`
		CreateTime string `gorm:"create_time"`
	}

	db := config.GetDB()

	var result Result
	db.SingularTable(true)
	sql := "select i.uname as uname,i.isreal_name as isreal_name,i.sex as sex,i.portrait as portrait,i.uid as uid,f.descr as descr,f.label as label,f.fileinfo as fileinfo ,f.create_time as create_time from user_info as i left join user_feedback as f on i.uid=f.uid where i.uid=?"
	db.Raw(sql, 1).Scan(&result)

}
func TestSa(t *testing.T) {
	db := config.GetDB()
	type Result struct {
		Sex string `json:"sex" gorm:"sex"`
		UId string `json:"uid" gorm:"uid"`
	}
	var result Result
	db.SingularTable(true)
	db.Table("user_info").Select("user_info.sex as sex ,user_info.uid as uid ").First(&result)
}
func TestConne(t *testing.T) {
	type Result struct {
		Uname      string `json:"Uname" gorm:"column:uname;" `
		IsrealName int    `json:"IsrealName" gorm:"column:isreal_name;"`
		Sex        int    `json:"Sex" gorm:"column:sex;"`
		Portrait   string `json:"Portrait" gorm:"column:portrait;"`
		Uid        int    `json:"Uid" gorm:"column:uid;"`
		Username   string `json:"Username" gorm:"column:username;"`
		Pwd        string `json:"Pwd" gorm:"column:pwd;"`
	}
	db := config.GetDB()
	defer db.Close()
	var result Result
	err := db.Table("user_info").Select("user_info.uname as uname ,user_info.isreal_name as isreal_name,user_info.sex as sex ,user_info.portrait as portrait,user_info.uid as uid,users.username as username ,users.pwd as pwd").
		Joins("left join  users ON users.id = user_info.uid").Where("users.id=1").First(&result).Error
	if err != nil {
		fmt.Println("error")
	}

}
func TestConn2(t *testing.T) {
	db := config.GetDB()
	defer db.Close()
	var user model.Users
	// db.Table("users").Select("users.id as id,users.username as username,users.pwd as pwd").Where("users.id=1").First(&user)
	err := db.Where("username = ?", "admin").Where("pwd=?", "admin").First(&user).Error
	if err != nil {
		fmt.Println("smd")
	}
}
