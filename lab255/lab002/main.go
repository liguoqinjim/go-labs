package main

import (
	"encoding/base64"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"io/ioutil"
	"log"
	"mime"
	"os"
)

func main() {
	// create new *SGMailV3
	m := mail.NewV3Mail()

	from := mail.NewEmail("liguoqinjim", "admin@liguoqinjim.cn")
	content := mail.NewContent("text/html", "<p>Sending different attachments.</p>")
	to := mail.NewEmail("Example User", "136542728@qq.com")

	m.SetFrom(from)
	m.AddContent(content)

	// create new *Personalization
	personalization := mail.NewPersonalization()
	personalization.AddTos(to)
	personalization.Subject = "Attachments - Demystified!"

	// add `personalization` to `m`
	m.AddPersonalizations(personalization)

	// read/attach .txt file
	a_txt := mail.NewAttachment()
	dat, err := ioutil.ReadFile("../data/testing.txt")
	if err != nil {
		log.Fatalf("read txt:%v", err)
	}
	encoded := base64.StdEncoding.EncodeToString([]byte(dat))
	a_txt.SetContent(encoded)
	a_txt.SetType("text/plain")
	a_txt.SetFilename(encodeRFC2047("测试.txt"))
	a_txt.SetDisposition("attachment")

	// read/attach .pdf file
	a_pdf := mail.NewAttachment()
	dat, err = ioutil.ReadFile("../data/testing.pdf")
	if err != nil {
		log.Fatalf("open pdf:%v", err)
	}
	encoded = base64.StdEncoding.EncodeToString([]byte(dat))
	a_pdf.SetContent(encoded)
	a_pdf.SetType("application/pdf")
	a_pdf.SetFilename(encodeRFC2047("第一份PDF.pdf"))
	a_pdf.SetDisposition("attachment")

	// read/attach inline .jpg file
	a_jpg := mail.NewAttachment()
	dat, err = ioutil.ReadFile("../data/testing.png")
	if err != nil {
		log.Fatalf("png read:%v", err)
	}
	encoded = base64.StdEncoding.EncodeToString([]byte(dat))
	a_jpg.SetContent(encoded)
	a_jpg.SetType("image/png")
	a_jpg.SetFilename("testing.png")
	a_jpg.SetDisposition("inline")
	a_jpg.SetContentID("Test Attachment")

	// add `a_txt`, `a_pdf` and `a_jpg` to `m`
	m.AddAttachment(a_txt)
	m.AddAttachment(a_pdf)
	m.AddAttachment(a_jpg)

	request := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	request.Body = mail.GetRequestBody(m)
	response, err := sendgrid.API(request)
	if err != nil {
		log.Fatalf("send error:%v", err)
	} else {
		log.Println(response.StatusCode)
		log.Println(response.Body)
		log.Println(response.Headers)
	}
}

func encodeRFC2047(s string) string {
	return mime.QEncoding.Encode("utf-8", s)
}
