# Kinodvor Schedule

This project fetches the schedule from the Kinodvor website, generates an HTML representation of the schedule, and sends it via email to specified recipients.

## Development

### Prerequisites

- Go 1.23.3
- Docker (optional, for containerization)


### Environment Variables

Create a `.env` file in the root directory with the following variables:

- SENDGRID_FROM=your_email@example.com
- SENDGRID_API_KEY=your_sendgrid_api_key
- RECIPIENTS=recipient1@example.com,recipient2@example.com
- PORT=8080


### Running the Application

To start the application, run:

```console
go run cmd/main.go
```

### Sending the Schedule

Make a GET request to the endpoint:
```console
curl -X GET http://localhost:8080/send-schedule
```
This request will send the Kinodvor schedule to the emails listed in the RECIPIENTS variable.


