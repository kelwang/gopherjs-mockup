package main

import (
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	js.Global.Get("document").Call("write", "ok")
	println("here")
}

func jsString(a interface{}) string {
	if a != nil {
		s := js.Global.Call("string", a).String()
		if s != "undefined" && s != "null" {
			return s
		}
	}
	return ""
}

type Svg struct {
	Width   float64
	Height  float64
	Content []SvgElement
}

func (svg Svg) String() string {
	result := `<svg width="` + jsString(svg.Width) + `" height="` + jsString(svg.Height) + `">`
	for _, v := range svg.Content {
		result += v.String()
	}
	result += `</svg>`
	return result
}

type SvgElement interface {
	String() string
}

type Rect struct {
	Width  float64
	Height float64
	X      float64
	Y      float64
	Rx     float64
	Ry     float64
}

type Circle struct {
}
