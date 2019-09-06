package services

import (
	"github.com/liu578101804/distributed-authentication/repositories"
	"github.com/liu578101804/distributed-authentication/utils"
	"fmt"
	"github.com/liu578101804/go-tool/mailman"
	"github.com/liu578101804/distributed-authentication/datamodels"
	"strings"
	"time"
)

type IEmailService interface {
	SendRestPasswordEmail(string,string) error
}

type EmailService struct {
	emailRepository 	repositories.IEmailRepository
	userCodeRepository 	repositories.IUserCodeRepository
	userRepository 		repositories.IUserRepository
	mailMan 			mailman.IMailman
}

func (e EmailService) SendRestPasswordEmail(baseUrl string, emailAddr string) error {

	//查询用户
	userM,err := e.userRepository.GetUserByEmail(emailAddr)
	if err != nil {
		return err
	}

	//生成随机code
	code := utils.CreatePasswordCode()
	//生成邮件内容
	body := fmt.Sprintf(`
<h1>找回密码</h1>
<a href="%v?code=%s">点击我</a>
<p>%v?code=%s</p>
`,baseUrl,code,baseUrl,code)
	toEmails := []string{userM.Email}
	subject := "找回密码"
	//发送邮件
	err = e.mailMan.SendEmail(toEmails,subject, body, mailman.MailContentTypeHtml)
	if err != nil {
		return err
	}

	//记录到历史
	newTime := time.Now()
	emailM := datamodels.Email{
		To: strings.Join(toEmails,","),
		Subject: subject,
		Body: body,
		ContentType: mailman.MailContentTypeHtml,
		EmailType: datamodels.EmailTypeResetPassword,
		CreateAt: newTime,
		UpdateAt: newTime,
	}
	_,err = e.emailRepository.InsertEmail(&emailM)
	if err != nil {
		return err
	}

	//写入到用户code数据表
	userCodeM := datamodels.UserCode{
		UserId: userM.Id,
		EmailId: emailM.Id,
		Code: code,
		CreateAt: newTime,
		UpdateAt: newTime,
	}
	_,err = e.userCodeRepository.InsertUserCode(&userCodeM)
	if err != nil {
		return err
	}

	return nil
}

func newMailman() mailman.IMailman {
	return mailman.NewMailman("smtp.qq.com","578101804@qq.com","cbwahrscbwskbcfi","liu",25)
}

func NewEmailService(emailRepository repositories.IEmailRepository, userCodeRepository repositories.IUserCodeRepository, userRepository repositories.IUserRepository) IEmailService {

	emailService := EmailService{
		emailRepository: emailRepository,
		userCodeRepository: userCodeRepository,
		userRepository: userRepository,
	}
	emailService.mailMan = newMailman()

	return &emailService
}
