package datamodels

import "time"

type TokenBlack struct {
	Id 			int64 	`json:"id" xorm:"int notnull autoincr pk 'id'"`

	Email 		string 	`json:"email" xorm:"varchar(25) notnull 'email'"`
	Token  		string 	`json:"token" xorm:"varchar(1000) notnull 'token'"`

	CreateAt 	time.Time 	`json:"create_at" xorm:"created 'create_at'"`
	UpdateAt 	time.Time 	`json:"update_at" xorm:"updated 'update_at'"`
}

func (t TokenBlack) TableName() string {
	return "token_black"
}
