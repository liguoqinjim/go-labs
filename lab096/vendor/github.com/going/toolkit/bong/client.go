/**
 * Author:        Tony.Shao
 * Email:         xiocode@gmail.com
 * Github:        github.com/xiocode
 * File:          client.go
 * Description:   Protorpc Client
 */
package bong

import (
	"bufio"
	"github.com/golang/protobuf/proto"
	"errors"
	"fmt"
	"io"
	"net"
	"net/rpc"
	"sync"
)

type clientCodec struct {
	mutex   sync.Mutex
	r       *bufio.Reader
	w       *bufio.Writer
	c       io.Closer
	pending map[uint64]*rpc.Request
	res     *Response
	next    uint64
}

func NewClientCodec(conn io.ReadWriteCloser) rpc.ClientCodec {
	return &clientCodec{
		r:       bufio.NewReader(conn),
		w:       bufio.NewWriter(conn),
		c:       conn,
		pending: make(map[uint64]*rpc.Request),
		next:    1,
	}
}

func (this *clientCodec) WriteRequest(rpcreq *rpc.Request, param interface{}) error {
	rr := *rpcreq
	req := &Request{}

	this.mutex.Lock()
	req.Id = proto.Uint64(this.next)
	this.next++
	this.pending[*req.Id] = &rr
	this.mutex.Unlock()

	req.Method = proto.String(rpcreq.ServiceMethod)
	if msg, ok := param.(proto.Message); ok {
		body, err := proto.Marshal(msg)
		if err != nil {
			return err
		}
		req.Body = body
	} else {
		return fmt.Errorf("marshal request param error: %s", param)
	}

	f, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	if err := write(this.w, f); err != nil {
		return err
	}

	return nil
}

func (this *clientCodec) ReadResponseHeader(rpcres *rpc.Response) error {
	f, err := read(this.r)
	if err != nil {
		return err
	}

	res := &Response{}
	if err := proto.Unmarshal(f, res); err != nil {
		return err
	}

	this.mutex.Lock()
	p, ok := this.pending[res.GetId()]
	if !ok {
		this.mutex.Unlock()
		return errors.New("invalid sequence number in response")
	}
	this.res = res
	delete(this.pending, res.GetId())
	this.mutex.Unlock()

	rpcres.Seq = p.Seq
	rpcres.ServiceMethod = p.ServiceMethod
	rpcres.Error = res.GetError()

	return nil
}

func (this *clientCodec) ReadResponseBody(value interface{}) error {
	if value == nil {
		return nil
	}

	if msg, ok := value.(proto.Message); ok {
		if err := proto.Unmarshal(this.res.Body, msg); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("unmarshal response body error: %s", value)
	}

	return nil
}

func (this *clientCodec) Close() error {
	return this.c.Close()
}

func NewClient(this io.ReadWriteCloser) *rpc.Client {
	return rpc.NewClientWithCodec(NewClientCodec(this))
}

func Dial(network, address string) (*rpc.Client, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return NewClient(conn), err
}
