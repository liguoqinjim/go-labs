package main

import (
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func main() {
	from := mail.NewEmail("liguoqinjim", "admin@liguoqinjim.cn")
	subject := "测试邮件"
	to := mail.NewEmail("Example User", "136542728@qq.com")
	plainTextContent := "测试内容"
	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Fatalf("send error:%v", err)
	} else {
		log.Println(response.StatusCode)
		log.Println(response.Body)
		log.Println(response.Headers)
	}
}
