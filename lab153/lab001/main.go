package main

import (
	"fmt"
	"github.com/tebeka/selenium"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Start a Selenium WebDriver server instance (if one is not already
	//	// running).
	const (
		// These paths will be different on your system.
		seleniumPath    = "../vendor/selenium-server-standalone-3.14.0.jar"
		geckoDriverPath = "../vendor/geckodriver.exe"
		port            = 8080
	)
	opts := []selenium.ServiceOption{
		//selenium.StartFrameBuffer(),           // Start an X frame buffer for the browser to run in.
		selenium.GeckoDriver(geckoDriverPath), // Specify the path to GeckoDriver in order to use Firefox.
		selenium.Output(ioutil.Discard),       // Output debug information to STDERR.
	}
	selenium.SetDebug(false)

	service, err := selenium.NewSeleniumService(seleniumPath, port, opts...)
	if err != nil {
		panic(err) // panic is used only as an example and is not otherwise recommended.
	}
	defer func() {
		log.Println("stop service")
		service.Stop()
	}()

	log.Println("start")

	// Connect to the WebDriver instance running locally.
	caps := selenium.Capabilities{"browserName": "firefox"}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		panic(err)
	}
	defer func() {
		log.Println("wd quit")
		wd.Quit()
	}()

	// Navigate to the simple playground interface.
	if err := wd.Get("https://www.baidu.com/"); err != nil {
		panic(err)
	}

	// Get a reference to the text box containing code.
	elem, err := wd.FindElement(selenium.ByCSSSelector, "#kw")
	if err != nil {
		log.Fatalf("wd.FindElement error:%v", err)
	}
	// Remove the boilerplate code already in the text box.
	if err := elem.Clear(); err != nil {
		log.Fatalf("elem.Clear error:%v", err)
	}

	// Enter some new code in text box.
	if err := elem.SendKeys("hello 你好"); err != nil {
		log.Fatalf("elem.SendKeys error:%v", err)
	}

	// Click the run button.
	btn, err := wd.FindElement(selenium.ByCSSSelector, "#su")
	if err != nil {
		log.Fatalf("wd.FindElement error:%v", err)
	}
	if err := btn.Click(); err != nil {
		log.Fatalf("bt.Click error:%v", err)
	}

	<-sigs
	log.Println("end")
}
