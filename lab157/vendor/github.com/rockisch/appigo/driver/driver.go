package driver

import (
	"encoding/json"

	"github.com/rockisch/appigo/jsonutils"

	"github.com/rockisch/appigo/client"
)

// Driver object containing some data related to your session
type Driver struct {
	driverClient       *client.Client
	driverCapabilities map[string]string
	sessionID          string
}

type appiumRequest struct {
	Method  string
	BodyMap map[string]string
	Path    string
}

// CreateDriver takes the create a Driver with the specified URL and and a map of
// capabilities, returning the driver after that.
func CreateDriver(url string, capabilities map[string]string) *Driver {
	newDriver := &Driver{
		client.CreateClient(url),
		capabilities,
		"",
	}

	return newDriver
}

func doAppiumRequest(appiumReq *appiumRequest, c *client.Client, name string) *client.Response {
	resp, err := c.MakeRequest(
		appiumReq.Method,
		jsonutils.StringMapToJSON(appiumReq.BodyMap, name),
		appiumReq.Path,
	)

	if err != nil {
		panic(err)
	}

	return &resp
}

// Init tries to start a appium session with the url and capabilities stored in the driver.
func (d *Driver) Init() {
	appiumReq := &appiumRequest{
		"POST",
		d.driverCapabilities,
		"/wd/hub/session",
	}

	resp := doAppiumRequest(appiumReq, d.driverClient, "desiredCapabilities")

	statusCodeErrorHandler(
		resp.StatusCode, 500,
		"appigo: unable to create session. please, check if the specified capabilities are corret",
	)

	mapBody := jsonutils.JSONToMap(resp.Body)

	err := json.Unmarshal(*mapBody["sessionId"], &d.sessionID)
	if err != nil {
		panic(err)
	}
}

// Close closes the session stored in the driver. It's always good practice to defer "driver.Close()"
// as soon as you cal "driver.Init()"
func (d *Driver) Close() {
	appiumReq := &appiumRequest{
		"DELETE",
		nil,
		"/wd/hub/session/" + d.sessionID,
	}

	resp := doAppiumRequest(appiumReq, d.driverClient, "")

	statusCodeErrorHandler(
		resp.StatusCode, 500,
		"appigo: unable to close session",
	)
}
