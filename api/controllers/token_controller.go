package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/liu578101804/distributed-authentication/services"
	"github.com/liu578101804/distributed-authentication/api/from"
	"github.com/liu578101804/distributed-authentication/datamodels"
	"time"
)

type TokenController struct {
	BaseController
	tokenBlackService services.ITokenBlackService
	jwtService services.IJwtService
}

func (t *TokenController) PostDestroyToken(c *gin.Context) {
	var request from.DestroyTokenFrom
	if err := c.ShouldBind(&request);err != nil{
		t.ResFailJson(err.Error(),300, c)
		return
	}

	//解析jwt
	data,err := t.jwtService.CheckTokenAndGetBody(request.Token)
	if err != nil {
		t.ResFailJson(err.Error(),301, c)
		return
	}

	nowTime := time.Now()
	tokenBlackM := datamodels.TokenBlack{
		Email: data["email"].(string),
		Token: request.Token,
		CreateAt: nowTime,
		UpdateAt: nowTime,
	}
	_, err = t.tokenBlackService.InsertTokenBlack(&tokenBlackM)
	if err != nil {
		t.ResFailJson(err.Error(),302, c)
		return
	}

	t.ResSuccessJson(gin.H{
		"info": tokenBlackM,
	}, c)
}

func (t *TokenController) PostCheckToken(c *gin.Context) {
	var request from.CheckTokenFrom
	if err := c.ShouldBind(&request);err != nil{
		t.ResFailJson(err.Error(),300, c)
		return
	}

	//解析jwt
	data,err := t.jwtService.CheckTokenAndGetBody(request.Token)
	if err != nil {
		t.ResFailJson(err.Error(),301, c)
		return
	}

	tokenBlack,err := t.tokenBlackService.GetTokenBlackByEmailAndToken( data["email"].(string),request.Token)
	if err != nil {
		t.ResFailJson(err.Error(),302, c)
		return
	}

	t.ResSuccessJson(gin.H{
		"info": tokenBlack,
	}, c)
}

func NewTokenController(tokenBlackService services.ITokenBlackService,jwtService services.IJwtService) *TokenController {
	return &TokenController{
		tokenBlackService: tokenBlackService,
		jwtService: jwtService,
	}
}