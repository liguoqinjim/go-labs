/*
Package httpclient wraps Go's build in standard library HTTP client providing
and API to:

	* set connection timeouts
	* set a per-request transport timeout

Internally, it uses a priority queue maintained in a single go-routine
(per client instance), leveraging the Go 1.1 `CancelRequest()` API.

Basic example:

	connectTimeout := 1 * time.Second
	httpClient := httpclient.New(connectTimeout)

	reqTimeout := 1 * time.Second
	req, _ := http.NewRequest("GET", "http://127.0.0.1/test", nil)
	resp, err := httpClient.Do(req, reqTimeout)
	if err != nil {
		// handle errors
	}

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle errors
	}
	resp.Body.Close()

Setting Transport Configuration (like TLS, etc.):

	connectTimeout := 1 * time.Second
	httpClient := httpclient.New(connectTimeout)

	// disable SSL cert checks
	httpClient.Transport.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: true,
	}

	// set timeout for receiving response headers
	httpClient.Transport.ResponseHeaderTimeout = 2 * time.Second
*/
package httpclientv2

import (
	"container/heap"
	"github.com/xiocode/toolkit/httpclientv2/pqueue"
	"io"
	"net"
	"net/http"
	"sync"
	"time"
)

// returns the current version of the package
func Version() string {
	return "0.4.0"
}

type httpRequest struct {
	req *http.Request
}

// HttpClient is a type that wraps Go's built in standard library HTTP client.
//
// To change properties of the underlying transport, set them directly on the
// Transport property of the HttpClient struct, ie:
//
// 	connectTimeout := 1 * time.Second
// 	httpClient := httpclient.New(connectTimeout)
// 	// disable SSL cert checks
// 	httpClient.Transport.TLSClientConfig = &tls.Config{
// 		InsecureSkipVerify: true,
// 	}
// 	// set timeout for receiving response headers
// 	httpClient.Transport.ResponseHeaderTimeout = 2 * time.Second

type HttpClient struct {
	sync.Mutex
	Client    *http.Client
	Transport *http.Transport
	Verbose   bool
	requests  pqueue.PriorityQueue
}

// New returns a new instance of HttpClient with the specified connect timeout
func New(connectTimeout time.Duration) *HttpClient {
	client := &http.Client{}
	h := &HttpClient{
		Client:   client,
		requests: pqueue.New(16),
	}
	h.Transport = &http.Transport{
		Dial: func(network, addr string) (net.Conn, error) {
			return net.DialTimeout(network, addr, connectTimeout)
		},
	}
	client.Transport = h.Transport
	go h.transportTimeoutWorker()
	return h
}

// Do executes an HTTP request with the given transport timeout.
//
// The request must fully complete within the specified transport
// timeout or it will be cancelled.
//
// Transport timeout must never be smaller than the HttpClient's configured
// connect timeout.
func (h *HttpClient) Do(req *http.Request, transportTimeout time.Duration) (*http.Response, error) {
	absTs := time.Now().Add(transportTimeout).UnixNano()
	item := &pqueue.Item{Value: req, Priority: absTs}
	h.Lock()
	heap.Push(&h.requests, item)
	h.Unlock()
	resp, err := h.Client.Do(req)
	if err != nil {
		return nil, err
	}
	resp.Body = &bodyCloseInterceptor{rc: resp.Body, item: item, h: h}
	return resp, nil
}

func (hc *HttpClient) transportTimeoutWorker() {
	defaultSleep := 25 * time.Millisecond
	for {
		time.Sleep(defaultSleep)
		now := time.Now().UnixNano()
		for {
			hc.Lock()
			item, _ := hc.requests.PeekAndShift(now)
			hc.Unlock()

			if item == nil {
				break
			}

			req := item.Value.(*http.Request)
			hc.Transport.CancelRequest(req)
		}
	}
}

type bodyCloseInterceptor struct {
	rc   io.ReadCloser
	item *pqueue.Item
	h    *HttpClient
}

func (bci *bodyCloseInterceptor) Read(p []byte) (n int, err error) {
	return bci.rc.Read(p)
}

func (bci *bodyCloseInterceptor) Close() error {
	err := bci.rc.Close()
	bci.h.Lock()
	if bci.item.Index != -1 {
		heap.Remove(&bci.h.requests, bci.item.Index)
	}
	bci.h.Unlock()
	return err
}
