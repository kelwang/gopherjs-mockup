package main

import (
	"fmt"
	"net/http"

	gbuild "github.com/gopherjs/gopherjs/build"
	"github.com/kelwang/gopherjs/tool"
)

var basePath = "github.com/kelwang/gopherjs-mockup/"

// Default Request Handler
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `<html><head><title>GopherJS Mockup</title></head><body><script type="text/javascript" src="//cdnjs.cloudflare.com/ajax/libs/jquery/2.2.4/jquery.min.js"></script><script src="/mockup/script/script.js"></script></body></html>`)
}

func main() {
	options := &gbuild.Options{CreateMapFile: true}
	http.Handle("/mockup/", tool.Handler(basePath+"mockup/", options, len("/mockup/")))
	http.HandleFunc("/", defaultHandler)
	http.ListenAndServe(":9390", nil)
}
