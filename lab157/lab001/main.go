package main

import (
	"github.com/rockisch/appigo/driver"
	"github.com/rockisch/appigo/mobileby"
)

func main() {
	caps := map[string]string{
		"deviceName":      "Genymotion",
		"platformName":    "Android",
		"platformVersion": "8.0.0",
		"appPackage":      "com.android.calculator2",
		"appActivity":     ".Calculator",
	}

	//capabilities.setCapability("appPackage", "com.android.calculator2");
	//capabilities.setCapability("appActivity", ".Calculator");

	driver := driver.CreateDriver("http://0.0.0.0:4723", caps)

	driver.Init()
	defer driver.Close()

	//点击9
	btn9 := driver.FindElement("digit_9", mobileby.ById)
	btn9.Click()

	//点击+号
	btnAdd := driver.FindElement("op_add", mobileby.ById)
	btnAdd.Click()

	//点击9
	btn9.Click()

	//点击=号
	btnEqual := driver.FindElement("eq", mobileby.ById)
	btnEqual.Click()
}
