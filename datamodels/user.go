package datamodels

import "time"

type User struct {
	Id 			int64 	`json:"id" xorm:"int notnull autoincr pk 'id'"`
	Name  		string 	`json:"name" xorm:"varchar(25) notnull 'name'"`
	OpenId 		string 	`json:"open_id" xorm:"varchar(25) notnull 'open_id'"`
	Password 	string 	`json:"-" xorm:"varchar(64) notnull 'password'"`
	Salt 		string 	`json:"salt" xorm:"varchar(32) notnull 'salt'"`
	Email 		string 	`json:"email" xorm:"varchar(25) unique notnull 'email'"`
	CreateAt 	time.Time 	`json:"create_at" xorm:"created 'create_at'"`
	UpdateAt 	time.Time 	`json:"update_at" xorm:"updated 'update_at'"`
}

func (u User) TableName() string {
	return "user"
}


//登录表单
type UserLoginFrom struct {
	Email 		string 	`form:"email" binding:"required"`
	Password 	string	`form:"password" binding:"required"`
}

//注册表单
type UserRegisterFrom struct {
	Email 		string 	`form:"email" binding:"required"`
	Password 	string	`form:"password" binding:"required"`
	Name 		string	`form:"name" binding:"required"`
}