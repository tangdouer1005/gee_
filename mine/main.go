package main

import (
	"fmt"
	"net/http"
	"gee"
)

func main()  {
	G := gee.New()
	G.GET("/1000", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "get http from %v/1000\n", r.URL.Path)
	})
	G.Run(":9999")
}