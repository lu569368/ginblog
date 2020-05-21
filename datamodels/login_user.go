package datamodels

type LoginUser struct {
	ID       int64  `json:"id" sql:"id" form:"id" `
	Username string `json:"username" sql:"username" form:"username" binding:"required"`
	Password string `json:"password" sql:"password" form:"password" binding:"required"`
	Captcha  string `form:"captcha"  binding:"required"`
}
