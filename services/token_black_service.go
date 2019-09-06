package services

import (
	"github.com/liu578101804/distributed-authentication/repositories"
	"github.com/liu578101804/distributed-authentication/datamodels"
)

type ITokenBlackService interface {
	InsertTokenBlack(token *datamodels.TokenBlack)(int64,error)
	GetTokenBlackByEmailAndToken(string,string)(datamodels.TokenBlack,error)
}

type TokenBlackService struct {
	tokenBlackRepository 		repositories.ITokenBlackRepository
}

func (t TokenBlackService) InsertTokenBlack(token *datamodels.TokenBlack)(int64,error) {
	return t.tokenBlackRepository.InsertTokenBlack(token)
}

func (t TokenBlackService) GetTokenBlackByEmailAndToken(email,token string)(datamodels.TokenBlack,error) {
	return t.tokenBlackRepository.GetTokenBlackByEmailAndToken(email, token)
}

func NewTokenBlackService(tokenBlackRepository repositories.ITokenBlackRepository) ITokenBlackService {
	return &TokenBlackService{}
}