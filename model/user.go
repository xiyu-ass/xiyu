package model

import (
	"encoding/json"
	"myproject3/utils"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type Users struct {
	// Model
	ID       int64  ` gorm:"type:int;primary_key"`                                         // 用户唯一标识
	Username string `form:"username" binding:"required" gorm:"type:varcahr(255);not null"` //账号
	Password string `form:"password" binding:"required" gorm:"colunm:pwd"`                 //密码
}

type UserInfo struct {
	ID         int64  `json:"id" gorm:"primary_key;"`      //用户信息唯一标识
	Uname      string `gorm:"column:uname;not null"`       //用户姓名
	IsRealName int64  `gorm:"column:isreal_name;not null"` //是否实名
	Sex        int64  `gorm:"column:sex;not null"`         //性别
	Portrait   string `gorm:"column:portrait;"`            //头像
	UId        uint   //用户id
}

type UserFeedback struct {
	ID         int64      `gorm:"type:int;primary_key"` //用户反馈信息唯一标识
	Descr      string     `gorm:"column:descr;not null"`
	Label      string     `gorm:"column:label;not null"`
	CreateTime *time.Time `gorm:"column:create_time;not null"`
	UId        int64      `gorm:"column:uid;"`
	FileInfo   string     `gorm:"column:fileinfo;"`
}

func GetUserSession(c *gin.Context) *Users {
	//创建session
	session := sessions.Default(c)
	//获取session的值
	s := session.Get("login")
	if s == nil {
		return nil
	}

	key, ok := s.([]byte)
	if !ok {
		return nil
	}

	user := &Users{}
	//把json数据转为结构体
	if err := json.Unmarshal([]byte(key), user); err != nil {
		utils.Fail(c, "数据转化异常", nil)
		return nil
	}
	return user

}
