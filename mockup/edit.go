package mockup

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/jquery"
	"github.com/kelwang/gopherjs-mockup/mockup/svg"
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

func jsInt(s interface{}) int {
	if s != nil && s != "" {
		return js.Global.Call("parseInt", s).Int()
	}
	return 0
}

type Editable struct {
	Movable
	Scalable
	Border
}

type Border struct {
	X1 float64
	Y1 float64
	X2 float64
	Y2 float64
}

func NewEditable(x1, y1, x2, y2 float64) *Editable {
	return &Editable{
		Movable:  movableNil,
		Scalable: scalableNill,
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

type Scalable struct {
	jquery.JQuery
}

var movableNil = Movable{JQuery: jQuery(nil)}
var scalableNill = Scalable{JQuery: jQuery(nil)}

func (ed *Editable) BindEvents(m map[string]MockupElement) {
	//dragging
	jQuery(document).On(jquery.MOUSEDOWN, svg.DRAGGABLE.JqSelector(), ed.startDragging)
	jQuery(document).On(jquery.MOUSEMOVE, func(e jquery.Event) { ed.dragging(e, m) })

	//scaling
	jQuery(document).On(jquery.MOUSEOVER, svg.EW_RESIZABLE.JqSelector(), ed.ewResizeMouseOver)
	jQuery(document).On(jquery.MOUSEOVER, svg.NS_RESIZABLE.JqSelector(), ed.nsResizeMouseOver)
	jQuery(document).On(jquery.MOUSEOVER, svg.NESW_RESIZABLE.JqSelector(), ed.neswResizeMouseOver)
	jQuery(document).On(jquery.MOUSEOVER, svg.NWSE_RESIZABLE.JqSelector(), ed.nwseResizeMouseOver)

	jQuery(document).On(jquery.MOUSEDOWN, svg.NWSE_RESIZABLE.JqSelector(), ed.startResize)
	jQuery(document).On(jquery.MOUSEDOWN, svg.NESW_RESIZABLE.JqSelector(), ed.startResize)
	jQuery(document).On(jquery.MOUSEDOWN, svg.NS_RESIZABLE.JqSelector(), ed.startResize)
	jQuery(document).On(jquery.MOUSEDOWN, svg.EW_RESIZABLE.JqSelector(), ed.startResize)

	// stopping
	jQuery(document).On(jquery.MOUSEUP, ed.stopDraggingResize)

}

func (ed *Editable) startResize(e jquery.Event) {
	ed.Movable = movableNil
	ed.Scalable = Scalable{JQuery: jQuery(e.CurrentTarget)}
}

func (ed *Editable) resizing(e jquery.Event, m map[string]MockupElement) {
	this := jQuery(e.CurrentTarget)
	id := this.Attr("id")
	ele := m[id[4:]]
	println(ele.Id())

}

func (ed *Editable) ewResizeMouseOver(e jquery.Event) {
	jQuery(e.CurrentTarget).SetCss("cursor", "ew-resize")
}

func (ed *Editable) nsResizeMouseOver(e jquery.Event) {
	jQuery(e.CurrentTarget).SetCss("cursor", "ns-resize")
}

func (ed *Editable) neswResizeMouseOver(e jquery.Event) {
	jQuery(e.CurrentTarget).SetCss("cursor", "nesw-resize")
}

func (ed *Editable) nwseResizeMouseOver(e jquery.Event) {
	jQuery(e.CurrentTarget).SetCss("cursor", "nwse-resize")
}

func (ed *Editable) dragging(e jquery.Event, m map[string]MockupElement) {

	clientX := e.Get("offsetX").Float()
	clientY := e.Get("offsetY").Float()

	if ed.Movable != movableNil {
		id := ed.Movable.Attr("id")
		ele := m[id]
		width := float64(ed.Movable.Width())
		height := float64(ed.Movable.Height())
		if clientX-width >= ed.X1 && clientX < ed.X2 && clientY-height >= ed.Y1 && clientY < ed.Y2 {
			ele.MoveTo(clientX-width, clientY)
		}
	}

	if ed.Scalable != scalableNill {
		id := ed.Scalable.Attr("id")
		sqr := id[2:3]
		id = id[4:]
		ele := m[id]
		switch sqr {
		case "1":
			ele.(*ScaleBox).NWResizeTo(clientX, clientY)
		case "2":
			ele.(*ScaleBox).NResizeTo(clientX, clientY)
		case "3":
			ele.(*ScaleBox).NEResizeTo(clientX, clientY)
		case "4":
			ele.(*ScaleBox).WResizeTo(clientX, clientY)
		case "5":
			ele.(*ScaleBox).EResizeTo(clientX, clientY)
		case "6":
			ele.(*ScaleBox).SWResizeTo(clientX, clientY)
		case "7":
			ele.(*ScaleBox).SResizeTo(clientX, clientY)
		case "8":
			ele.(*ScaleBox).SEesizeTo(clientX, clientY)
		}

	}
}

func (ed *Editable) startDragging(e jquery.Event) {
	if ed.Scalable == scalableNill {
		println("start dragging")
		ed.Movable = Movable{JQuery: jQuery(e.CurrentTarget)}
		ed.Movable.SetCss("cursor", "move")
	}

}

func (ed *Editable) stopDraggingResize(e jquery.Event) {
	jQuery(e.CurrentTarget).SetCss("cursor", "pointer")
	if ed.Movable != movableNil {
		ed.Movable = movableNil
	}
	if ed.Scalable != scalableNill {
		ed.Scalable = scalableNill
	}
}
