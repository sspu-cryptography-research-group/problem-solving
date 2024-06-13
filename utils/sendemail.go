package utils

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/mail"
	"net/smtp"
	"time"
)

func SendEmail(branch, email string) {
	from := mail.Address{"今天commit了吗？", "chengwenq@qq.com"}
	to := mail.Address{"不积跬步，无以至千里", email}

	smtpServer := "smtp.qq.com:587"
	emailUser := "chengwenq@qq.com"
	emailPass := "gdwgfvyyqwwtecba" // 确保这是你的SMTP独立密码

	//subject := "这是一封定时提醒邮件"
	// body := "This is a test email sent via SMTP from QQ email."
	body := `
<html>
<head>
<title>今日GitHub提交提醒</title>
</head>
<body>
<h1>🌟 今日GitHub提交提醒 🌟</h1>
<p>亲爱的开发者，<br>
<strong>提醒：</strong> 看起来今天你的GitHub仓库还没有新的提交哦！</p>
<p>每一个伟大的项目都始于一个小小的commit，不要让你的项目等太久。<br>
<strong>加油：</strong> 继续你的优秀工作，让代码的火焰继续燃烧吧！</p>
<p>🔥 一点小建议：<br>
- 尝试每天至少提交一次，保持项目的活跃度。<br>
- 分享你的进展，见证你的成长。<br>
- 不要忘记给自己一点掌声，每一点进步都值得庆祝！</p>
<p><strong>鼓励：</strong> 记得，每个成功的开发者都有过无数次的提交。你的贡献是宝贵的！</p>
<p>祝你编码愉快，<br>
<strong>你的GitHub小助手:大橘同学</strong></p>
</body>
</html>
`

	client, err := smtp.Dial(smtpServer)
	if err != nil {
		log.Printf(time.Now().String()+":Failed to connect to the SMTP server: %v\n", err)
		return
	}
	defer client.Quit()

	if err := client.StartTLS(&tls.Config{ServerName: "smtp.qq.com"}); err != nil {
		log.Printf(time.Now().String()+":Failed to start TLS: %v\n", err)
		return
	}

	if err := client.Auth(smtp.PlainAuth("", emailUser, emailPass, "smtp.qq.com")); err != nil {
		log.Printf(time.Now().String()+":Failed to authenticate: %v\n", err)
		return
	}

	fmt.Println(time.Now().String() + ":SMTP服务器登陆成功")

	if err := client.Mail(from.Address); err != nil {
		log.Printf(time.Now().String()+":Failed to set sender: %v\n", err)
		return
	}
	if err := client.Rcpt(to.Address); err != nil {
		log.Printf(time.Now().String()+":Failed to set recipient: %v\n", err)
		return
	}

	// 写入邮件，设置正确的Content-Type为text/html
	w, err := client.Data()
	if err != nil {
		log.Fatalf("Failed to initialize data connection: %v\n", err)
		return
	}
	defer w.Close()

	_, err = io.WriteString(w, "From: "+from.String()+"\r\n")
	if err != nil {
		log.Fatalf("Failed to write From header: %v\n", err)
		return
	}
	_, err = io.WriteString(w, "To: "+to.String()+"\r\n")
	if err != nil {
		log.Fatalf("Failed to write To header: %v\n", err)
		return
	}
	_, err = io.WriteString(w, "Subject: 今日GitHub提交提醒\r\n")
	if err != nil {
		log.Fatalf("Failed to write Subject header: %v\n", err)
		return
	}
	_, err = io.WriteString(w, "MIME-Version: 1.0\r\n")
	if err != nil {
		log.Fatalf("Failed to write MIME-Version header: %v\n", err)
		return
	}
	_, err = io.WriteString(w, "Content-Type: text/html; charset=UTF-8\r\n\r\n")
	if err != nil {
		log.Fatalf("Failed to write Content-Type header: %v\n", err)
		return
	}

	// 写入HTML邮件正文
	_, err = io.WriteString(w, body)
	if err != nil {
		log.Fatalf("Failed to write HTML body: %v\n", err)
		return
	}

	// 邮件内容结束后，发送一个单独的点来结束邮件
	_, err = w.Write([]byte{'.', '\r', '\n'})
	if err != nil {
		log.Fatalf("Failed to send end-of-message signal: %v\n", err)
		return
	}

	fmt.Println("Email sent successfully!")
	log.Println(time.Now().String() + ":Email sent successfully! The receiver is:" + email)
}
