package from


type DestroyTokenFrom struct {
	Token 		string 	`form:"token" binding:"required"`
}

type CheckTokenFrom struct {
	Token 		string 	`form:"token" binding:"required"`
}