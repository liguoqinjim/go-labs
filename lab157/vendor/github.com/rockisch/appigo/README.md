# appigo
A library used to create automated tests for iOS and Android using Appium with golang

## What is working

Appigo is still in development, so a lot of features aren't implemented yet

Still, some are already working:

* Starting a session
* Closing the session
* Finding an element (with all locator strategies)
* Tapping on an element
* Sending keys to an element
* Setting a implicit wait

## Example

A simple example of an appigo implementation

```go
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

	driver := driver.CreateDriver("http://0.0.0.0:4723", caps)

	driver.Init()
	defer driver.Close()

	okButton := driver.FindElement("Ok", mobileby.ById)
	okButton.Click()
}

```
