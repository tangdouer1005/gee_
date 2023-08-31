package gee

import (
	"fmt"
	"net/http"
)

type Handler func (w http.ResponseWriter, r *http.Request)

type My_struct struct{
	rounter map[string]Handler
}

func New() *My_struct {
	return &My_struct{rounter: make(map[string]Handler)}
}

func (my_struct *My_struct) AddRounter(key string, handler Handler){
	my_struct.rounter[key] = handler
}

func (my_struct *My_struct) GET(path string, handler Handler){
	my_struct.AddRounter("GET-" + path, handler)
}
func (my_struct *My_struct) POST(path string, handler Handler){
	my_struct.AddRounter("POST-" + path, handler)
}
func (my_struct *My_struct) RUN(addr string) error{
	return http.ListenAndServe(addr, my_struct)
}


func (my_struct *My_struct) ServeHTTP(w http.ResponseWriter, r *http.Request){
	for k, v := range my_struct.rounter{
		fmt.Fprintf(w, "key: %v value: %v\n", k, v)
	}
	
	key := r.Method + "-" + r.URL.Path
	fmt.Fprintf(w, "key: %v \n", key)
	if handler, ok := my_struct.rounter[key]; ok{
		handler(w, r)
	}else{
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", r.URL.Path)
	}
}