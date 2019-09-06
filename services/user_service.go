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
	Login(email,password string) (datamodels.User, error)
	//注册
	Register(email,password,name string) (datamodels.User, error)
	//修改密码
	ChangePassword(email,oldPassword,newPassword string) error
	//重置密码
	ReSetPassword(email,code,password string) error
}

type UserService struct {
	userRepository repositories.IUserRepository
	userCodeRepository 	repositories.IUserCodeRepository
}

func (u UserService) Login(email string, password string) (datamodels.User,error)  {
	user,err := u.userRepository.GetUserByEmail(email)
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

	_,err := u.userRepository.InsertUser(&user)
	if err != nil {
		return datamodels.User{},err
	}
	return user,nil
}


func (u UserService)ChangePassword(email,oldPassword,newPassword string) error  {

	user,err := u.userRepository.GetUserByEmail(email)
	if err != nil {
		return err
	}

	oldPassword, _ = utils.PasswordCreate(oldPassword, user.Salt)
	if user.Password != oldPassword {
		return errors.New("old password is error")
	}

	newPassword, _ = utils.PasswordCreate(newPassword, user.Salt)
	user.Password = newPassword
	user.UpdateAt = time.Now()

	af,err := u.userRepository.UpdateUserByUserId(&user)
	if err != nil {
		return err
	}
	if af == 0 {
		return errors.New("affect is equal 0")
	}

	return nil
}


func (u UserService)ReSetPassword(email, code, password string) error {

	//查询用户
	user,err := u.userRepository.GetUserByEmail(email)
	if err != nil {
		return err
	}

	//验证
	userCode,err := u.userCodeRepository.GetCodeByUserId(user.Id)
	if userCode.Code != code {
		return errors.New("code error")
	}

	//设置新密码
	newPassword, _ := utils.PasswordCreate(password, user.Salt)
	if newPassword == user.Password{
		return errors.New("new password equal old password")
	}

	//生成新密码
	user.Password, user.Salt = utils.PasswordCreate(password, "")
	user.UpdateAt = time.Now()

	//更新数据库里面的密码
	af,err := u.userRepository.UpdateUserByUserId(&user)
	if err != nil {
		return err
	}
	if af == 0 {
		return errors.New("affect is equal 0")
	}

	return nil
}

func NewUserService(userRepository repositories.IUserRepository,userCodeRepository repositories.IUserCodeRepository) IUserService {
	return UserService{
		userRepository: userRepository,
		userCodeRepository: userCodeRepository,
	}
}