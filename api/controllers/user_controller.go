package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/liu578101804/distributed-authentication/services"
	"github.com/liu578101804/distributed-authentication/api/from"
	"time"
)

type UserController struct {
	BaseController
	userService services.IUserService
	jwtService 	services.IJwtService
}

func (u *UserController) PostLogin(c *gin.Context) {
	var request from.UserLoginFrom
	if err := c.ShouldBind(&request);err != nil{
		u.ResFailJson(err.Error(),300, c)
		return
	}
	user,err := u.userService.Login(request.Email, request.Password)
	if err != nil {
		u.ResFailJson(err.Error(),301, c)
		return
	}

	//生成token
	token, err := u.jwtService.CreateToken(map[string]interface{}{
		"expiration_time": time.Now().AddDate(0,0,7).Format("2006-01-02 15:04:05"),
		"open_id": user.OpenId,
		"email": user.Email,
		"name": user.Name,
	})
	if err != nil {
		u.ResFailJson(err.Error(),302, c)
		return
	}

	u.ResSuccessJson(gin.H{
		"token": token,
	},c)
}

func (u *UserController) PostRegister(c *gin.Context) {
	var request from.UserRegisterFrom
	if err := c.ShouldBind(&request);err != nil{
		u.ResFailJson(err.Error(),300, c)
		return
	}
	user,err := u.userService.Register(request.Email, request.Password, request.Name)
	if err != nil {
		u.ResFailJson(err.Error(),301, c)
		return
	}
	u.ResSuccessJson(gin.H{
		"user": user,
	},c)
}

func (u *UserController) PostPasswordChange(c *gin.Context) {
	var request from.UserPasswordChangeFrom
	if err := c.ShouldBind(&request);err != nil{
		u.ResFailJson(err.Error(),300, c)
		return
	}
	err := u.userService.ChangePassword(request.Email, request.OldPassword, request.NewPassword)
	if err != nil {
		u.ResFailJson(err.Error(),301, c)
		return
	}
	u.ResSuccessJson(gin.H{
	},c)
}

func (u *UserController) PostPasswordReset(c *gin.Context) {
	var request from.UserPasswordResetFrom
	if err := c.ShouldBind(&request);err != nil{
		u.ResFailJson(err.Error(),300, c)
		return
	}
	err := u.userService.ReSetPassword(request.Email,request.Code,request.Password)
	if err != nil {
		u.ResFailJson(err.Error(),301, c)
		return
	}
	u.ResSuccessJson(gin.H{
	},c)
}

func NewUserController(userService services.IUserService, jwtService services.IJwtService) *UserController {
	return &UserController{
		userService: userService,
		jwtService: jwtService,
	}
}