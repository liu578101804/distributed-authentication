package main

import (
	"github.com/gin-gonic/gin"
	"github.com/liu578101804/distributed-authentication/api/controllers"
)

func main() {
	router := gin.Default()

	//注册控制器
	controllers.RegisterController(router.RouterGroup)

	router.Run(":8086")
}
