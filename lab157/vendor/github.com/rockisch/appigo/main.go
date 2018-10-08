package main

import (
	"github.com/rockisch/appigo/driver"
	"github.com/rockisch/appigo/mobileby"
)

func main() {
	caps := map[string]string{
		"deviceName":      "iPhone 6",
		"platformName":    "iOS",
		"platformVersion": "11.4",
		"app":             "/Users/joaohaas/Documents/apps/TestApp.app",
	}

	driver := driver.CreateDriver("http://192.168.0.55:4723", caps)

	driver.Init()
	defer driver.Close()

	okButton := driver.FindElement("Ok", mobileby.ById)
	okButton.Click()
}
