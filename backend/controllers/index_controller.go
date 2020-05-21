package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetIndex(ctx *gin.Context) {
	// username, err := ctx.Get("username")
	// if err != true {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"status": 0, "msg": "获取用户名错误"})
	// }
	ctx.HTML(http.StatusOK, "index/index.html", gin.H{})
}
