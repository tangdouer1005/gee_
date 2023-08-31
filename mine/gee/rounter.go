package gee

import (
	//"fmt"
	//"net/http"
)

type Handler func ( *Context)

type Rounter struct{
	handlers map[string]Handler
}
func NewRounter() *Rounter{
	return &Rounter{
		handlers: make (map[string]Handler),
	}
}
func (r *Rounter) Add(key string, val Handler){
	r.handlers[key] = val
}

func (r *Rounter) Handle(c *Context){
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok{
		handler(c)
	}else{
		c.String(0, "404 NOT FOUND %v", c.Method + "-" + c.Path)
	}
}


