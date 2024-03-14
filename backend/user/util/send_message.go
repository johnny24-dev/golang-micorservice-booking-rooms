package util

import (
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
)

const (
	//Email constant
	NameEmail         = "Booking Service"
	UsernameEmail     = "drminhvipoi1@gmail.com"
	PasswordEmail     = "kzumaozifvmpzdkf"
	SmtpAuthAddress   = "smtp.gmail.com"
	SmtpServerAddress = "smtp.gmail.com:587"
)

type GmailSender struct {
	name              string
	fromEmailAddress  string
	fromEmailPassword string
}

type EmailSender interface {
	SendEmail(subject string, content string, to []string, cc []string, bcc []string, attachFiles []string) error
}

func NewGmailSender() EmailSender {
	return &GmailSender{
		name:              NameEmail,
		fromEmailAddress:  UsernameEmail,
		fromEmailPassword: PasswordEmail,
	}
}

func (sender *GmailSender) SendEmail(subject string, otp string, to []string, cc []string, bcc []string, attachFiles []string) error {
	a := fmt.Sprintf("<!DOCTYPE html>\n<html>\n<head>\n    <meta charset=\"UTF-8\">\n    <title>OTP Verification</title>\n    <style>\n        .container {\n            width: 600px;\n            margin: 0 auto;\n            font-family: Arial, sans-serif;\n        }\n        .header {\n            background-color: #d900ff;\n            padding: 15px;\n            color: #ffffff;\n            text-align: center;\n        }\n        .content {\n            background-color: #ffffff;\n            padding: 40px;\n        }\n        .footer {\n            background-color: #f1f1f1;\n            padding: 15px;\n            text-align: center;\n            font-size: 10px;\n        }\n    </style>\n</head>\n<body>\n<div class=\"container\">\n    <div class=\"header\">\n        <h2>OTP Verification</h2>\n    </div>\n    <div class=\"content\">\n        <p>Dear User,</p>\n        <p>Your OTP for verification is:</p>\n        <h3 style=\"text-align: center; font-size: 40px; margin-top: 20px;\">%s</h3>\n        <p style=\"text-align: center; margin-top: 20px;\">Please enter OTP to complete the verification process.</p>\n    </div>\n    <div class=\"footer\">\n        <p>&copy; 2023 Your Company. All rights reserved.</p>\n    </div>\n</div>\n</body>\n</html>", otp)

	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", sender.name, sender.fromEmailAddress)
	e.Subject = subject
	e.HTML = []byte(a)
	e.To = to
	e.Cc = cc
	e.Bcc = bcc

	for _, f := range attachFiles {
		_, err := e.AttachFile(f)
		if err != nil {
			return fmt.Errorf("failed to attach file %s: %w", f, err)
		}
	}

	smtpAuth := smtp.PlainAuth("", sender.fromEmailAddress, sender.fromEmailPassword, SmtpAuthAddress)
	return e.Send(SmtpServerAddress, smtpAuth)
}
