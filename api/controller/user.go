package controller

import (
	"encoding/json"
	"errors"
	"flag"
	"myproject3/config"
	"myproject3/model"
	"myproject3/utils"
	"myproject3/utils/logger"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// 用户登录
func UserLogin(c *gin.Context) {
	params := &struct {
		Username string `form:"username"   binding:"required"`
		Password string `form:"password"   binding:"required"`
	}{}

	if err := c.ShouldBind(params); err != nil {
		utils.Fail(c, "数据绑定出错", nil)
		return
	}

	var user model.Users
	DB := config.GetDB()

	err := DB.Where("username = ?", params.Username).Where("pwd=?", params.Password).First(&user).Error
	if err != nil {
		utils.Fail(c, "账号或密码错误", nil)
		return
	} else {
		// session := sessions.Default(c)
		// session.Set("login", &user)
		// session.Save()
		session := sessions.Default(c)
		key, _ := json.Marshal(&user)
		session.Set("login", key)
		session.Save()
		utils.Success(c, "登录成功", session)
		return
	}

}

// 用户首页信息
func UserIndex(c *gin.Context) {
	loginUser := model.GetUserSession(c)

	type Result struct {
		Uname      string     `json:"Uname" gorm:"column:uname;"`
		IsrealName int64      `json:"IsrealName" gorm:"column:isreal_name;"`
		Sex        int64      `json:"Sex" gorm:"column:sex;"`
		Portrait   string     `json:"Portrait" gorm:"column:portrait;"`
		Uid        int64      `json:"Uid" gorm:"column:uid;"`
		Descr      string     `json:"Descr" gorm:"column:descr;"`
		Label      string     `json:"Label" gorm:"column:label;"`
		Fileinfo   string     `json:"Fileinfo" gorm:"column:fileinfo;"`
		CreateTime *time.Time `json:"CreateTime" gorm:"column:create_time;"`
	}
	db := config.GetDB()
	var result Result
	err := db.Table("user_info").Select("user_info.uname as uname ,user_info.isreal_name as isreal_name,user_info.sex as sex ,user_info.portrait as portrait,user_info.uid as uid,user_feedback.descr as descr ,user_feedback.label as label ,user_feedback.uid as uid ,user_feedback.fileinfo as fileinfo,user_feedback.create_time as create_time").
		Joins("left join  user_feedback ON user_feedback.uid = user_info.uid").Where("user_feedback.uid=?", loginUser.ID).First(&result).Error

	if err != nil {
		utils.Fail(c, "查询失败", nil)
		return
	}
	utils.Success(c, "查询成功", &result)
	return

}

// 用户反馈
func UserFeedback(c *gin.Context) {
	userLogin := model.GetUserSession(c)

	//1.获取上传的文件
	file, err := c.FormFile("file")
	if err == nil {
		//2.获取后缀名，判断类型是否正确 .jpg .png .gif .jpeg
		extName := path.Ext(file.Filename)
		allowExtmap := map[string]bool{
			".jpg":  true,
			".png":  true,
			",jpeg": true,
		}

		if _, ok := allowExtmap[extName]; !ok {
			logger.PanicError(errors.New("文件类型不合法"), "上传错误", false)
			utils.Fail(c, "文件类型不合法", nil)
			return
		}

		//创建图片保存目录,Linux下需设置权限（0755可读可写）
		currentTime := time.Now().Format("2006010")
		//使用flag定义路径字符常量
		dir := flag.String("b", "/home/xiyu/GOCODE/src/myproject3/upload/"+currentTime, "file name")
		utils.Success(c, "ok", dir)
		//生成目录文件夹
		if err := os.MkdirAll(*dir, 0755); err != nil {
			logger.PanicError(err, "上传失败", false)
			utils.Fail(c, "上传失败", dir)
			return
		}
		//4.生成文件名字
		fileUnixName := strconv.FormatInt(time.Now().UnixMicro(), 10)

		//上传文件
		saveDir := path.Join(*dir, fileUnixName+extName)

		err := c.SaveUploadedFile(file, saveDir)
		if err != nil {
			logger.PanicError(err, "上传错误", false)
			utils.Fail(c, "文件保存失败", nil)
			return
		}
		descr := c.PostForm("descr")
		label := c.PostForm("label")
		//将文件路径保存到数据库
		var user model.UserFeedback
		db := config.GetDB()
		db.SingularTable(true)

		user.Descr = descr
		user.Label = label
		user.UId = userLogin.ID
		times := time.Now()
		user.CreateTime = &times
		user.FileInfo = saveDir

		db.Save(&user)

		if err != nil {
			utils.Fail(c, "插入数据失败", nil)
			return
		}
		utils.Success(c, "上传成功", &user)
		return
	}
}

// 测试跨域
func ListUser(c *gin.Context) {
	db := config.GetDB()
	var user model.Users
	if err := db.Where("id=?", 1).First(&user); err != nil {
		utils.Fail(c, "错误", nil)
		return
	}
	utils.Success(c, "成功", nil)
	return
}
