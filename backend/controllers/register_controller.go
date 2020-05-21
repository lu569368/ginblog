package controllers

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/student/ginblog/datamodels"
	"github.com/student/ginblog/repositories"
	"github.com/student/ginblog/services"

	"github.com/gin-gonic/gin"
	"github.com/student/ginblog/common"
)

// type User struct {
// 	Username        string `form:"username"  binding:"required"`
// 	Pssword         string `form:"password"  binding:"required"`
// 	ConfirmPassword string `form:"confirmPassword"  binding:"required"`
// 	Captcha         string `form:"captcha"  binding:"required"`
// }

func GetRegister(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "register/register.html", gin.H{})
}
func PostRegister(ctx *gin.Context) {
	var user datamodels.AllUser
	if err := ctx.Bind(&user); err != nil {
		common.Logger().WithFields(logrus.Fields{
			"name": "PostRegister",
		}).Error("用户数据绑定错误", "error")
		ctx.JSON(http.StatusBadRequest, gin.H{"status": 0, "msg": err.Error()})
		return
	}
	if user.User.Password != user.ConfirmPassword {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": 0, "msg": "两次密码输入不一致"})
		return
	}
	if !common.CaptchaVerify(ctx, user.Captcha) {
		ctx.JSON(http.StatusOK, gin.H{"status": 0, "msg": "验证码错误"})
		return
	}
	db := common.GormInit()
	userRepository := repositories.NewUserManager("users", db)
	u := services.NewUserService(userRepository)
	res := u.GetUserByName(&user.User, user.User.Username)
	if res != 0 {
		ctx.JSON(http.StatusOK, gin.H{"status": 0, "msg": "用户名已存在"})
		return
	}
	passwordbyte, err := common.GeneratePassword(user.User.Password)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"status": 0, "msg": "密码加密失败"})
		return
	}
	user.User.Password = string(passwordbyte)
	result := u.InsertUser(&user.User)
	if result == 0 {
		ctx.JSON(http.StatusOK, gin.H{"status": 0, "msg": "注册失败请重新尝试"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": 1, "msg": "注册成功"})
}
func Captcha(ctx *gin.Context) {
	common.Captcha(ctx, 4)
}
