package mockup

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/jquery"
)

var jQuery = jquery.NewJQuery
var document = js.Global.Get("document")
var console = js.Global.Get("console")

func jsFloat(s interface{}) float64 {
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

func (dr *Draggable) BindEvents(m map[string]MockupElement) {
	jQuery(document).On(jquery.MOUSEDOWN, ".draggable", dr.start)
	jQuery(document).On(jquery.MOUSEUP, ".draggable", dr.stop)
	jQuery(document).On(jquery.MOUSEMOVE, func(e jquery.Event) {
		dr.dragging(e, m)
	})
}

func (dr *Draggable) dragging(e jquery.Event, m map[string]MockupElement) {
	//container := jQuery("svg")
	if dr.Movable != MovableNil {
		id := dr.Movable.Attr("id")
		ele := m[id]
		width := float64(dr.Movable.Width())
		height := float64(dr.Movable.Height())
		clientX := e.Get("offsetX").Float()
		clientY := e.Get("offsetY").Float()
		if clientX-width >= dr.X1 && clientX < dr.X2 && clientY-height >= dr.Y1 && clientY < dr.Y2 {
			ele.MoveTo(clientX-width, clientY)
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
