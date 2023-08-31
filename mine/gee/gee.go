package gee

import (
	//"fmt"
	"net/http"
)

type Engine struct{
	rounter *Rounter
}

func New() *Engine {
	return &Engine{rounter: NewRounter(),}
}

func (engine *Engine) AddRounter(methond string, pattern string, handler Handler){
	key := methond + "-" + pattern
	engine.rounter.Add(key, handler)
}

func (engine *Engine) Get(pattern string, handler Handler){
	engine.AddRounter("GET", pattern, handler)
}
func (engine *Engine) Post(pattern string, handler Handler){
	engine.AddRounter("POST", pattern, handler)
}
func (engine *Engine) Run(addr string) error{
	return http.ListenAndServe(addr, engine)
}


func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request){
	c := NewContext(w, r)
	engine.rounter.Handle(c)
}