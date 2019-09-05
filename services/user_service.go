package services

import (
	"github.com/liu578101804/distributed-authentication/datamodels"
	"github.com/liu578101804/distributed-authentication/repositories"
	"errors"
	"time"
	"github.com/liu578101804/distributed-authentication/utils"
)

type IUserService interface {
	//登录
	Login(email, password string) (datamodels.User, error)
	//注册
	Register(email,password,name string) (datamodels.User, error)
	//修改密码
	ChangePassword(email,oldPassword,newPassword string) error
	//重置密码
	ReSetPassword(email,password string) error
}

type UserService struct {
	repository repositories.IUserRepository
}

func (u UserService) Login(email string, password string) (datamodels.User,error)  {
	user,err := u.repository.GetUserByEmail(email)
	if err != nil {
		return datamodels.User{},err
	}
	newPassword,_ := utils.PasswordCreate(password, user.Salt)
	if user.Password != newPassword {
		return datamodels.User{},errors.New("password error")
	}
	return user,nil
}

func (u UserService)Register(email,password,name string) (datamodels.User,error)  {
	newPassword,salt := utils.PasswordCreate(password, "")
	openId := utils.CreateOpenId()
	nowTime := time.Now()

	user := datamodels.User{
		Name: name,
		OpenId: openId,
		Password: newPassword,
		Salt: salt,
		Email: email,
		CreateAt: nowTime,
		UpdateAt: nowTime,
	}
	_,err := u.repository.InsertUser(&user)
	if err != nil {
		return datamodels.User{},err
	}
	return user,nil
}

//TODO: 待完成
func (u UserService)ChangePassword(email,oldPassword,newPassword string) error  {
	return nil
}

//TODO: 待完成
func (u UserService)ReSetPassword(email, password string) error {
	return nil
}

func NewUserService(repository repositories.IUserRepository) IUserService {
	return UserService{
		repository: repository,
	}
}