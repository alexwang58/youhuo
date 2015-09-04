package components

import (
	//"github.com/gorilla/mux"
	"net/http"
)

type Context struct {
	// host name
	host string
	// controller name
	name string
	// action name
	action string
	// request IP
	remoteAddr string
	// request URI
	uri string
	// quyer params
	param   map[string]string
	writer  http.ResponseWriter
	request *http.Request
	layout  string
	//status  int
}

// New returns an Context object
func NewContext() *Context {
	return &Context{}
}

func (c *Context) SetHost(host string) {
	c.host = host
}

func (c *Context) GetHost() string {
	return c.host
}

func (c *Context) SetName(name string) {
	c.name = name
}

func (c *Context) GetName() string {
	return c.name
}

func (c *Context) SetAction(action string) {
	c.action = action
}

func (c *Context) GetAction() string {
	return c.action
}

func (c *Context) SetRemoteAddr(raddr string) {
	c.remoteAddr = raddr
}

func (c *Context) GetRemoteAddr() string {
	return c.remoteAddr
}

func (c *Context) SetUri(uri string) {
	c.uri = uri
}

func (c *Context) GetUri() string {
	return c.uri
}

func (c *Context) SetLayout(layout string) {
	c.layout = layout
}

func (c *Context) GetLayout() string {
	return c.layout
}

/*
func (c *Context) SetStatus(status int) {
	c.status = status
}

func (c *Context) GetStatus() int {
	return c.status
}
*/
func (c *Context) SetQueryParam(param map[string]string) {
	c.param = param
}

func (c *Context) GetQueryParam() map[string]string {
	return c.param
}

func (c *Context) GetParam(_var string) string {
	return c.param[_var]
}

// Get POST parameter
func (c *Context) GetPost(_var string) string {
	val := c.GetResquest().PostFormValue(_var)
	if val != "" {
		return val
	}
	return c.GetResquest().FormValue(_var)
}

func (c *Context) SetResponseWriter(writer http.ResponseWriter) {
	c.writer = writer
}

func (c *Context) GetResponseWriter() http.ResponseWriter {
	return c.writer
}

func (c *Context) SetResquest(r *http.Request) {
	c.request = r
}

func (c *Context) GetResquest() *http.Request {
	return c.request
}
