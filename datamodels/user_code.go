package datamodels

import "time"

type UserCode struct {
	Id     int64  	`json:"id" xorm:"int notnull autoincr pk 'id'"`

	UserId 	int64 	`json:"user_id" xorm:"int notnull 'user_id'"`
	EmailId int64 	`json:"email_id" xorm:"int notnull 'email_id'"`
	Code   	string 	`json:"code" xorm:"varchar(25) notnull 'code'"`

	ExpirationTime 	time.Time 	`json:"expiration_time" xorm:"datetime 'expiration_time'"`
	CreateAt 		time.Time 	`json:"create_at" xorm:"created 'create_at'"`
	UpdateAt 		time.Time 	`json:"update_at" xorm:"updated 'update_at'"`
}

func (u UserCode) TableName() string {
	return "user_code"
}

