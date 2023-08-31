package gee

import (
	"fmt"
	"net/http"
	"strings"
)

type Handler func ( *Context)

type Rounter struct{
	root map[string] *Node
	handlers map[string]Handler
}
func NewRounter() *Rounter{
	return &Rounter{
		root: map[string]*Node{
			"GET": &Node{pattern: "/"},
			"POST": &Node{pattern: "/"},
		},
		handlers: make (map[string]Handler),
	}
}

func TransPattern(pattern string) []string{
	vs := strings.Split(pattern, "/")
	//fmt.Printf("%v %v\n", pattern, len(vs))
	// if(len(vs) == 0){
	// 	return vs
	// }
	
	// for _, item := range vs{
	// 	fmt.Printf("%v\n", item)
	// }

	parts := make([]string, 0)

	for _, item := range vs{
		if item != "" && item != "\n" && item != " " && item != "/"{
			parts = append(parts, item)
			if len(item) != 0 && item[0] == '*'{
				break
			}
		}
	}
	
	// for _, item := range parts{
	// 	fmt.Printf("%v\n", item)
	// }
	return parts
}

func (r *Rounter) Add(methond string, pattern string, handler Handler){
	key := methond + "-" + pattern
	r.handlers[key] = handler

	n, is := r.root[methond]
	if !is{
		fmt.Printf("no methond %v", methond)
		return
	}

	parts := TransPattern(pattern)
	//fmt.Printf("%v %v\n", pattern, len(parts))
	n.insert(pattern, parts, 0)
}

func (r *Rounter) Search(methond string, pattern string) (*Node, map[string]string){
	n, is := r.root[methond]
	if !is{
		fmt.Printf("no methond %v", methond)
		return nil, make(map[string]string, 0)
	}

	parts := TransPattern(pattern)



	aim := n.search(parts, 0)
	if aim == nil{
		fmt.Printf("no pattern %v", pattern)
		return nil, make(map[string]string, 0)
	}

	diff := make(map[string]string, 0)

	fmt.Printf("%v\n", aim.pattern)

	_parts := TransPattern(aim.pattern)

	for i, item := range parts{
		if _parts[i] != item{
			diff[strings.TrimLeft(_parts[i], ":*")] = item
		}
	}
	return aim, diff
}

func (r *Rounter) Handle(c *Context){
	n, diff := r.Search(c.Method, c.Path)
	if n != nil{
		c.Params = diff

	for key, val := range c.Params{
		fmt.Printf(" %v, %v\n", key, val);
	}

		key := c.Method + "-" + n.pattern
		handle, is := r.handlers[key]
		//handle, is := r.handlers[n.pattern]
		// for key, val := range r.handlers{
		// 	fmt.Printf("%v %v\n", key, val)
		// }
		// fmt.Printf("%v\n", key)

		if is{
			handle(c)
		}else{
			c.String(http.StatusNotFound, "meizhaodao: %v\n", n.pattern)
		}
	}else{
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}


