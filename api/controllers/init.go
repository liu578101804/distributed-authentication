package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/liu578101804/distributed-authentication/repositories"
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/liu578101804/distributed-authentication/services"
)

func createConn() (*xorm.Engine,error) {
	return xorm.NewEngine("mysql","root:root@tcp(127.0.0.1:3306)/authentication?charset=utf8")
}

func RegisterController(group gin.RouterGroup)  {
	var(
		engine *xorm.Engine
		err error

		userRepository repositories.IUserRepository
		userService services.IUserService
		userController UserController
	)

	if engine,err = createConn();err != nil{
		panic(err.Error())
	}
	engine.ShowSQL(true)

	//new repository
	userRepository = repositories.NewUserRepository(engine)
	//new service
	userService = services.NewUserService(userRepository)
	//new controller
	userController = NewUserController(userService)


	group.POST("/user/login", userController.PostLogin)
	group.POST("/user/register", userController.PostRegister)
}

