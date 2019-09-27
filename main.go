package main

import (
	"context"
	"time"
	"fmt"
	"github.com/caseymrm/menuet"
)

// mainMenu returns the default menu items.
// todo: proper ones.
func mainMenu() []menuet.MenuItem {
	return []menuet.MenuItem{
		{
			Text: "Go Fast",
			FontWeight: menuet.WeightBold,
		},
		{
			Type: menuet.Separator,
		},
		{
			Text: "Run Test",
			Clicked: checkSpeed,
		},
		{
			Type: menuet.Separator,
		},
		{
			Text: "History",
			Clicked: placeholderAction,
		},
	}
}

// placeholderAction is an action that opens an alert when clicked.
// This should be used as a placeHolder when behaviour is not yet ready.
func placeholderAction() {
	menuet.App().Alert(menuet.Alert{
		MessageText: "This button is unimplemented",
		InformativeText: "Please close this popup, and potentially open an issue on GitHub",
	})
}

// this is currently spaghetti code, will fix
func checkSpeed() {
	menu := mainMenu()
	menu[2].Clicked = nil

	children := func() []menuet.MenuItem{
		return menu
	}
	menuet.App().Children = children

	// Makes menu look pretty when updating.
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		var count int
		for {
			select {
			case <-ctx.Done():
				return
			default:
				menuet.App().Children = func() []menuet.MenuItem{
					menu[2].Text = "Measuring."
					for i := 0; i < count % 3; i++ {
						menu[2].Text += "."
					}
					count++
					return menu
				}
				menuet.App().MenuChanged()
				time.Sleep(500 * time.Millisecond)	
			}
		}
	}()

	// Check results, then reset menu state
	result := client.Measure()
	cancel()
	menuet.App().Children = mainMenu
	menuet.App().MenuChanged()
	menuet.App().Alert(menuet.Alert{
		MessageText: fmt.Sprintf("Your internet speed is %.2f Mbps\n", result.Average/1000),
		InformativeText: fmt.Sprintf("The fastest speed measured was %.2f Mbps, and the slowest recorded was %.2f Mbps", result.Max/1000, result.Min/1000),
	})
	
	// TODO: Store alert to do history for it.
}


func main() {
	app := menuet.App()
	app.Name = "GoFast"
	app.Label = "com.github.com.johnaoss.gofast"
	app.Children = mainMenu
	app.SetMenuState(&menuet.MenuState{
		Image: "icon.icns",
	})
	
	// The main running of the application. As such every other process should
	// be called by a given action on the menu bar, or somehow else.
	app.RunApplication()
}
