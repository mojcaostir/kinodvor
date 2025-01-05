package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/mojcaostir/kinodvor/crawlerService"
	"github.com/mojcaostir/kinodvor/emailService"
	"github.com/mojcaostir/kinodvor/htmlService"
	"golang.org/x/net/html"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	http.HandleFunc("GET /send-schedule", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("https://www.kinodvor.org/spored/")
		if err != nil {
			http.Error(w, "Error fetching URL", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()
	
		doc, err := html.Parse(resp.Body)
		if err != nil {
			http.Error(w, "Error parsing HTML", http.StatusInternalServerError)
			return
		}
	
		schedules := crawlerService.ExtractData(doc)
		htmlContent := htmlService.GenerateHTML(schedules)
		emailService.SendEmail(os.Getenv("RECIPIENTS"), "Kinodvor Spored", htmlContent)
	
		fmt.Fprintln(w, "Email sent successfully")
	})
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Server is running on port", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}




