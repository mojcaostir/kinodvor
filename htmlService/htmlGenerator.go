package htmlService

import (
	"github.com/mojcaostir/kinodvor/crawlerService"
)

func GenerateHTML(schedules []crawlerService.Schedule) string {
	type Data struct { 
		Time   string
		Link   string
		Title  string
		Author string
	}

	type DaySchedule struct {
		Day     string
		DayData []Data
	}

	var daySchedules []DaySchedule
	for _, schedule := range schedules {
		found := false
		for i, day := range daySchedules {
			if day.Day == schedule.Day {
				daySchedules[i].DayData = append(day.DayData, Data{
					Time:   schedule.Time,
					Link:   schedule.Link,
					Title:  schedule.Title,
					Author: schedule.Author,
				});
                found = true
                break
			}  
		}
		if !found {
            daySchedules = append(daySchedules, DaySchedule{
                Day: schedule.Day,
                DayData: []Data{
                    {
                        Time:   schedule.Time,
                        Link:   schedule.Link,
                        Title:  schedule.Title,
                        Author: schedule.Author,
                    },
                },
            })
        }
	}

    html := "<html><body>"
    for _, daySchedule := range daySchedules {
        html += "<h1>" + daySchedule.Day + "</h1>"
        for _, data := range daySchedule.DayData {
            html += "<p>" + data.Time + " - <a href='" + data.Link + "'>" + data.Title + "</a> (režija: " + data.Author + ")</p>"
        }
    }
    html += "</body></html>"

	return html
}
