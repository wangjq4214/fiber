package client

import (
	"sync"
	"time"

	"github.com/gofiber/fiber/v3/utils"
	"github.com/valyala/fasthttp"
)

// The Client is used to create a Fiber Client with
// client-level settings that apply to all requests
// raise from the client.
//
// Fiber Client also provides an option to override
// or merge most of the client settings at the request.
type Client struct {
	core *core

	baseUrl   string
	userAgent string
	referer   string
	header    *Header
	params    *QueryParam
	cookies   *Cookie
	path      *PathParam

	timeout time.Duration
}

// R raise a request from the client.
func (c *Client) R() *Request {
	return AcquireRequest().SetClient(c)
}

// Request returns user-defined request hooks.
func (c *Client) RequestHook() []RequestHook {
	return c.core.userRequestHooks
}

// Add user-defined request hooks.
func (c *Client) AddRequestHook(h ...RequestHook) *Client {
	c.core.userRequestHooks = append(c.core.userRequestHooks, h...)
	return c
}

// ResponseHook return user-define reponse hooks.
func (c *Client) ResponseHook() []ResponseHook {
	return c.core.userResponseHooks
}

// Add user-defined response hooks.
func (c *Client) AddResponseHook(h ...ResponseHook) *Client {
	c.core.userResponseHooks = append(c.core.userResponseHooks, h...)
	return c
}

// Set HostClient dial, this method for unit test,
// maybe don't use it.
func (c *Client) SetDial(f fasthttp.DialFunc) *Client {
	c.core.client.Dial = f
	return c
}

// JSONMarshal returns json marshal function in Core.
func (c *Client) JSONMarshal() utils.JSONMarshal {
	return c.core.jsonMarshal
}

// Set json encoder.
func (c *Client) SetJSONMarshal(f utils.JSONMarshal) *Client {
	c.core.jsonMarshal = f
	return c
}

// JSONUnmarshal returns json unmarshal function in Core.
func (c *Client) JSONUnmarshal() utils.JSONUnmarshal {
	return c.core.jsonUnmarshal
}

// Set json decoder.
func (c *Client) SetJSONUnmarshal(f utils.JSONUnmarshal) *Client {
	c.core.jsonUnmarshal = f
	return c
}

// XMLMarshal returns xml marshal function in Core.
func (c *Client) XMLMarshal() utils.XMLMarshal {
	return c.core.xmlMarshal
}

// Set xml encoder.
func (c *Client) SetXMLMarshal(f utils.XMLMarshal) *Client {
	c.core.xmlMarshal = f
	return c
}

// XMLUnmarshal returns xml unmarshal function in Core.
func (c *Client) XMLUnmarshal() utils.XMLUnmarshal {
	return c.core.xmlUnmarshal
}

// Set xml decoder.
func (c *Client) SetXMLUnmarshal(f utils.XMLUnmarshal) *Client {
	c.core.xmlUnmarshal = f
	return c
}

// BaseURL returns baseurl in Client instance.
func (c *Client) BaseURL() string {
	return c.baseUrl
}

// Set baseUrl which is prefix of real url.
func (c *Client) SetBaseURL(url string) *Client {
	c.baseUrl = url
	return c
}

// AddHeader method adds a single header field and its value in the client instance.
// These headers will be applied to all requests raised from this client instance.
// Also it can be overridden at request level header options.
func (c *Client) AddHeader(key, val string) *Client {
	c.header.Add(key, val)
	return c
}

// SetHeader method sets a single header field and its value in the client instance.
// These headers will be applied to all requests raised from this client instance.
// Also it can be overridden at request level header options.
func (c *Client) SetHeader(key, val string) *Client {
	c.header.Set(key, val)
	return c
}

// AddHeaders method adds multiple headers field and its values at one go in the client instance.
// These headers will be applied to all requests raised from this client instance. Also it can be
// overridden at request level headers options.
func (c *Client) AddHeaders(h map[string][]string) *Client {
	c.header.AddHeaders(h)
	return c
}

// SetHeaders method sets multiple headers field and its values at one go in the client instance.
// These headers will be applied to all requests raised from this client instance. Also it can be
// overridden at request level headers options.
func (c *Client) SetHeaders(h map[string]string) *Client {
	c.header.SetHeaders(h)
	return c
}

// AddParam method adds a single query param field and its value in the client instance.
// These params will be applied to all requests raised from this client instance.
// Also it can be overridden at request level param options.
func (c *Client) AddParam(key, val string) *Client {
	c.params.Add(key, val)
	return c
}

// SetParam method sets a single query param field and its value in the client instance.
// These params will be applied to all requests raised from this client instance.
// Also it can be overridden at request level param options.
func (c *Client) SetParam(key, val string) *Client {
	c.params.Set(key, val)
	return c
}

// AddParams method adds multiple query params field and its values at one go in the client instance.
// These params will be applied to all requests raised from this client instance. Also it can be
// overridden at request level params options.
func (c *Client) AddParams(m map[string][]string) *Client {
	c.params.AddParams(m)
	return c
}

// SetParams method sets multiple params field and its values at one go in the client instance.
// These params will be applied to all requests raised from this client instance. Also it can be
// overridden at request level params options.
func (c *Client) SetParams(m map[string]string) *Client {
	c.params.SetParams(m)
	return c
}

// SetParamsWithStruct method sets multiple params field and its values at one go in the client instance.
// These params will be applied to all requests raised from this client instance. Also it can be
// overridden at request level params options.
func (c *Client) SetParamsWithStruct(v any) *Client {
	c.params.SetParamsWithStruct(v)
	return c
}

// DelParams method deletes single or multiple params field and its valus in client.
func (c *Client) DelParams(key ...string) *Client {
	for _, v := range key {
		c.params.Del(v)
	}
	return c
}

// SetUserAgent method sets userAgent field and its value in the client instance.
// This ua will be applied to all requests raised from this client instance.
// Also it can be overridden at request level ua options.
func (c *Client) SetUserAgent(ua string) *Client {
	c.userAgent = ua
	return c
}

// SetReferer method sets referer field and its value in the client instance.
// This referer will be applied to all requests raised from this client instance.
// Also it can be overridden at request level referer options.
func (c *Client) SetReferer(r string) *Client {
	c.referer = r
	return c
}

// SetPathParam method sets a single path param field and its value in the client instance.
// These path params will be applied to all requests raised from this client instance.
// Also it can be overridden at request level path params options.
func (c *Client) SetPathParam(key, val string) *Client {
	c.path.SetParam(key, val)
	return c
}

// SetPathParams method sets multiple path params field and its values at one go in the client instance.
// These path params will be applied to all requests raised from this client instance. Also it can be
// overridden at request level path params options.
func (c *Client) SetPathParams(m map[string]string) *Client {
	c.path.SetParams(m)
	return c
}

// SetPathParamsWithStruct method sets multiple path params field and its values at one go in the client instance.
// These path params will be applied to all requests raised from this client instance. Also it can be
// overridden at request level path params options.
func (c *Client) SetPathParamsWithStruct(v any) *Client {
	c.path.SetParamsWithStruct(v)
	return c
}

// DelPathParams method deletes single or multiple path params field and its valus in client.
func (c *Client) DelPathParams(key ...string) *Client {
	c.path.DelParams(key...)
	return c
}

// SetCookie method sets a single cookie field and its value in the client instance.
// These cookies will be applied to all requests raised from this client instance.
// Also it can be overridden at request level cookie options.
func (c *Client) SetCookie(key, val string) *Client {
	c.cookies.SetCookie(key, val)
	return c
}

// SetCookies method sets multiple cookies field and its values at one go in the client instance.
// These cookies will be applied to all requests raised from this client instance. Also it can be
// overridden at request level cookie options.
func (c *Client) SetCookies(m map[string]string) *Client {
	c.cookies.SetCookies(m)
	return c
}

// SetCookiesWithStruct method sets multiple cookies field and its values at one go in the client instance.
// These cookies will be applied to all requests raised from this client instance. Also it can be
// overridden at request level cookies options.
func (c *Client) SetCookiesWithStruct(v any) *Client {
	c.cookies.SetCookiesWithStruct(v)
	return c
}

// DelCookies method deletes single or multiple cookies field and its valus in client.
func (c *Client) DelCookies(key ...string) *Client {
	c.cookies.DelCookies(key...)
	return c
}

// SetTimeout method sets timeout val in client instance.
// This value will be applied to all requests raised from this client instance.
// Also it can be overridden at request level timeout options.
func (c *Client) SetTimeout(t time.Duration) *Client {
	c.timeout = t
	return c
}

// Get provide a API like axios which send get request.
func (c *Client) Get(url string, setter ...SetRequestOptionFunc) (*Response, error) {
	req := AcquireRequest().SetClient(c)

	for _, v := range setter {
		v(req)
	}

	return req.Get(url)
}

// Post provide a API like axios which send post request.
func (c *Client) Post(url string, setter ...SetRequestOptionFunc) (*Response, error) {
	req := AcquireRequest().SetClient(c)

	for _, v := range setter {
		v(req)
	}

	return req.Post(url)
}

// Head provide a API like axios which send head request.
func (c *Client) Head(url string, setter ...SetRequestOptionFunc) (*Response, error) {
	req := AcquireRequest().SetClient(c)

	for _, v := range setter {
		v(req)
	}

	return req.Head(url)
}

// Put provide a API like axios which send put request.
func (c *Client) Put(url string, setter ...SetRequestOptionFunc) (*Response, error) {
	req := AcquireRequest().SetClient(c)

	for _, v := range setter {
		v(req)
	}

	return req.Put(url)
}

// Delete provide a API like axios which send delete request.
func (c *Client) Delete(url string, setter ...SetRequestOptionFunc) (*Response, error) {
	req := AcquireRequest().SetClient(c)

	for _, v := range setter {
		v(req)
	}

	return req.Delete(url)
}

// Options provide a API like axios which send options request.
func (c *Client) Options(url string, setter ...SetRequestOptionFunc) (*Response, error) {
	req := AcquireRequest().SetClient(c)

	for _, v := range setter {
		v(req)
	}

	return req.Options(url)
}

// Patch provide a API like axios which send patch request.
func (c *Client) Patch(url string, setter ...SetRequestOptionFunc) (*Response, error) {
	req := AcquireRequest().SetClient(c)

	for _, v := range setter {
		v(req)
	}

	return req.Patch(url)
}

// Reset clear Client object
func (c *Client) Reset() {
	c.baseUrl = ""
	c.timeout = 0
	c.userAgent = ""
	c.referer = ""

	c.path.Reset()
	c.cookies.Reset()
	c.header.Reset()
	c.params.Reset()
}

type SetRequestOptionFunc func(r *Request)

func SetRequestHeaders(m map[string]string) SetRequestOptionFunc {
	return func(r *Request) {
		r.SetHeaders(m)
	}
}

func SetRequestQueryParams(m map[string]string) SetRequestOptionFunc {
	return func(r *Request) {
		r.SetParams(m)
	}
}

func SetRequestUserAgent(ua string) SetRequestOptionFunc {
	return func(r *Request) {
		r.SetUserAgent(ua)
	}
}

func SetRequestReferer(referer string) SetRequestOptionFunc {
	return func(r *Request) {
		r.SetReferer(referer)
	}
}

func SetRequestData(v any) SetRequestOptionFunc {
	return func(r *Request) {
		r.SetJSON(v)
	}
}

func SetRequestFormDatas(m map[string]string) SetRequestOptionFunc {
	return func(r *Request) {
		r.SetFormDatas(m)
	}
}

func SetRequestPathParams(m map[string]string) SetRequestOptionFunc {
	return func(r *Request) {
		r.SetPathParams(m)
	}
}

func SetRequestFiles(files ...*File) SetRequestOptionFunc {
	return func(r *Request) {
		r.AddFiles(files...)
	}
}

var (
	defaultClient    *Client
	defaultUserAgent = "fiber"
	clientPool       = &sync.Pool{
		New: func() any {
			return &Client{
				core: newCore(),
				header: &Header{
					RequestHeader: &fasthttp.RequestHeader{},
				},
				params: &QueryParam{
					Args: fasthttp.AcquireArgs(),
				},
				cookies: &Cookie{},
				path:    &PathParam{},
			}
		},
	}
)

func init() {
	defaultClient = AcquireClient()
}

// AcquireClient returns an empty Client object from the pool.
//
// The returned Client object may be returned to the pool with ReleaseClient when no longer needed.
// This allows reducing GC load.
func AcquireClient() *Client {
	return clientPool.Get().(*Client)
}

// ReleaseClient returns the object acquired via AcquireClient to the pool.
//
// Do not access the released Client object, otherwise data races may occur.
func ReleaseClient(c *Client) {
	c.Reset()
	clientPool.Put(c)
}

// Get default client.
func C() *Client {
	return defaultClient
}

// Replce the defaultClient, the returned function can undo.
func Replace(c *Client) func() {
	oldClient := defaultClient
	defaultClient = c

	return func() {
		defaultClient = oldClient
	}
}

// Get send a get request use defaultClient, a convenient method.
func Get(url string, setter ...SetRequestOptionFunc) (*Response, error) {
	return defaultClient.Get(url, setter...)
}

// Post send a post request use defaultClient, a convenient method.
func Post(url string, setter ...SetRequestOptionFunc) (*Response, error) {
	return defaultClient.Post(url, setter...)
}

// Head send a head request use defaultClient, a convenient method.
func Head(url string, setter ...SetRequestOptionFunc) (*Response, error) {
	return defaultClient.Head(url, setter...)
}

// Put send a put request use defaultClient, a convenient method.
func Put(url string, setter ...SetRequestOptionFunc) (*Response, error) {
	return defaultClient.Put(url, setter...)
}

// Delete send a delete request use defaultClient, a convenient method.
func Delete(url string, setter ...SetRequestOptionFunc) (*Response, error) {
	return defaultClient.Delete(url, setter...)
}

// Options send a options request use defaultClient, a convenient method.
func Options(url string, setter ...SetRequestOptionFunc) (*Response, error) {
	return defaultClient.Options(url, setter...)
}

// Patch send a patch request use defaultClient, a convenient method.
func Patch(url string, setter ...SetRequestOptionFunc) (*Response, error) {
	return defaultClient.Patch(url, setter...)
}