package from

//登录表单
type UserLoginFrom struct {
	Email 		string 	`form:"email" binding:"required"`
	Password 	string	`form:"password" binding:"required"`
}

//注册表单
type UserRegisterFrom struct {
	Email 		string 	`form:"email" binding:"required"`
	Password 	string	`form:"password" binding:"required"`
	Name 		string	`form:"name" binding:"required"`
}

//修改密码
type UserPasswordChangeFrom struct {
	Email 			string 	`form:"email" binding:"required"`
	OldPassword 	string	`form:"old_password" binding:"required"`
	NewPassword 	string	`form:"new_password" binding:"required"`
}

//重置密码
type UserPasswordResetFrom struct {
	Email 		string 	`form:"email" binding:"required"`
	Code 		string 	`form:"code" binding:"required"`
	Password 	string	`form:"password" binding:"required"`
}