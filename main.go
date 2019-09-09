package main

import (
	"time"
	"github.com/caseymrm/menuet"
)


// Currently just the hello world example.
func helloClock() {
	for {
		menuet.App().SetMenuState(&menuet.MenuState{
			Title: "Hello World " + time.Now().Format(":05"),
		})
		time.Sleep(time.Second)
	}
}

func main() {
	go helloClock()
	menuet.App().RunApplication()
}
