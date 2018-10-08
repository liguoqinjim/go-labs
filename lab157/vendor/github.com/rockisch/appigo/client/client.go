package client

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Client contains info related on where and how to do the requests.
type Client struct {
	ClientURL  string
	HTTPClient *http.Client
}

// Response is used when returning requests.
type Response struct {
	StatusCode int
	Body       *[]byte
}

// CreateClient create and returns a Client struct.
func CreateClient(url string) *Client {
	client := &http.Client{}
	newClient := &Client{url, client}

	return newClient
}

// MakeRequest does a request using the data contained in the Client caller and the
// specified parameters. Returns an error/panics if anything wrong happens with the request or
// with the reading of the response.
func (c *Client) MakeRequest(method string, reqBody *[]byte, path string) (Response, error) {
	req, err := http.NewRequest(method, c.ClientURL+path, bytes.NewReader(*reqBody))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return Response{}, fmt.Errorf("client: error connecting to the url %s", c.ClientURL)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	clientResponse := Response{
		resp.StatusCode,
		&respBody,
	}

	return clientResponse, nil
}
