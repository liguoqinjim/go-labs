package driver

import (
	"encoding/json"

	"github.com/rockisch/appigo/jsonutils"
)

type Element struct {
	Driver *Driver
	ID     string
}

func (d *Driver) FindElement(elName string, elBy string) *Element {
	reqBody := map[string]string{
		"using": elBy,
		"value": elName,
	}

	appiumReq := &appiumRequest{
		"POST",
		reqBody,
		"/wd/hub/session/" + d.sessionID + "/element",
	}

	res := doAppiumRequest(appiumReq, d.driverClient, "")

	if res.StatusCode != 200 {
		statusCodeErrorHandler(res.StatusCode, 404,
			"driver: the driver was unable to find an element on the screen using the specified arguments")
		statusCodeErrorHandler(res.StatusCode, 400,
			"driver: an invalid argument was passed to the findElement function")
	}

	mapBody := jsonutils.JSONToMap(res.Body)
	value := map[string]string{}

	err := json.Unmarshal(*mapBody["value"], &value)
	if err != nil {
		panic(err)
	}

	return &Element{d, value["ELEMENT"]}
}

func (el *Element) Click() {
	appiumReq := &appiumRequest{
		"POST",
		nil,
		"/wd/hub/session/" + el.Driver.sessionID + "/element/" + el.ID + "/click",
	}

	res := doAppiumRequest(appiumReq, el.Driver.driverClient, "")

	if res.StatusCode != 200 {
		panic("ERROR IN ELEMENT CLICK")
	}
}

func (el *Element) SendKeys(keys string) {
	reqBody := map[string]string{
		"value": keys,
	}

	appiumReq := &appiumRequest{
		"POST",
		reqBody,
		"/wd/hub/session" + el.Driver.sessionID + "/element/" + el.ID + "/sendkeys",
	}

	res := doAppiumRequest(appiumReq, el.Driver.driverClient, "")

	if res.StatusCode != 0 {
		panic("ERROR IN SEND KEYS")
	}
}
