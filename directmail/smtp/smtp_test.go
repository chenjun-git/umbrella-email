package smtp

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendToMail(t *testing.T) {
	assert := assert.New(t)

	user := "chenjun@umbrellacs.top"
	password := "ChenJun1990"
	host := "smtpdm.aliyun.com:25"
	to := "876765378@qq.com"
	subject := "test Golang to sendmail"
	mailtype := "html"
	replyToAddress := "swustcj@foxmail.com"
	body := `
        <html>
        <body>
        <h3>
        "保护伞：欢迎注册保护伞平台，验证码为：123456"
        </h3>
        </body>
        </html>
        `

	fmt.Println("send email")
	err := SendToMail(user, password, host, to, subject, body, mailtype, replyToAddress)
	assert.Nil(err)
	if err != nil {
		fmt.Println("Send mail error!")
		fmt.Println(err)
	} else {
		fmt.Println("Send mail success!")
	}
}
