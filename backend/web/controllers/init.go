package controllers

import "github.com/gin-gonic/gin"


func RegisterController(group gin.RouterGroup)  {

	group.GET("/", DefHomeController().GetHome)
}
