package main

import (
	"github.com/gin-gonic/gin"
	"github.com/student/ginblog/backend/controllers"
	"github.com/student/ginblog/common"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(common.LoggerToFile())
	r.Use(common.Session("topgoer"))
	r.LoadHTMLGlob("views/**/*")
	// r.Static("/assets", "./assets")
	r.GET("/api/auth/register", controllers.GetRegister)
	r.POST("/api/auth/register", controllers.PostRegister)
	r.GET("/api/auth/captcha", controllers.Captcha)
	r.GET("/api/auth/login", controllers.GetLogin)
	r.POST("/api/auth/login", controllers.PostLogin)
	r.GET("/api/auth/index", controllers.GetIndex)
	r.GET("/api/auth/user", controllers.GetUser)
	r.POST("/api/auth/user", controllers.PostUser)
	r.GET("/api/auth/user/:id", controllers.GetoneUser)
	r.PUT("/api/auth/user", controllers.PutUser)
	r.DELETE("/api/auth/user/:id", controllers.DeleteUser)
	return r
}
