package gee

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type Context struct{
	Writer http.ResponseWriter
	Req *http.Request
	Path string
	Method string
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer: w,
		Req: r,
		Path: r.URL.Path,
		Method: r.Method,
	}
}

func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}


func (c *Context) SetHeader(key string, val string){
	c.Writer.Header().Set(key, val)
}

func (c *Context) Status(code int){
	c.Writer.WriteHeader(code)
}

func (c *Context) SetData(code int, data []byte){
	c.Status(code)
	c.Writer.Write(data)
}

func (c *Context) String(code int, format string, values ...interface{}){
	c.SetHeader("Content-Type", "text/plain")
	c.SetData(code, ([]byte)(fmt.Sprintf(format, values)))
}

func (c *Context) HTML(code int, content string){
	c.SetHeader("Content-Type", "text/html")
	c.SetData(code, ([]byte)(content))
}

func (c *Context) JSON(code int, obj interface{}){
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

