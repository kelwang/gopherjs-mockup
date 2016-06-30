package mockup

import (
	"github.com/kelwang/gopherjs-mockup/mockup/svg"
)

var EditablePrefix = "M_"

type Position struct {
	X float64
	Y float64
}

type Dimension struct {
	Width  float64
	Height float64
}

type MockupElement interface {
	Svg() svg.SvgElement
	Id() string

	GetWHXY() (float64, float64, float64, float64)
	MoveTo(x, y float64)
	ResizeTo(x, y, w, h float64)
	GetBase() BaseElement
}

type BaseElement struct {
	Position  Position
	Dimension Dimension
}

func (be *BaseElement) GetWHXY() (float64, float64, float64, float64) {
	return be.Dimension.Width, be.Dimension.Height, be.Position.X, be.Position.Y
}

func (be *BaseElement) GetBase() BaseElement {
	return *be
}

func (be *BaseElement) MoveTo(x, y float64) {
	be.Position.X = x
	be.Position.Y = y
}

func (be *BaseElement) ResizeTo(w, h float64) {
	be.Dimension.Width = w
	be.Dimension.Height = h
}

func newBaseElement(width, height, x, y float64) BaseElement {
	return BaseElement{
		Dimension: Dimension{
			Width:  width,
			Height: height,
		},
		Position: Position{
			X: x,
			Y: y,
		},
	}
}

type Text struct {
	Content string
	Color   string
	//	Size    int
}

type Thickness int

const (
	None Thickness = iota
	VeryThin
	Thin
	Medium
	Thick
	VeryThick
)

type Stroke struct {
	Color     string
	Thickness Thickness
}

func (thickness Thickness) Float64() float64 {
	return 0.5 * float64(thickness)
}

type idable struct {
	id string
}

func (id idable) Id() string {
	return id.id
}

func (i *idable) SetId(id string) {
	(*i).id = id
}

type textBox struct {
	idable
	BaseElement
	Text
	Stroke
}

var (
	WHITE    = "white"
	DARKGREY = "#555"
)

func NewTextBox(w, h, x, y float64, content string, id string) *textBox {
	return &textBox{
		BaseElement: newBaseElement(w, h, x, y),
		Text: Text{
			Content: content,
			Color:   DARKGREY,
		},
		Stroke: Stroke{
			Thickness: Medium,
			Color:     DARKGREY,
		},
		idable: idable{id: id},
	}
}

func (ele *textBox) Svg() svg.SvgElement {
	return &svg.Group{
		Editable: svg.EDITABLE,
		Idable: svg.Idable{
			ID: ele.id,
		},
		Content: []svg.SvgElement{
			&svg.Rect{
				Width:    ele.BaseElement.Dimension.Width,
				Height:   ele.BaseElement.Dimension.Height,
				X:        ele.BaseElement.Position.X,
				Y:        ele.BaseElement.Position.Y,
				Fillable: svg.NewFillable(WHITE, 1),
				Strokeable: svg.Strokeable{
					Stroke:      ele.Stroke.Color,
					StrokeWidth: ele.Stroke.Thickness.Float64(),
				},
				Idable: svg.Idable{ID: ele.idable.id + "_outter"},
			},
			&svg.Text{
				Content: ele.Text.Content,
				X:       ele.BaseElement.Position.X + (ele.BaseElement.Dimension.Width-float64(7*len(ele.Text.Content)))/2,
				Y:       ele.BaseElement.Position.Y + (ele.BaseElement.Dimension.Height+7)/2,
				Strokeable: svg.Strokeable{
					Stroke:      ele.Text.Color,
					StrokeWidth: ele.Stroke.Thickness.Float64(),
				},
				Idable: svg.Idable{ID: ele.idable.id + "_inner"},
			},
		},
	}
}

func (ele *textBox) MoveTo(x, y float64) {
	w, h, _, _ := ele.GetWHXY()
	content := ele.Svg().(*svg.Group).Content
	content[0].MoveTo(x, y)
	x1 := x + w/2 - float64(7*len(ele.Text.Content))/2
	y1 := y + (h+7)/2
	content[1].MoveTo(x1, y1)
	ele.BaseElement.MoveTo(x, y)
}

func (ele *textBox) ResizeTo(x, y, w, h float64) {
	content := ele.Svg().(*svg.Group).Content
	content[0].ResizeTo(w, h)
	ele.BaseElement.ResizeTo(w, h)
	content[0].MoveTo(x, y)
	ele.BaseElement.MoveTo(x, y)
	x1 := x + w/2 - float64(7*len(ele.Text.Content))/2
	y1 := y + (h+7)/2
	content[1].MoveTo(x1, y1)
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

type button struct {
	idable
	BaseElement
	Text
	Stroke
}

func NewButton(w, h, x, y float64, content string, id string) *button {
	return &button{
		idable:      idable{id: id},
		BaseElement: newBaseElement(w, h, x, y),
		Text: Text{
			Content: content,
			Color:   DARKGREY,
		},
		Stroke: Stroke{
			Thickness: Medium,
			Color:     DARKGREY,
		},
	}
}

func (ele *button) Svg() svg.SvgElement {
	return &svg.Group{
		Idable: svg.Idable{
			ID: ele.idable.id,
		},
		Editable: svg.EDITABLE,
		Content: []svg.SvgElement{
			&svg.Rect{
				Width:    ele.BaseElement.Dimension.Width,
				Height:   ele.BaseElement.Dimension.Height,
				X:        ele.BaseElement.Position.X,
				Y:        ele.BaseElement.Position.Y,
				RX:       min(ele.BaseElement.Dimension.Width, ele.BaseElement.Dimension.Height) / 4,
				RY:       min(ele.BaseElement.Dimension.Width, ele.BaseElement.Dimension.Height) / 4,
				Fillable: svg.NewFillable(WHITE, 1),
				Strokeable: svg.Strokeable{
					Stroke:      ele.Stroke.Color,
					StrokeWidth: ele.Stroke.Thickness.Float64(),
				},
				Idable: svg.Idable{
					ID: ele.idable.id + "_outer",
				},
			},
			&svg.Text{
				Content: ele.Text.Content,
				X:       ele.BaseElement.Position.X + (ele.BaseElement.Dimension.Width-float64(7*len(ele.Text.Content)))/2,
				Y:       ele.BaseElement.Position.Y + (ele.BaseElement.Dimension.Height+7)/2,
				Strokeable: svg.Strokeable{
					Stroke:      ele.Text.Color,
					StrokeWidth: ele.Stroke.Thickness.Float64(),
				},
				Idable: svg.Idable{
					ID: ele.idable.id + "_inner",
				},
			},
		},
	}
}

func (ele *button) MoveTo(x, y float64) {
	w, h, _, _ := ele.GetWHXY()
	content := ele.Svg().(*svg.Group).Content
	content[0].MoveTo(x, y)
	x1 := x + w/2 - float64(7*len(ele.Text.Content))/2
	y1 := y + (h+7)/2
	content[1].MoveTo(x1, y1)
	ele.BaseElement.MoveTo(x, y)
}

func (ele *button) ResizeTo(x, y, w, h float64) {
	content := ele.Svg().(*svg.Group).Content
	content[0].ResizeTo(w, h)
	ele.BaseElement.ResizeTo(w, h)
	content[0].MoveTo(x, y)
	ele.BaseElement.MoveTo(x, y)
	x1 := x + w/2 - float64(7*len(ele.Text.Content))/2
	y1 := y + (h+7)/2
	content[1].MoveTo(x1, y1)
}

type box struct {
	idable
	BaseElement
	Stroke
}

func NewBox(w, h, x, y float64, id string) *box {
	return &box{
		idable:      idable{id: id},
		BaseElement: newBaseElement(w, h, x, y),
		Stroke: Stroke{
			Thickness: Medium,
			Color:     DARKGREY,
		},
	}
}

func (ele *box) Svg() svg.SvgElement {
	return &svg.Rect{
		Width:    ele.BaseElement.Dimension.Width,
		Height:   ele.BaseElement.Dimension.Height,
		X:        ele.BaseElement.Position.X,
		Y:        ele.BaseElement.Position.Y,
		RX:       min(ele.BaseElement.Dimension.Width, ele.BaseElement.Dimension.Height) / 8,
		RY:       min(ele.BaseElement.Dimension.Width, ele.BaseElement.Dimension.Height) / 8,
		Fillable: svg.NewFillable(WHITE, 1),
		Strokeable: svg.Strokeable{
			Stroke:      ele.Stroke.Color,
			StrokeWidth: ele.Stroke.Thickness.Float64(),
		},
		Editable: svg.EDITABLE,
		Idable: svg.Idable{
			ID: ele.id,
		},
	}
}

func (ele *box) MoveTo(x, y float64) {
	ele.Svg().MoveTo(x, y)
	ele.BaseElement.MoveTo(x, y)
}

func (ele *box) ResizeTo(x, y, w, h float64) {
	ele.Svg().ResizeTo(w, h)
	ele.BaseElement.ResizeTo(w, h)
	ele.Svg().MoveTo(x, y)
	ele.BaseElement.MoveTo(x, y)

}

type label struct {
	idable
	BaseElement
	Text
	Stroke
	svg.Editable
}

func NewLabel(w, h, x, y float64, content string, id string, editable svg.Editable) *label {
	return &label{
		idable:      idable{id: id},
		BaseElement: newBaseElement(w, h, x, y),
		Text: Text{
			Content: content,
			Color:   DARKGREY,
		},
		Stroke: Stroke{
			Thickness: Medium,
			Color:     DARKGREY,
		},
		Editable: editable,
	}
}

func (ele *label) Svg() svg.SvgElement {
	return &svg.Group{
		Editable: ele.Editable,
		Idable: svg.Idable{
			ID: ele.id,
		},
		Content: []svg.SvgElement{
			&svg.Rect{
				Width:    ele.BaseElement.Dimension.Width,
				Height:   ele.BaseElement.Dimension.Height,
				X:        ele.BaseElement.Position.X,
				Y:        ele.BaseElement.Position.Y,
				Fillable: svg.NewFillable(WHITE, 0),
				Idable:   svg.Idable{ID: ele.idable.id + "_outter"},
			},
			&svg.Text{
				Content: ele.Text.Content,
				X:       ele.BaseElement.Position.X + (ele.BaseElement.Dimension.Width-float64(7*len(ele.Text.Content)))/2,
				Y:       ele.BaseElement.Position.Y + (ele.BaseElement.Dimension.Height+7)/2,
				Strokeable: svg.Strokeable{
					Stroke:      ele.Text.Color,
					StrokeWidth: ele.Stroke.Thickness.Float64(),
				},
				Editable: svg.EDITABLE,
				Idable:   svg.Idable{ID: ele.idable.id + "_inner"},
			},
		},
	}
}

func (ele *label) MoveTo(x, y float64) {
	w, h, _, _ := ele.GetWHXY()
	content := ele.Svg().(*svg.Group).Content
	content[0].MoveTo(x, y)
	x1 := x + w/2 - float64(7*len(ele.Text.Content))/2
	y1 := y + (h+7)/2
	content[1].MoveTo(x1, y1)
	ele.BaseElement.MoveTo(x, y)
}

func (ele *label) ResizeTo(x, y, w, h float64) {
	content := ele.Svg().(*svg.Group).Content
	content[0].ResizeTo(w, h)
	ele.BaseElement.ResizeTo(w, h)
	content[0].MoveTo(x, y)
	ele.BaseElement.MoveTo(x, y)
	x1 := x + w/2 - float64(7*len(ele.Text.Content))/2
	y1 := y + (h+7)/2
	content[1].MoveTo(x1, y1)
}

type line struct {
	idable
	BaseElement
	Stroke
}

func NewLine(w, h, x, y float64, id string) *line {
	return &line{
		idable:      idable{id: id},
		BaseElement: newBaseElement(w, h, x, y),
		Stroke: Stroke{
			Thickness: Medium,
			Color:     DARKGREY,
		},
	}
}

func (ele *line) Svg() svg.SvgElement {
	return &svg.Line{
		X1: ele.BaseElement.Position.X,
		Y1: ele.BaseElement.Position.Y,
		X2: ele.BaseElement.Position.X + ele.BaseElement.Dimension.Width,
		Y2: ele.BaseElement.Position.Y + ele.BaseElement.Dimension.Height,
		Strokeable: svg.Strokeable{
			Stroke:      DARKGREY,
			StrokeWidth: ele.Stroke.Thickness.Float64(),
		},
		Editable: svg.LINABLE,
		Idable: svg.Idable{
			ID: ele.id,
		},
	}
}

func (ele *line) MoveTo(x, y float64) {
	ele.BaseElement.MoveTo(x, y)
	ele.Svg().MoveTo(x, y)
}

func (ele *line) ResizeTo(x, y, w, h float64) {

}

func (ele *line) PointTo(x, y float64, pt int) {
	if pt == 1 {
		ele.BaseElement.ResizeTo(ele.BaseElement.Position.X-x+ele.BaseElement.Dimension.Width, ele.BaseElement.Position.Y-y+ele.BaseElement.Dimension.Height)
		ele.BaseElement.MoveTo(x, y)

	} else {
		ele.BaseElement.ResizeTo(x-ele.BaseElement.Position.X, y-ele.BaseElement.Position.Y)
	}
	ele.Svg().(*svg.Line).PointTo(x, y, pt)
}

type ScaleBox struct {
	idable
	MockupElement
}

func NewScaleBox(mockupElement MockupElement) *ScaleBox {
	return &ScaleBox{
		idable:        idable{id: EditablePrefix + mockupElement.Id()},
		MockupElement: mockupElement,
	}
}

func (ele *ScaleBox) Id() string {
	return ele.idable.id
}

func (ele *ScaleBox) MoveTo(x, y float64) {
	content := ele.Svg().(*svg.Group).Content
	w, h, _, _ := ele.MockupElement.GetWHXY()
	content[1].MoveTo(x-square_height/2, y-square_height/2)
	content[2].MoveTo(x+w/2-square_height/2, y-square_height/2)
	content[3].MoveTo(x-square_height/2, y+h/2-square_height/2)
	content[4].MoveTo(x-square_height/2, y+h-square_height/2)
	content[5].MoveTo(x+w-square_height/2, y-square_height/2)
	content[6].MoveTo(x+w/2-square_height/2, y+h-square_height/2)
	content[7].MoveTo(x+w-square_height/2, y+h/2-square_height/2)
	content[8].MoveTo(x+w-square_height/2, y+h-square_height/2)

	ele.MockupElement.MoveTo(x, y)
}

var stroke_width = float64(2)
var square_height = float64(8)

func (ele *ScaleBox) Svg() svg.SvgElement {

	w, h, x, y := ele.MockupElement.GetWHXY()
	return &svg.Group{
		Content: []svg.SvgElement{
			ele.MockupElement.Svg(),
			scaleboxRect(x-square_height/2, y-square_height/2, stroke_width, square_height, "sq1_"+ele.idable.id, svg.NWSE_RESIZABLE),
			scaleboxRect(x+w/2-square_height/2, y-square_height/2, stroke_width, square_height, "sq2_"+ele.idable.id, svg.NS_RESIZABLE),
			scaleboxRect(x+w-square_height/2, y-square_height/2, stroke_width, square_height, "sq3_"+ele.idable.id, svg.NESW_RESIZABLE),
			scaleboxRect(x-square_height/2, y+h/2-square_height/2, stroke_width, square_height, "sq4_"+ele.idable.id, svg.EW_RESIZABLE),
			scaleboxRect(x+w-square_height/2, y+h/2-square_height/2, stroke_width, square_height, "sq5_"+ele.idable.id, svg.EW_RESIZABLE),
			scaleboxRect(x-square_height/2, y+h-square_height/2, stroke_width, square_height, "sq6_"+ele.idable.id, svg.NESW_RESIZABLE),
			scaleboxRect(x+w/2-square_height/2, y+h-square_height/2, stroke_width, square_height, "sq7_"+ele.idable.id, svg.NS_RESIZABLE),
			scaleboxRect(x+w-square_height/2, y+h-square_height/2, stroke_width, square_height, "sq8_"+ele.idable.id, svg.NWSE_RESIZABLE),
		},
		Editable: svg.DRAGGABLE,
		Fillable: svg.NewFillable(WHITE, 1),
		Idable: svg.Idable{
			ID: ele.idable.id,
		},
	}

}

func (ele *ScaleBox) NWResizeTo(x, y float64) {
	w0, h0, x0, y0 := ele.MockupElement.GetWHXY()
	content := ele.Svg().(*svg.Group).Content
	w := w0 + x0 - x
	h := h0 + y0 - y

	content[1].MoveTo(x-square_height/2, y-square_height/2)
	content[2].MoveTo((x+content[8].(*svg.Rect).X)/2-square_height/2, y-square_height/2)
	content[3].MoveTo(content[8].(*svg.Rect).X, y-square_height/2)
	content[4].MoveTo(x-square_height/2, (y+content[8].(*svg.Rect).Y)/2)
	content[5].MoveTo(content[8].(*svg.Rect).X, (y+content[8].(*svg.Rect).Y)/2)
	content[6].MoveTo(x-square_height/2, content[8].(*svg.Rect).Y)
	content[7].MoveTo((x+content[8].(*svg.Rect).X)/2-square_height/2, content[8].(*svg.Rect).Y)

	ele.MockupElement.ResizeTo(x, y, w, h)
}

func (ele *ScaleBox) NResizeTo(x, y float64) {
	w0, h0, x0, y0 := ele.MockupElement.GetWHXY()
	content := ele.Svg().(*svg.Group).Content

	h := h0 + y0 - y

	content[1].MoveTo(x0-square_height/2, y-square_height/2)
	content[2].MoveTo(content[2].(*svg.Rect).X-square_height/2, y-square_height/2)
	content[3].MoveTo(content[3].(*svg.Rect).X-square_height/2, y-square_height/2)
	content[4].MoveTo(x0-square_height/2, (y+content[8].(*svg.Rect).Y)/2)
	content[5].MoveTo(content[5].(*svg.Rect).X, (y+content[8].(*svg.Rect).Y)/2)

	ele.MockupElement.ResizeTo(x0, y, w0, h)
}

func (ele *ScaleBox) NEResizeTo(x, y float64) {
	_, h0, x0, y0 := ele.MockupElement.GetWHXY()
	content := ele.Svg().(*svg.Group).Content
	w := x - x0
	h := h0 + y0 - y

	content[1].MoveTo(content[6].(*svg.Rect).X, y-square_height/2)
	content[2].MoveTo((x+content[6].(*svg.Rect).X)/2, y-square_height/2)
	content[3].MoveTo(x-square_height/2, y-square_height/2)
	content[4].MoveTo(content[6].(*svg.Rect).X, (y+content[6].(*svg.Rect).Y)/2)
	content[5].MoveTo(x-square_height/2, (y+content[6].(*svg.Rect).Y)/2)
	content[7].MoveTo((x+content[6].(*svg.Rect).X)/2, content[6].(*svg.Rect).Y)
	content[8].MoveTo(x-square_height/2, content[6].(*svg.Rect).Y)

	ele.MockupElement.ResizeTo(x0, y, w, h)
}

func (ele *ScaleBox) WResizeTo(x, y float64) {
	w0, h0, x0, y0 := ele.MockupElement.GetWHXY()
	content := ele.Svg().(*svg.Group).Content

	w := w0 + x0 - x

	content[1].MoveTo(x-square_height/2, content[1].(*svg.Rect).Y)
	content[2].MoveTo((x+content[3].(*svg.Rect).X)/2, content[3].(*svg.Rect).Y)
	content[4].MoveTo(x-square_height/2, content[4].(*svg.Rect).Y)
	content[6].MoveTo(x-square_height/2, content[6].(*svg.Rect).Y)
	content[7].MoveTo((x+content[3].(*svg.Rect).X)/2, content[7].(*svg.Rect).Y)

	ele.MockupElement.ResizeTo(x, y0, w, h0)
}

func (ele *ScaleBox) EResizeTo(x, y float64) {
	_, h0, x0, y0 := ele.MockupElement.GetWHXY()
	content := ele.Svg().(*svg.Group).Content

	w := x - x0

	content[3].MoveTo(x-square_height/2, content[3].(*svg.Rect).Y)
	content[2].MoveTo((x0+content[3].(*svg.Rect).X)/2, content[3].(*svg.Rect).Y)
	content[5].MoveTo(x-square_height/2, content[5].(*svg.Rect).Y)
	content[7].MoveTo((x0+content[3].(*svg.Rect).X)/2, content[7].(*svg.Rect).Y)
	content[8].MoveTo(x-square_height/2, content[8].(*svg.Rect).Y)

	ele.MockupElement.ResizeTo(x0, y0, w, h0)
}

func (ele *ScaleBox) SWResizeTo(x, y float64) {
	w0, _, x0, y0 := ele.MockupElement.GetWHXY()
	content := ele.Svg().(*svg.Group).Content
	w := w0 + x0 - x
	h := y - y0

	content[1].MoveTo(x-square_height/2, y0)
	content[2].MoveTo((x+content[3].(*svg.Rect).X)/2-square_height/2, y0)
	content[4].MoveTo(x-square_height/2, (y+y0)/2)
	content[5].MoveTo(content[3].(*svg.Rect).X, (y+content[3].(*svg.Rect).Y)/2)
	content[6].MoveTo(x-square_height/2, y-square_height/2)
	content[7].MoveTo((x+content[3].(*svg.Rect).X)/2, y-square_height/2)
	content[8].MoveTo(content[3].(*svg.Rect).X, y-square_height/2)

	ele.MockupElement.ResizeTo(x, y-h, w, h)
}

func (ele *ScaleBox) SResizeTo(x, y float64) {
	w0, _, x0, y0 := ele.MockupElement.GetWHXY()
	content := ele.Svg().(*svg.Group).Content
	h := y - y0

	content[4].MoveTo(x0-square_height/2, (y+content[3].(*svg.Rect).Y)/2)
	content[5].MoveTo(content[3].(*svg.Rect).X, (y+content[3].(*svg.Rect).Y)/2)
	content[6].MoveTo(x0-square_height/2, y-square_height/2)
	content[7].MoveTo(content[2].(*svg.Rect).X, y-square_height/2)
	content[8].MoveTo(content[3].(*svg.Rect).X, y-square_height/2)

	ele.MockupElement.ResizeTo(x0, y0, w0, h)
}

func (ele *ScaleBox) SEesizeTo(x, y float64) {
	_, _, x0, y0 := ele.MockupElement.GetWHXY()
	content := ele.Svg().(*svg.Group).Content
	w := x - x0
	h := y - y0

	content[2].MoveTo((x+x0)/2, y0-square_height/2)
	content[3].MoveTo(x-square_height/2, content[3].(*svg.Rect).Y)
	content[4].MoveTo(x0-square_height/2, (y+content[3].(*svg.Rect).Y)/2)
	content[5].MoveTo(content[3].(*svg.Rect).X, (y+content[3].(*svg.Rect).Y)/2)
	content[6].MoveTo(x0-square_height/2, y-square_height/2)
	content[7].MoveTo((x+x0-square_height/2)/2, y-square_height/2)
	content[8].MoveTo(content[3].(*svg.Rect).X, y-square_height/2)

	ele.MockupElement.ResizeTo(x0, y0, w, h)

}

func scaleboxRect(x, y, stroke_width, square_height float64, id string, ed svg.Editable) *svg.Rect {
	return &svg.Rect{
		X:        x,
		Y:        y,
		Width:    square_height,
		Height:   square_height,
		Fillable: svg.NewFillable(WHITE, 1),
		Strokeable: svg.Strokeable{
			Stroke:      DARKGREY,
			StrokeWidth: stroke_width,
		},
		Idable: svg.Idable{
			ID: id,
		},
		Editable: ed,
	}
}

type ScaleLine struct {
	idable
	*line
}

func NewScaleLine(l MockupElement) *ScaleLine {
	ll := l.(*line)
	return &ScaleLine{
		idable: idable{id: EditablePrefix + ll.idable.id},
		line:   ll,
	}
}

func (ele *ScaleLine) Line() *line {
	return ele.line
}

func (ele *ScaleLine) Svg() svg.SvgElement {
	w, h, x, y := ele.line.GetWHXY()
	return &svg.Group{
		Content: []svg.SvgElement{
			ele.line.Svg(),
			scaleboxRect(x-square_height/2, y-square_height/2, stroke_width, square_height, "sql1_"+ele.idable.id, svg.LINE_VERTEX),
			scaleboxRect(x+w-square_height/2, y+h-square_height/2, stroke_width, square_height, "sql2_"+ele.idable.id, svg.LINE_VERTEX),
		},
		Editable: svg.DRAGGABLE,
		Fillable: svg.NewFillable(WHITE, 1),
		Idable: svg.Idable{
			ID: ele.idable.id,
		},
	}
}

func (ele *ScaleLine) MoveTo(x, y float64) {
}

func (ele *ScaleLine) ResizeTo(x, y, w, h float64) {

}

func (ele *ScaleLine) PointTo(x, y float64, sqr int) {
	content := ele.Svg().(*svg.Group).Content
	content[sqr].MoveTo(x-square_height/2, y-square_height/2)
	ele.line.PointTo(x, y, sqr)
}

type CloneBox struct {
	idable
	MockupElement
}

func newCloneBox(ele MockupElement, id string) *CloneBox {
	return &CloneBox{
		idable: idable{
			id: id,
		},
		MockupElement: ele,
	}
}

func (ele *CloneBox) Id() string {
	return ele.idable.id
}

func (ele *CloneBox) Svg() svg.SvgElement {
	mockup := ele.MockupElement
	eleSvg := ele.MockupElement.Svg()
	eleSvg.(*svg.Group).Editable = svg.EDITABLE
	eleSvg.(*svg.Group).ID = ele.idable.id
	println(eleSvg.String())
	return mockup.Svg()
}

func (ele *CloneBox) MoveTo(x, y float64) {
	ele.MockupElement.MoveTo(x, y)
}
