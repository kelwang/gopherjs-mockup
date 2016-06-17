package mockup

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/jquery"
)

var jQuery = jquery.NewJQuery
var document = js.Global.Get("document")
var console = js.Global.Get("console")

func jsFloat(s interface{}) float64 {
	println(s)
	if s != nil && s != "" {
		return js.Global.Call("parseFloat", s).Float()
	}
	return 0
}

type Draggable struct {
	Movable
	Border
}

type Border struct {
	X1 float64
	Y1 float64
	X2 float64
	Y2 float64
}

func NewDraggable(ele Movable, x1, y1, x2, y2 float64) *Draggable {
	return &Draggable{
		Movable: ele,
		Border: Border{
			X1: x1,
			Y1: y1,
			X2: x2,
			Y2: y2,
		},
	}
}

type Movable struct {
	jquery.JQuery
}

var MovableNil = Movable{JQuery: jQuery(nil)}

func (dr *Draggable) BindEvents() {
	jQuery(document).On(jquery.MOUSEDOWN, ".draggable", dr.start)
	jQuery(document).On(jquery.MOUSEUP, ".draggable", dr.stop)
	jQuery(document).On(jquery.MOUSEMOVE, dr.dragging)
}

func (dr *Draggable) dragging(e jquery.Event) {
	if dr.Movable != MovableNil {
		width := float64(dr.Movable.Width())
		height := float64(dr.Movable.Height())
		clientX := e.Get("offsetX").Float()
		clientY := e.Get("offsetY").Float()
		if clientX-width >= dr.X1 && clientX < dr.X2 && clientY-height >= dr.Y1 && clientY < dr.Y2 {
			diffX := clientX - width - jsFloat(dr.Movable.Attr("x"))
			diffY := clientY - jsFloat(dr.Movable.Attr("y"))
			dr.Movable.Children("").Each(func(i int, ele interface{}) {
				el := jQuery(ele)
				el.SetAttr("x", jsFloat(el.Attr("x"))+diffX)
				el.SetAttr("y", jsFloat(el.Attr("y"))+diffY)
			})
			dr.Movable.SetAttr("x", jsFloat(dr.Movable.Attr("x"))+diffX)
			dr.Movable.SetAttr("y", jsFloat(dr.Movable.Attr("y"))+diffY)
		}
	}
}

func (dr *Draggable) start(e jquery.Event) {
	dr.Movable = Movable{JQuery: jQuery(e.CurrentTarget)}
	dr.Movable.SetCss("cursor", "move")
}

func (dr *Draggable) stop(e jquery.Event) {
	jQuery(e.CurrentTarget).SetCss("cursor", "auto")
	dr.Movable = MovableNil
}
