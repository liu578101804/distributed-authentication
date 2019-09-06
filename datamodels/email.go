package datamodels

import "time"

type EmailType string

const (
	EmailTypeResetPassword = "reset_password"
)

type Email struct {
	Id 			int64 	`json:"id" xorm:"int notnull autoincr pk 'id'"`

	To  		string 	`json:"to" xorm:"varchar(250) notnull 'to'"`
	Subject 	string	`json:"subject" xorm:"varchar(250) notnull 'subject'"`
	Body 		string	`json:"body" xorm:"varchar(2500) notnull 'body'"`
	ContentType string	`json:"content_type" xorm:"varchar(100) notnull 'content_type'"`
	EmailType 	string 	`json:"email_type" xorm:"varchar(50) notnull 'email_type'"`

	CreateAt 	time.Time 	`json:"create_at" xorm:"created 'create_at'"`
	UpdateAt 	time.Time 	`json:"update_at" xorm:"updated 'update_at'"`
}

func (e Email) TableName() string {
	return "email"
}
