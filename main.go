package main

import (
	"fmt"
	"net/http"
	//"os"
	"github.com/joho/godotenv"
	"github.com/mojcaostir/kinodvor/crawlerService"
	"github.com/mojcaostir/kinodvor/emailService"
	"github.com/mojcaostir/kinodvor/htmlService"
	"golang.org/x/net/html"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	resp, err := http.Get("https://www.kinodvor.org/spored/")
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println("Error parsing HTML:", err)
		return
	}

/* 	file, err := os.Open("_examples/raw.html")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    doc, err := html.Parse(file)
    if err != nil {
        fmt.Println("Error parsing HTML:", err)
        return
    } */

	schedules := crawlerService.ExtractData(doc)

    htmlContent := htmlService.GenerateHTML(schedules)

    emailService.SendEmail("jarvexostirpotrc@gmail.com", "Kinodvor Spored", htmlContent)
	//fmt.Println(htmlContent)
}

