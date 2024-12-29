package crawlerService

import (
	"strings"

	"golang.org/x/net/html"
)

type Schedule struct {
	Day    string
	Time   string
	Link   string
	Title  string
	Author string
}

func ExtractData(n *html.Node) []Schedule {
	var schedules []Schedule
	var currentDay string

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n == nil {
			return
		}

		if n.Type == html.ElementNode && n.Data == "div" {
			for _, a := range n.Attr {
				if a.Key == "class" && strings.Contains(a.Val, "day-wrappper") {
					currentDay = extractDay(n)
				}
			}
		}

		if n.Type == html.ElementNode && n.Data == "div" {
			for _, a := range n.Attr {
				if a.Key == "class" && a.Val == "card-block" {
					schedule := extractSchedule(n)
					schedule.Day = currentDay
					schedules = append(schedules, schedule)
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(n)

	var filteredSchedules []Schedule
	for _, schedule := range schedules {
		if schedule.Time != "" && schedule.Link != "" && schedule.Title != "" {
			filteredSchedules = append(filteredSchedules, schedule)
		}
	}
	return filteredSchedules
}

func extractDay(n *html.Node) string {
	if n == nil {
		return ""
	}

	var day string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n == nil {
			return
		}

		for _, a := range n.Attr {
			if a.Key == "class" && a.Val == "day" {
				for c := n.FirstChild; c != nil; c = c.NextSibling {
					if c.Type == html.ElementNode && c.Data == "span" {
						day = c.FirstChild.Data
						return
					}
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(n)
	return day
}

func extractSchedule(n *html.Node) Schedule {
	if n == nil {
		return Schedule{}
	}

	var schedule Schedule

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && c.Data == "p" && len(c.Attr) > 0 && c.Attr[0].Val == "mb-2" {
			if c.FirstChild.Data == "small" && c.FirstChild.FirstChild.Data == "b" {
				schedule.Time = c.FirstChild.FirstChild.FirstChild.Data
			}

			for sc := c.FirstChild.NextSibling; sc != nil; sc = sc.NextSibling {
				if sc.Type == html.ElementNode && sc.Data == "small" {
					schedule.Time = extractText(sc)
				}
			}

		}

		if c.Type == html.ElementNode && c.Data == "a" {
			for _, a := range c.Attr {
				if a.Key == "href" {
					schedule.Link = a.Val
				}
			}
			for ac := c.FirstChild; ac != nil; ac = ac.NextSibling {
				if ac.Type == html.ElementNode && ac.Data == "h6" {
					schedule.Title = extractText(ac)
				}
			}
		}

		if c.Type == html.ElementNode && c.Data == "p" && len(c.Attr) == 0 {
			if c.FirstChild != nil && c.FirstChild.NextSibling.Data == "small" {
				schedule.Author = c.FirstChild.NextSibling.FirstChild.Data
			}
		}
	}

	return schedule
}

func extractText(n *html.Node) string {
	if n == nil {
		return ""
	}

	var text string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n == nil {
			return
		}
		if n.Type == html.TextNode {
			text += n.Data
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(n)
	return strings.TrimSpace(text)
}
