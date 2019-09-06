package repositories

import (
	"github.com/go-xorm/xorm"
	"github.com/liu578101804/distributed-authentication/datamodels"
	"errors"
)

type ITokenBlackRepository interface {
	InsertTokenBlack(token *datamodels.TokenBlack)(int64,error)
	GetTokenBlackByEmailAndToken(string,string)(datamodels.TokenBlack,error)
}

type TokenBlackRepository struct {
	conn *xorm.Engine
}

func (t *TokenBlackRepository) InsertTokenBlack(token *datamodels.TokenBlack)(int64,error) {
	return t.conn.InsertOne(token)
}

func (t *TokenBlackRepository) GetTokenBlackByEmailAndToken(email,token string)(datamodels.TokenBlack,error) {
	tokenBlackM := datamodels.TokenBlack{
		Email: email,
		Token: token,
	}
	has,err := t.conn.Get(&tokenBlackM)
	if err != nil {
		return datamodels.TokenBlack{},err
	}
	if !has {
		return datamodels.TokenBlack{},errors.New("can't find token_black")
	}
	return tokenBlackM,nil
}

func NewTokenBlackRepository(conn *xorm.Engine) ITokenBlackRepository {
	return &TokenBlackRepository{
		conn: conn,
	}
}
