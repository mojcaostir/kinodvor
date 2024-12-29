package htmlService

import (
	"fmt"
	"strings"

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

	var htmlBuilder strings.Builder
	
	htmlBuilder.WriteString("<html><body>")
	htmlBuilder.WriteString("<h1>Kinodvor Spored</h1>")
	htmlBuilder.WriteString("<table border='1'>")
	htmlBuilder.WriteString("<tr><th>Dan</th><th>Ura</th><th>Naslov</th><th>Režiser</th></tr>")

	for _, schedule := range daySchedules {
		var span = len(schedule.DayData)
		htmlBuilder.WriteString("<tr>")
		htmlBuilder.WriteString(fmt.Sprintf("<th rowspan='%d'>%s</th>", span, schedule.Day))
		var time string
		var title string
		var author string
		for i := 0; i < span; i++  {
			scheduleData := schedule.DayData[i]
			time = fmt.Sprintf("<td>%s</td>", scheduleData.Time)
			title = fmt.Sprintf("<td><a href='%s'>%s</a></td>",scheduleData.Link, scheduleData.Title)
			author = fmt.Sprintf("<td>%s</td>", scheduleData.Author)
			htmlBuilder.WriteString(time+title+author)
			htmlBuilder.WriteString("</tr>")
		}
		htmlBuilder.WriteString("</tr>")
	}

	htmlBuilder.WriteString("</table>")
	htmlBuilder.WriteString("</body></html>")

	return htmlBuilder.String()
}
