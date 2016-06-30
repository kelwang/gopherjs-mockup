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
	LineMovable
	Clonable
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
		Movable:     movableNil,
		Scalable:    scalableNill,
		LineMovable: lineMovableNil,
		Clonable:    clonableNil,
		Border: Border{
			X1: x1,
			Y1: y1,
			X2: x2,
			Y2: y2,
		},
	}
}

type LineMovable struct {
	jquery.JQuery
}

type Movable struct {
	jquery.JQuery
}

type Scalable struct {
	jquery.JQuery
}

type Clonable struct {
	jquery.JQuery
}

var movableNil = Movable{JQuery: jQuery(nil)}
var scalableNill = Scalable{JQuery: jQuery(nil)}
var lineMovableNil = LineMovable{JQuery: jQuery(nil)}
var clonableNil = Clonable{JQuery: jQuery(nil)}

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

	//line moving
	jQuery(document).On(jquery.MOUSEDOWN, svg.LINE_VERTEX.JqSelector(), ed.startLineEditing)

	// stopping
	jQuery(document).On(jquery.MOUSEUP, ed.stopDraggingResize)

	// clonable
	jQuery(document).On(jquery.MOUSEDOWN, svg.CLONABLE.JqSelector(), func(e jquery.Event) {
		ed.startClone(e, m)
	})

}

func (ed *Editable) startClone(e jquery.Event, m map[string]MockupElement) {
	ed.Movable = movableNil
	ed.Scalable = scalableNill
	ed.LineMovable = lineMovableNil
	id := jQuery(e.CurrentTarget).Attr("id")
	ele := m[id]

	clo := newCloneBox(ele, "M1")
	m["M1"] = clo
	cloJq := clo.Svg().Jq()
	jQuery("svg").Append(cloJq)

	ed.Clonable = Clonable{JQuery: cloJq}
}

func (ed *Editable) startResize(e jquery.Event) {
	ed.Movable = movableNil
	ed.Scalable = Scalable{JQuery: jQuery(e.CurrentTarget)}
	ed.LineMovable = lineMovableNil
	ed.Clonable = clonableNil
}

func (ed *Editable) startLineEditing(e jquery.Event) {
	ed.Movable = movableNil
	ed.Scalable = scalableNill
	ed.LineMovable = LineMovable{JQuery: jQuery(e.CurrentTarget)}
	ed.Clonable = clonableNil
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

	if ed.LineMovable != lineMovableNil {
		id := ed.LineMovable.Attr("id")
		ele := m[id[5:]]
		sqr := id[3:4]
		ele.(*ScaleLine).PointTo(clientX, clientY, jsInt(sqr))
	}

	if ed.Clonable != clonableNil {
		id := ed.Clonable.Attr("id")
		ele := m[id]
		ele.MoveTo(clientX, clientY)
		println(id)
	}

}

func (ed *Editable) startDragging(e jquery.Event) {
	if ed.Scalable == scalableNill {
		ed.Movable = Movable{JQuery: jQuery(e.CurrentTarget)}
		ed.Movable.SetCss("cursor", "move")
	}

	ed.Scalable = scalableNill
	ed.LineMovable = lineMovableNil
	ed.Clonable = clonableNil
}

func (ed *Editable) stopDraggingResize(e jquery.Event) {
	jQuery(e.CurrentTarget).SetCss("cursor", "pointer")
	ed.Movable = movableNil
	ed.Scalable = scalableNill
	ed.LineMovable = lineMovableNil
	ed.Clonable = clonableNil
}
