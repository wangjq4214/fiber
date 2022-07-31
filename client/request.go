package client

import (
	"context"
	"net/http"
	"sync"

	"github.com/gofiber/fiber/v3"
	"github.com/valyala/fasthttp"
)

type Request struct {
	url        string
	method     string
	ctx        context.Context
	header     *Header
	rawRequest *fasthttp.Request
}

// setMethod will set method for Request object,
// user should use request method to set method.
func (r *Request) setMethod(method string) *Request {
	r.method = method
	return r
}

// SetURL will set url for Request object.
func (r *Request) SetURL(url string) *Request {
	r.url = url
	return r
}

// Context returns the Context if its already set in request
// otherwise it creates new one using `context.Background()`.
func (r *Request) Context() context.Context {
	if r.ctx == nil {
		return context.Background()
	}
	return r.ctx
}

// SetContext sets the context.Context for current Request. It allows
// to interrupt the request execution if ctx.Done() channel is closed.
// See https://blog.golang.org/context article and the "context" package
// documentation.
func (r *Request) SetContext(ctx context.Context) *Request {
	r.ctx = ctx
	return r
}

// AddHeader method adds a single header field and its value in the request instance.
// It will override header which set in client instance.
func (r *Request) AddHeader(key, val string) *Request {
	r.header.Add(key, val)
	return r
}

// SetHeader method sets a single header field and its value in the request instance.
// It will override header which set in client instance.
func (r *Request) SetHeader(key, val string) *Request {
	r.header.Set(key, val)
	return r
}

// AddHeaders method adds multiple headers field and its values at one go in the request instance.
// It will override header which set in client instance.
func (r *Request) AddHeaders(h map[string][]string) *Request {
	r.header.AddHeaders(h)
	return r
}

// SetHeaders method sets multiple headers field and its values at one go in the request instance.
// It will override header which set in client instance.
func (r *Request) SetHeaders(h map[string]string) *Request {
	r.header.SetHeaders(h)
	return r
}

// Reset clear Request object, used by ReleaseRequest method.
func (r *Request) Reset() {
	r.url = ""
	r.method = fiber.MethodGet
	r.ctx = nil

	for k := range r.header.Header {
		delete(r.header.Header, k)
	}

	r.rawRequest.Reset()
}

type Header struct {
	http.Header
}

func (h *Header) AddHeaders(r map[string][]string) {
	for k, v := range r {
		for _, vv := range v {
			h.Header.Add(k, vv)
		}
	}
}

func (h *Header) SetHeaders(r map[string]string) {
	for k, v := range r {
		h.Header.Set(k, v)
	}
}

var requestPool sync.Pool

// AcquireRequest returns an empty request object from the pool.
//
// The returned request may be returned to the pool with ReleaseRequest when no longer needed.
// This allows reducing GC load.
func AcquireRequest() (req *Request) {
	reqv := requestPool.Get()
	if reqv != nil {
		req = reqv.(*Request)
		return
	}

	req = &Request{
		header:     &Header{Header: make(http.Header)},
		rawRequest: fasthttp.AcquireRequest(),
	}
	return
}

// ReleaseRequest returns the object acquired via AcquireRequest to the pool.
//
// Do not access the released Request object, otherwise data races may occur.
func ReleaseRequest(req *Request) {
	req.Reset()
	requestPool.Put(req)
}
