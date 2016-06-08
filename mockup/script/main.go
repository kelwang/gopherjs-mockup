package main

import (
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	js.Global.Get("document").Call("write", "ok")
	println("here")
}
