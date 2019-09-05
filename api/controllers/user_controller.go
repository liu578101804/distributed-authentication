package controllers

import (
	"github.com/liu578101804/distributed-authentication/services"
	"github.com/gin-gonic/gin"
	"github.com/liu578101804/distributed-authentication/datamodels"
)

type UserController struct {
	BaseController
	userService services.IUserService
}

func (u UserController) PostLogin(c *gin.Context) {
	var request datamodels.UserLoginFrom
	if err := c.ShouldBind(&request);err != nil{
		u.ResFailJson(err.Error(),300, c)
		return
	}
	user,err := u.userService.Login(request.Email, request.Password)
	if err != nil {
		u.ResFailJson(err.Error(),301, c)
		return
	}
	u.ResSuccessJson(gin.H{
		"token": "",
		"user": user,
	},c)
}

func (u UserController) PostRegister(c *gin.Context) {
	var request datamodels.UserRegisterFrom
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
		"token": "",
		"user": user,
	},c)
}

func NewUserController(userService services.IUserService) UserController {
	return UserController{
		userService: userService,
	}
}