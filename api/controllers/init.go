package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/liu578101804/distributed-authentication/repositories"
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/liu578101804/distributed-authentication/services"
	"github.com/liu578101804/distributed-authentication/datamodels"
)

func createConn() (*xorm.Engine,error) {
	return xorm.NewEngine("mysql","root:root@tcp(127.0.0.1:3306)/authentication?charset=utf8")
}

func RegisterController(group gin.RouterGroup)  {
	var(
		engine *xorm.Engine
		err error

		userRepository repositories.IUserRepository
		emailRepository repositories.IEmailRepository
		userCodeRepository repositories.IUserCodeRepository
		tokenBlackRepository repositories.ITokenBlackRepository

		userService services.IUserService
		emailService services.IEmailService
		jwtService services.IJwtService
		tokenBlackService services.ITokenBlackService

		userController *UserController
		emailController *EmailController
		tokenController *TokenController
	)

	if engine,err = createConn();err != nil{
		panic(err.Error())
	}

	//同步表结构
	engine.Sync2(new(datamodels.User),new(datamodels.UserCode),new(datamodels.Email),new(datamodels.TokenBlack))

	engine.ShowSQL(true)

	//new repository
	userRepository = repositories.NewUserRepository(engine)
	emailRepository = repositories.NewEmailRepository(engine)
	userCodeRepository = repositories.NewUserCodeRepository(engine)
	tokenBlackRepository = repositories.NewTokenBlackRepository(engine)

	//new service
	userService = services.NewUserService(userRepository, userCodeRepository)
	emailService = services.NewEmailService(emailRepository, userCodeRepository, userRepository)
	jwtService = services.NewJwtService()
	tokenBlackService = services.ITokenBlackService(tokenBlackRepository)

	//new controller
	userController = NewUserController(userService,jwtService)
	emailController = NewEmailController(emailService)
	tokenController = NewTokenController(tokenBlackService,jwtService)

	group.POST("/user/login", userController.PostLogin)
	group.POST("/user/register", userController.PostRegister)
	group.POST("/user/password/change", userController.PostPasswordChange)
	group.POST("/user/password/reset", userController.PostPasswordReset)

	group.POST("/token/destroy", tokenController.PostDestroyToken)
	group.POST("/token/check", tokenController.PostCheckToken)

	group.POST("/email/send/password/reset", emailController.PostSendEmailForRestPassword)

}

