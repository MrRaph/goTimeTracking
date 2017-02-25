package main

import (
	"log"

	toast "github.com/jacobmarshall/go-toast"
)

func DoNotify(appID, title, message, url string) {
	notification := toast.Notification{
		AppID:   appID, // Shows up in the action center (lack of accent is due to encoding issues)
		Title:   title,
		Message: message,
		Icon:    "C:\\Users\\rcharrat\\go\\src\\github.com\\MrRaph\\goTimeTracking\\icon\\clock.png",
		Actions: []toast.Action{
			{"protocol", "Open " + appID, url},
		},
	}
	err := notification.Push()
	if err != nil {
		log.Fatalln(err)
	}
}
