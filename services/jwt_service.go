package services

import (
	"github.com/liu578101804/go-tool/jwt"
	"errors"
	"strings"
	"encoding/base64"
	"encoding/json"
)

type IJwtService interface {
	CreateToken(date map[string]interface{}) (string,error)
	CheckTokenAndGetBody(string)(map[string]interface{},error)
}

type JwtService struct {
	jwtTool 	jwt.IJwt
}

func (j JwtService) CreateToken(date map[string]interface{}) (string,error) {
	j.jwtTool.WriteBody(date)
	return j.jwtTool.CreateJwtString()
}

func (j JwtService) CheckTokenAndGetBody(token string)(data map[string]interface{},err error) {
	if !j.jwtTool.CheckJwtString(token) {
		return make(map[string]interface{}),errors.New("token error")
	}
	//解密
	arr := strings.Split(token,".")
	dataByte ,err := base64.URLEncoding.DecodeString(arr[1])
	if err != nil {
		return make(map[string]interface{}),err
	}
	//解析json
	err = json.Unmarshal(dataByte,&data)
	if err != nil {
		return make(map[string]interface{}),err
	}
	return
}

func NewJwtService() IJwtService {

	jwtService := JwtService{}

	jwtService.jwtTool = jwt.NewRS256()
	jwtService.jwtTool.SetHeader("RS256")
	jwtService.jwtTool.SetEncodeFunc(func(bytes []byte) string {
		return base64.URLEncoding.EncodeToString(bytes)
	})

	return &jwtService
}