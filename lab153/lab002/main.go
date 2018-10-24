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

	const (
		seleniumPath     = "../selenium-server-standalone-3.14.0.jar"
		chromeDriverPath = "D:/Coding/chromedriver_win32/chromedriver.exe"
		port             = 8087
	)
	opts := []selenium.ServiceOption{
		selenium.ChromeDriver(chromeDriverPath),
		selenium.Output(ioutil.Discard), // Output debug information to STDERR.
	}
	selenium.SetDebug(false)

	//service, err := selenium.NewChromeDriverService(chromeDriverPath, port, opts...)
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
	caps := selenium.Capabilities{"browserName": "chrome"}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		panic(err)
	}
	defer func() {
		log.Println("wd quit")
		if err := wd.Quit(); err != nil {
			log.Printf("wd.Quit error:%v", err)
		} else {
			log.Println("wd.Quit success")
		}
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
