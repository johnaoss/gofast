package main

import (
	"time"

	"github.com/caseymrm/menuet"
)

// Currently just the hello world example.
func helloClock() {
	for {
		menuet.App().SetMenuState(&menuet.MenuState{
			Title: "Hello Nerd " + time.Now().Format(time.RFC3339),
		})
		time.Sleep(time.Second)
	}
}

func main() {
	go helloClock()
	app := menuet.App()
	app.Name = "GoFast"
	app.RunApplication()
}
