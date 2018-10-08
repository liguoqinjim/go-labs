package driver

func (d *Driver) ImplicitWait(seconds int) {
	reqBody := map[string]string{
		"ms": string(seconds * 1000),
	}

	appiumReq := appiumRequest{
		"POST",
		reqBody,
		"/wd/hub/session/" + d.sessionID + "/timeouts/implicit_wait",
	}
	resp := doAppiumRequest(&appiumReq, d.driverClient, "")

	if resp.StatusCode != 200 {
		panic("IMPLICIT WAIT ERROR")
	}
}
