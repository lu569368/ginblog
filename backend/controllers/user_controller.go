package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/student/ginblog/dto"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/student/ginblog/common"
	"github.com/student/ginblog/datamodels"
	"github.com/student/ginblog/repositories"
	"github.com/student/ginblog/services"
)

func GetUser(ctx *gin.Context) {
	var userArray []datamodels.User
	var countUser int64
	db := common.GormInit()
	userRepository := repositories.NewUserManager("users", db)
	u := services.NewUserService(userRepository)
	res, err := u.GetAllUser()
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range res {
		v1 := *v
		// dto.ToUserDto(v1)
		userArray = append(userArray, v1)
	}
	db.Table("users").Count(&countUser)
	fmt.Println(countUser)
	// fmt.Println(us)
	ctx.HTML(http.StatusOK, "user/index.html", gin.H{"userArray": userArray})
}
func PostUser(ctx *gin.Context) {
	var user datamodels.User
	if err := ctx.Bind(&user); err != nil {
		common.Logger().WithFields(logrus.Fields{
			"name": "PostUser",
		}).Error("", "error")
		ctx.JSON(http.StatusBadRequest, gin.H{"status": 0, "msg": err.Error()})
		return
	}
	passwordbyte, err := common.GeneratePassword(user.Password)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"status": 0, "msg": "密码加密失败"})
		return
	}
	user.Password = string(passwordbyte)
	db := common.GormInit()
	userRepository := repositories.NewUserManager("users", db)
	u := services.NewUserService(userRepository)
	res := u.InsertUser(&user)
	if res == 0 {
		ctx.JSON(http.StatusOK, gin.H{"status": 0, "msg": "新增用户失败"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": 1, "msg": "新增用户成功"})
}
func PutUser(ctx *gin.Context) {
	var user datamodels.User
	if err := ctx.Bind(&user); err != nil {
		common.Logger().WithFields(logrus.Fields{
			"name": "PutUser",
		}).Error("", "error")
		ctx.JSON(http.StatusBadRequest, gin.H{"status": 0, "msg": err.Error()})
		return
	}
	passwordbyte, err := common.GeneratePassword(user.Password)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"status": 0, "msg": "密码加密失败"})
		return
	}
	user.Password = string(passwordbyte)
	db := common.GormInit()
	userRepository := repositories.NewUserManager("users", db)
	u := services.NewUserService(userRepository)
	res := u.UpdateUser(&user)
	if res != nil {
		ctx.JSON(http.StatusOK, gin.H{"status": 0, "msg": "更新失败请重新尝试"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": 1, "mas": "更新成功"})
}
func GetoneUser(ctx *gin.Context) {
	id := ctx.Param("id")
	var user datamodels.User
	IntId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		common.Logger().WithFields(logrus.Fields{
			"name": "GetoneUser",
		}).Error("字符串转int60出错", "error")
	}
	db := common.GormInit()
	userRepository := repositories.NewUserManager("users", db)
	u := services.NewUserService(userRepository)
	res := u.GetOneUserById(&user, IntId)
	OneUser := dto.ToUserDto(*res)
	ctx.JSON(http.StatusOK, gin.H{"status": 1, "OneUser": OneUser})
}
func DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	IntId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		common.Logger().WithFields(logrus.Fields{
			"name": "GetoneUser",
		}).Error("字符串转int60出错", "error")
	}
	db := common.GormInit()
	userRepository := repositories.NewUserManager("users", db)
	u := services.NewUserService(userRepository)
	isOk := u.DeleteUserByID(IntId)
	if isOk != true {
		ctx.JSON(400, gin.H{"status": 1, "msg": "删除错误请重试！"})
	}
	ctx.JSON(http.StatusOK, gin.H{"status": 1, "msg": "删除成功"})
}
