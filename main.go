package main

import (
	"fmt"

	"github.com/caseymrm/menuet"
)

// basestate shows a default icon for the application.
// todo: add speedtest buttons, as well as historic data.
func basestate() {
	// menuet.App().Alert(Info("Title", "Info"))
	menuet.App().SetMenuState(&menuet.MenuState{
		Image: "icon.icns",
	})
	menuet.App().MenuChanged()
}


// // Info returns a generic informative alert.
// func Info(title, info string) menuet.Alert {
// 	return menuet.Alert{
// 		MessageText:     title,
// 		InformativeText: info,
// 		Buttons:         []string{"Okay"},
// 	}
// }

// menu returns the default menu items.
// todo: proper ones.
func menu() []menuet.MenuItem {
	return []menuet.MenuItem{
		{
			Type: menuet.Regular,
			Text: "Example Menu Button",
			FontSize: 14,
			FontWeight: menuet.WeightHeavy,
			State: false,
			Clicked: menutext,
		},
	}
}

func menutext() {
	menuet.App().SetMenuState(&menuet.MenuState{
		Title: "example text",
	})
	menuet.App().MenuChanged()
}


func main() {
	go basestate()

	app := menuet.App()
	app.Name = "GoFast"
	app.Label = "com.github.com.johnaoss.gofast"
	app.Children = menu

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
