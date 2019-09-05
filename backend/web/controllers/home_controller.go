package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HomeController struct {

}

func (h *HomeController) GetHome(c *gin.Context)  {

	c.HTML(http.StatusOK, "index.html", gin.H{
		"message": "hello",
		"title": "首页",
	})
}

func DefHomeController() *HomeController {
	return &HomeController{}
}