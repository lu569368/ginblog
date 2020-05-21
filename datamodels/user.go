package datamodels

type User struct {
	ID       int64  `json:"id" sql:"id" form:"id" `
	Username string `json:"username" sql:"username" form:"username" binding:"required"`
	Password string `json:"password" sql:"password" form:"password" binding:"required"`
}

type AllUser struct {
	User
	ConfirmPassword string `form:"confirmPassword"  binding:"required"`
	Captcha         string `form:"captcha"  binding:"required"`
}
