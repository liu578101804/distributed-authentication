package repositories

import (
	"github.com/go-xorm/xorm"
	"github.com/liu578101804/distributed-authentication/datamodels"
	"errors"
	"time"
)

type IUserCodeRepository interface {
	InsertUserCode(email *datamodels.UserCode)(int64,error)
	GetCodeByUserId(userId int64)(datamodels.UserCode,error)
}

type UserCodeRepository struct {
	conn *xorm.Engine
}

func (u UserCodeRepository)InsertUserCode(email *datamodels.UserCode)(int64,error){
	return u.conn.InsertOne(email)
}

func (u UserCodeRepository)GetCodeByUserId(userId int64)(datamodels.UserCode,error){

	userCode := datamodels.UserCode{
		UserId: userId,
	}
	has,err := u.conn.Desc("create_at").Limit(1,0).Get(&userCode)
	if err != nil {
		return datamodels.UserCode{},err
	}
	if !has {
		return datamodels.UserCode{},errors.New("can't find user_code")
	}

	//判断是否过期
	if time.Now().Sub(userCode.ExpirationTime).Seconds() > 0 {
		return datamodels.UserCode{},errors.New("code expiration")
	}

	return userCode,nil
}

func NewUserCodeRepository(conn *xorm.Engine) IUserCodeRepository {
	return &UserCodeRepository{
		conn: conn,
	}
}
