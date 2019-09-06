package controllers

import (
	"github.com/liu578101804/distributed-authentication/services"
	"github.com/gin-gonic/gin"
	"github.com/liu578101804/distributed-authentication/api/from"
)

type EmailController struct {
	BaseController
	emailService services.IEmailService
}

//发送忘记密码邮件
func (e *EmailController) PostSendEmailForRestPassword(c *gin.Context) {

	var request from.SendEmailForRestPasswordFrom
	if err := c.ShouldBind(&request);err != nil{
		e.ResFailJson(err.Error(),300, c)
		return
	}

	err := e.emailService.SendRestPasswordEmail("http://www.baidu.com", request.Email)
	if err != nil{
		e.ResFailJson(err.Error(),302, c)
		return
	}

	e.ResSuccessJson(gin.H{
	}, c)
}

func NewEmailController(emailService services.IEmailService) *EmailController {
	return &EmailController{
		emailService: emailService,
	}
}