package mockup

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/jquery"
)

var jQuery = jquery.NewJQuery
var document = js.Global.Get("document")
var console = js.Global.Get("console")

type Draggable struct {
	//make sure there is only one movable
	Movable jquery.JQuery
	Border
}

type Border struct {
	X1 float64
	Y1 float64
	X2 float64
	Y2 float64
}

func NewDraggable(ele jquery.JQuery, x1, y1, x2, y2 float64) *Draggable {
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

func (dr *Draggable) BindEvents() {
	jQuery(document).On(jquery.MOUSEDOWN, ".draggable", dr.start)
	jQuery(document).On(jquery.MOUSEUP, ".draggable", dr.stop)
	jQuery(document).On(jquery.MOUSEMOVE, dr.dragging)
}

func (dr *Draggable) dragging(e jquery.Event) {
	if dr.Movable != jQuery(nil) {
		width := float64(dr.Movable.Width())
		height := float64(dr.Movable.Height())
		clientX := e.Get("offsetX").Float()
		clientY := e.Get("offsetY").Float()
		if clientX-width >= dr.X1 && clientX < dr.X2 && clientY-height >= dr.Y1 && clientY < dr.Y2 {
			dr.Movable.SetAttr("x", clientX-width)
			dr.Movable.SetAttr("y", clientY)
		}
	}
}
func (dr *Draggable) start(e jquery.Event) {
	dr.Movable = jQuery(e.CurrentTarget)
	dr.Movable.SetCss("cursor", "move")
}

func (dr *Draggable) stop(e jquery.Event) {
	jQuery(e.CurrentTarget).SetCss("cursor", "auto")
	dr.Movable = jQuery(nil)
}