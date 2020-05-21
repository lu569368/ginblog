package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/student/ginblog/common"
	"github.com/student/ginblog/datamodels"
	"github.com/student/ginblog/repositories"
	"github.com/student/ginblog/services"
)

func GetLogin(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login/login.html", gin.H{})
}
func PostLogin(ctx *gin.Context) {
	var user datamodels.LoginUser
	if err := ctx.Bind(&user); err != nil {
		common.Logger().WithFields(logrus.Fields{
			"name": "PostLogin",
		}).Error("用户数据绑定错误", "error")
		ctx.JSON(http.StatusBadRequest, gin.H{"status": 0, "msg": err.Error()})
		return
	}
	if !common.CaptchaVerify(ctx, user.Captcha) {
		ctx.JSON(http.StatusOK, gin.H{"status": 0, "msg": "验证码错误"})
		return
	}
	post := user.Password
	db := common.GormInit()
	userRepository := repositories.NewUserManager("users", db)
	u := services.NewUserService(userRepository)
	res := u.GetOneUser(&user, user.Username)
	isOK, err := common.ValidatePassword(post, res.Password)
	// fmt.Println(isOK)
	if err != nil || isOK != true {
		common.Logger().WithFields(logrus.Fields{
			"name": "PostRegister",
		}).Error("账户或密码错误", "error")
		ctx.JSON(http.StatusOK, gin.H{"status": 0, "msg": "账户或密码错误"})
		return
	}
	token, err := common.ReleaseToken(user)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"status": 0, "msg": "生成token失败"})
	}
	ctx.JSON(http.StatusOK, gin.H{"status": 1, "token": token, "msg": "登录成功"})
	// fmt.Println(post)
	// fmt.Println(res.Password)
}
