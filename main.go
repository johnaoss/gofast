package main

import (
	"fmt"
	"time"

	"github.com/caseymrm/menuet"
)

// helloClock is the generic clock example from the menuet example.
// todo: display icon for title, remove alert, add speedtest buttons.
func helloClock() {
	menuet.App().Alert(Info("Title", "Info"))
	for {
		menuet.App().SetMenuState(&menuet.MenuState{
			Title: time.Now().Format(time.RFC3339),
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

	// No idea what label does.
	app.Label = "Label"

	// This stores preferences in the ~/Library/Preferences/{BUNDLEID}.plist
	// Heavily cached, so removing the actual file won't necessarily remove it.
	// Look at the Makefile for more info.
	prefs := menuet.Defaults()

	// Todo: remove for debug.
	fmt.Println("The key for test:", prefs.String("test"))

	// The main running of the application. As such every other process should
	// be called by a given action on the menu bar, or somehow else.
	app.RunApplication()
}
