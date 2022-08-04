package client

import (
	"strings"
	"sync"

	"github.com/valyala/fasthttp"
)

type Response struct {
	client      *Client
	request     *Request
	cookie      []*fasthttp.Cookie
	rawResponse *fasthttp.Response
}

// setClient method sets client object in response instance.
// Use core object in the client.
func (r *Response) setClient(c *Client) {
	r.client = c
}

// setRequest method sets Request object in response instance.
// The request will be released when the Response.Close is called.
func (r *Response) setRequest(req *Request) {
	r.request = req
}

// Status method returns the HTTP status string for the executed request.
func (r *Response) Status() string {
	return string(r.rawResponse.Header.StatusMessage())
}

// StatusCode method returns the HTTP status code for the executed request.
func (r *Response) StatusCode() int {
	return r.rawResponse.StatusCode()
}

// Protocol method returns the HTTP response protocol used for the request.
func (r *Response) Protocol() string {
	return string(r.rawResponse.Header.Protocol())
}

// Header method returns the response headers.
func (r *Response) Header() fasthttp.ResponseHeader {
	return r.rawResponse.Header
}

// Cookies method to access all the response cookies.
func (r *Response) Cookies() []*fasthttp.Cookie {
	return r.cookie
}

// Body method returns HTTP response as []byte array for the executed request.
func (r *Response) Body() []byte {
	return r.rawResponse.Body()
}

// String method returns the body of the server response as String.
func (r *Response) String() string {
	return strings.TrimSpace(string(r.Body()))
}

// JSON method will unmarshal body to json.
func (r *Response) JSON(v any) error {
	return r.client.core.jsonUnmarshal(r.Body(), v)
}

// XML method will unmarshal body to xml.
func (r *Response) XML(v any) error {
	return r.client.core.xmlUnmarshal(r.Body(), v)
}

// Reset clear Response object.
func (r *Response) Reset() {
	r.client = nil
	r.request = nil
	copied := r.cookie
	r.cookie = []*fasthttp.Cookie{}
	for _, v := range copied {
		fasthttp.ReleaseCookie(v)
	}

	r.rawResponse.Reset()
}

// Close method will release Request object and Response object,
// after call Close please don't use these object.
func (r *Response) Close() {
	if r.request != nil {
		tmp := r.request
		r.request = nil
		ReleaseRequest(tmp)
	}
	ReleaseResponse(r)
}

var responsePool sync.Pool

// AcquireResponse returns an empty response object from the pool.
//
// The returned response may be returned to the pool with ReleaseResponse when no longer needed.
// This allows reducing GC load.
func AcquireResponse() (resp *Response) {
	respv := responsePool.Get()
	if respv != nil {
		resp = respv.(*Response)
		return
	}
	resp = &Response{
		cookie:      []*fasthttp.Cookie{},
		rawResponse: fasthttp.AcquireResponse(),
	}

	return
}

// ReleaseResponse returns the object acquired via AcquireResponse to the pool.
//
// Do not access the released Response object, otherwise data races may occur.
func ReleaseResponse(resp *Response) {
	resp.Reset()
	responsePool.Put(resp)
}
