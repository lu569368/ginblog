package common

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCaptcha(t *testing.T) {
	r := gin.Default()
	r.GET("/register", CeCaptcha)
}
func CeCaptcha(ctx *gin.Context) {
	Captcha(ctx, 4)
}
