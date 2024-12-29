package emailService

import (
    "fmt"
    "net/smtp"
    "os"
)

func SendEmail(to, subject, body string) {
    from := os.Getenv("EMAIL_FROM")
    password := os.Getenv("EMAIL_PASSWORD")

    // SMTP server configuration.
    smtpHost := "smtp.gmail.com"
    smtpPort := "587"

    // Message.
    message := []byte("MIME-Version: 1.0\r\n" +
        "Content-Type: text/html; charset=\"UTF-8\"\r\n" +
        "Subject: " + subject + "\r\n" +
        "\r\n" + body + "\r\n")

    // Authentication.
    auth := smtp.PlainAuth("", from, password, smtpHost)

    // Sending email.
    err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, message)
    if err != nil {
        fmt.Println("Error sending email:", err)
        return
    }
    fmt.Println("Email sent successfully")
}
