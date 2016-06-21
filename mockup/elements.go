package mockup

import (
	"github.com/kelwang/gopherjs-mockup/mockup/svg"
)

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
}

type BaseElement struct {
	Position  Position
	Dimension Dimension
}

func (be *BaseElement) GetWHXY() (float64, float64, float64, float64) {
	return be.Dimension.Width, be.Dimension.Height, be.Position.X, be.Position.Y
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
		Draggable: svg.Draggable{
			Draggable: true,
		},
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
			},
			&svg.Text{
				Content: ele.Text.Content,
				X:       ele.BaseElement.Position.X + (ele.BaseElement.Dimension.Width-float64(7*len(ele.Text.Content)))/2,
				Y:       ele.BaseElement.Position.Y + (ele.BaseElement.Dimension.Height+7)/2,
				Strokeable: svg.Strokeable{
					Stroke:      ele.Text.Color,
					StrokeWidth: ele.Stroke.Thickness.Float64(),
				},
			},
		},
	}
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
			ID: ele.id,
		},
		Draggable: svg.Draggable{
			Draggable: true,
		},
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
			},
			&svg.Text{
				Content: ele.Text.Content,
				X:       ele.BaseElement.Position.X + (ele.BaseElement.Dimension.Width-float64(7*len(ele.Text.Content)))/2,
				Y:       ele.BaseElement.Position.Y + (ele.BaseElement.Dimension.Height+7)/2,
				Strokeable: svg.Strokeable{
					Stroke:      ele.Text.Color,
					StrokeWidth: ele.Stroke.Thickness.Float64(),
				},
			},
		},
	}
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
		Draggable: svg.Draggable{
			Draggable: true,
		},
		Idable: svg.Idable{
			ID: ele.id,
		},
	}
}

type label struct {
	idable
	BaseElement
	Text
	Stroke
	Draggable bool
}

func NewLabel(w, h, x, y float64, content string, id string, draggable bool) *label {
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
		Draggable: draggable,
	}
}

func (ele *label) Svg() svg.SvgElement {
	return &svg.Text{
		Content: ele.Text.Content,
		X:       ele.BaseElement.Position.X,
		Y:       ele.BaseElement.Position.Y,
		Strokeable: svg.Strokeable{
			Stroke:      ele.Text.Color,
			StrokeWidth: ele.Stroke.Thickness.Float64(),
		},
		Draggable: svg.Draggable{
			Draggable: ele.Draggable,
		},
		Idable: svg.Idable{
			ID: ele.id,
		},
	}
}

func (ele *label) MoveTo(x, y float64) {
	ele.Svg().MoveTo(x, y)
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
		Draggable: svg.Draggable{
			Draggable: true,
		},
		Idable: svg.Idable{
			ID: ele.id,
		},
	}
}

type ScaleBox struct {
	idable
	MockupElement
}

func NewScaleBox(mockupElement MockupElement) *ScaleBox {
	return &ScaleBox{
		idable:        idable{id: "M_" + mockupElement.Id()},
		MockupElement: mockupElement,
	}
}

func (ele *ScaleBox) Id() string {
	return ele.idable.id
}

func (ele *ScaleBox) MoveTo(x, y float64) {
	content := ele.Svg().(*svg.Group).Content
	w, h, _, _ := ele.MockupElement.GetWHXY()
	x = x - w/4
	y = y - w/4
	w = w * 1.25
	h = h * 1.25
	square_height := float64(8)
	content[0].MoveTo(x, y)
	content[1].MoveTo(x-square_height/2, y-square_height/2)
	content[2].MoveTo(x+w/2-square_height/2, y-square_height/2)
	content[3].MoveTo(x-square_height/2, y+h/2-square_height/2)
	content[4].MoveTo(x-square_height/2, y+h-square_height/2)
	content[5].MoveTo(x+w-square_height/2, y-square_height/2)
	content[6].MoveTo(x+w/2-square_height/2, y+h-square_height/2)
	content[7].MoveTo(x+w-square_height/2, y+h/2-square_height/2)
	content[8].MoveTo(x+w-square_height/2, y+h-square_height/2)

	//content[9].MoveTo(x+w/2, y+h/2)
}

func (ele *ScaleBox) Svg() svg.SvgElement {
	stroke_width := float64(2)
	square_height := float64(8)
	w, h, x, y := ele.MockupElement.GetWHXY()
	x = x - w/4
	y = y - w/4
	w = w * 1.25
	h = h * 1.25
	return &svg.Group{
		Content: []svg.SvgElement{
			&svg.Path{
				D: []svg.PathItem{
					{Action: svg.MOVETO, Point: svg.NewPoint(x, y)},
					{Action: svg.LINETO, Point: svg.NewPoint(x, y+h)},
					{Action: svg.LINETO, Point: svg.NewPoint(x+w, y+h)},
					{Action: svg.LINETO, Point: svg.NewPoint(x+w, y)},
					{Action: svg.CLOSEPATH},
				},
				Fillable: svg.NewFillable(WHITE, 0),
				Strokeable: svg.Strokeable{
					Stroke:          DARKGREY,
					StrokeDashArray: []float64{square_height},
					StrokeWidth:     stroke_width,
				},
				Idable: svg.Idable{
					ID: ele.idable.id + "_outter",
				},
			},
			scaleboxRect(x-square_height/2, y-square_height/2, stroke_width, square_height, ele.idable.id+"_sq1"),
			scaleboxRect(x+w/2-square_height/2, y-square_height/2, stroke_width, square_height, ele.idable.id+"_sq2"),
			scaleboxRect(x-square_height/2, y+h/2-square_height/2, stroke_width, square_height, ele.idable.id+"_sq3"),
			scaleboxRect(x-square_height/2, y+h-square_height/2, stroke_width, square_height, ele.idable.id+"_sq4"),
			scaleboxRect(x+w-square_height/2, y-square_height/2, stroke_width, square_height, ele.idable.id+"_sq5"),
			scaleboxRect(x+w/2-square_height/2, y+h-square_height/2, stroke_width, square_height, ele.idable.id+"_sq6"),
			scaleboxRect(x+w-square_height/2, y+h/2-square_height/2, stroke_width, square_height, ele.idable.id+"_sq7"),
			scaleboxRect(x+w-square_height/2, y+h-square_height/2, stroke_width, square_height, ele.idable.id+"_sq8"),
			ele.MockupElement.Svg(),
		},
		Draggable: svg.Draggable{
			Draggable: true,
		},
		Fillable: svg.NewFillable(WHITE, 1),
		Idable: svg.Idable{
			ID: ele.idable.id,
		},
	}
}

func scaleboxRect(x, y, stroke_width, square_height float64, id string) *svg.Rect {
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
	}
}
