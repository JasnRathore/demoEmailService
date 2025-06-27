package main

import (
	"fmt"
	"strconv"

	"github.com/jasnrathore/goemail"
	"github.com/jasnrathore/trackingmail"
)

type Target struct {
	Mail string
	Id   int
}

func main() {
	// change to your credentials
	prof := goemail.NewProfile(
		"Jordan Phisher",
		"yourEmail.com",
		"Jordan Phisher <yourEmail.com>",
		"smtp.gmail.com:587",
		"",
	)

	// change to your targets
	mails := []Target{
		{
			Mail: "test@gmail.com",
			Id:   12,
		},
		{
			Mail: "test2@gmail.com",
			Id:   13,
		},
	}
	tracker := emailtracker.NewTracker(
		emailtracker.Config{
			Port:   8080,
			Domain: "localhost:8080",
			Path:   "/pixel",
		},
		func(evt emailtracker.OpenEvent) {
			var mail string
			for _, item := range mails {
				if strconv.Itoa(item.Id) == evt.ID {
					mail = item.Mail
				}
			}
			fmt.Printf("%s\n", mail)
			fmt.Printf("Email opened: %+v\n", evt)

		},
	)
	go tracker.Start()
	for _, item := range mails {
		trackURL := tracker.GenerateLink(strconv.Itoa(item.Id))
		body := `
    <html>
        <body>
            Hello, this is a tracked email.<br>
        </body>
    </html>
`
		err := prof.SendMailWithTracking(item.Mail, "demo for tracker and mail", body, nil, trackURL)
		if err != nil {
			fmt.Printf("%v\n", err)
		}
	}
	select {}
}
