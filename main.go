package main

import (
	"fmt"
	"net/smtp"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Body struct {
	Destination string
	Code        string
}

func mailMailBody(code string) string {
	temp := fmt.Sprintf("Your verification code is %s", code)
	return temp
}

func sendMail(dest string, code string) error {

	auth := smtp.PlainAuth(
		"",
		os.Getenv("ADMIN_EMAIL_ADDRESS"),
		os.Getenv("ADMIN_EMAIL_PASSWORD"),
		"smtp.gmail.com",
	)

	// msg := fmt.Sprintf("Subject: Email Verification Code \n\n %s", code)

	message := []byte(fmt.Sprintf("To: %s\r\n", dest) +
		"Subject: Thanks for signing up to Student Lifeline\r\n" +
		"\r\n" +
		mailMailBody(code))

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		os.Getenv("ADMIN_EMAIL_ADDRESS"),
		[]string{dest},
		message,
	)

	return err

}

func handleRequest(c *gin.Context) {
	var body Body
	c.Bind(&body)
	err := sendMail(body.Destination, body.Code)

	if err != nil {
		c.JSON(200, gin.H{
			"success": true,
			"message": "Email sent successfully",
		})
		return
	}
	c.JSON(200, gin.H{
		"success": false,
		"message": string(err.Error()),
	})

}

func main() {
	godotenv.Load()
	router := gin.Default()
	router.POST("/", handleRequest)
	router.Run()
}
