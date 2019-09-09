package main

import (
	"time"

	"github.com/caseymrm/menuet"
)

// helloClock is the generic clock example from the menuet example.
func helloClock() {
	menuet.App().Alert(Info("Title", "Info"))
	for {
		menuet.App().SetMenuState(&menuet.MenuState{
			Title: "Hello Nerd " + time.Now().Format(time.RFC3339),
		})
		time.Sleep(time.Second)
	}
}

// Info returns a generic informative alert.
func Info(title, info string) menuet.Alert {
	return menuet.Alert{
		MessageText:     title,
		InformativeText: info,
		Buttons:         []string{"Okay"},
	}
}

func main() {
	go helloClock()
	app := menuet.App()
	app.Name = "GoFast"
	app.Label = "Label"
	app.RunApplication()
}
