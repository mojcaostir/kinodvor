package emailService

import (
	"fmt"
	"os"
	"strings"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendEmail(recipients, subject, body string) error {
    from := mail.NewEmail("Mojca Oštir", os.Getenv("SENDGRID_FROM"))
    client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))


    recipientList := strings.Split(recipients, ",")
    for _, recipient := range recipientList {
        emailParts := strings.Split(recipient, "@")
        if len(emailParts) != 2 {
            return fmt.Errorf("invalid email address: %s", recipient)
        }
        name := emailParts[0]
        to := mail.NewEmail(name, recipient)
        message := mail.NewSingleEmail(from, subject, to, body, body)
        response, err := client.Send(message)
        if err != nil {
            return fmt.Errorf("error sending email: %w", err)
        }
        if response.StatusCode >= 400 {
            return fmt.Errorf("failed to send email: %s", response.Body)
        }
    }
    fmt.Println("Email sent successfully")
    return nil
}
