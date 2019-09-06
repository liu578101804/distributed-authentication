package repositories

import (
	"github.com/go-xorm/xorm"
	"github.com/liu578101804/distributed-authentication/datamodels"
	"errors"
)

type IEmailRepository interface {
	InsertEmail(email *datamodels.Email)(int64,error)
	UpdateUserByEmailId(*datamodels.Email)(int64,error)
	DeleteUserByEmailId(int64)(int64,error)
	GetEmailById(int64)(datamodels.Email,error)
}

type EmailRepository struct {
	conn *xorm.Engine
}

func (e EmailRepository) InsertEmail(email *datamodels.Email)(int64,error) {
	return e.conn.Insert(email)
}

func (e EmailRepository)UpdateUserByEmailId(email *datamodels.Email)(int64,error){
	return e.conn.Update(email)
}

func (e EmailRepository)DeleteUserByEmailId(id int64)(int64,error){
	email := datamodels.Email{
		Id: id,
	}
	return e.conn.Delete(&email)
}

func (e EmailRepository)GetEmailById(id int64)(datamodels.Email,error){
	email := datamodels.Email{
		Id: id,
	}
	has,err := e.conn.Get(&email)
	if err != nil {
		return datamodels.Email{},err
	}
	if !has {
		return datamodels.Email{},errors.New("can't find email")
	}
	return email,nil
}

func NewEmailRepository(conn *xorm.Engine) IEmailRepository {
	return &EmailRepository{
		conn: conn,
	}
}
