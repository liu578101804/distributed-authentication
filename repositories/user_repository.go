package repositories

import (
	"github.com/go-xorm/xorm"
	"github.com/liu578101804/distributed-authentication/datamodels"
	"errors"
)

type IUserRepository interface {
	InsertUser(*datamodels.User)(int64,error)
	UpdateUserByUserId(*datamodels.User)(int64,error)
	DeleteUserByUserOpenId(string)(int64,error)
	GetUserByEmail(string)(datamodels.User,error)
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

func (u *UserRepository) UpdateUserByUserId(user *datamodels.User)(int64,error) {
	return u.conn.Update(user)
}

func (u *UserRepository) DeleteUserByUserOpenId(openid string)(int64,error) {
	user := datamodels.User{
		OpenId: openid,
	}
	return u.conn.Delete(&user)
}

func NewUserRepository(conn *xorm.Engine) IUserRepository {
	return &UserRepository{
		conn:conn,
	}
}