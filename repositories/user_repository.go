package repositories

import (
	"github.com/go-xorm/xorm"
	"github.com/liu578101804/distributed-authentication/datamodels"
	"errors"
)

type IUserRepository interface {
	GetUserByEmail(string) (datamodels.User,error)
	InsertUser(*datamodels.User) (int64,error)
}

type UserRepository struct {
	conn *xorm.Engine
}

func (u *UserRepository) GetUserByEmail(email string) (datamodels.User,error) {
	user := datamodels.User{
		Email:email,
	}
	has,err := u.conn.Get(&user)
	if err != nil {
		return datamodels.User{},err
	}
	if !has {
		return datamodels.User{},errors.New("can't find user")
	}
	return user,nil
}

func (u *UserRepository) InsertUser(user *datamodels.User) (int64,error) {
	return u.conn.Insert(user)
}

func NewUserRepository(conn *xorm.Engine) IUserRepository {
	return &UserRepository{
		conn:conn,
	}
}