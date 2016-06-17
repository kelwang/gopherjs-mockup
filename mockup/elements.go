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
	GetWHXY() (float64, float64, float64, float64)
}

type BaseElement struct {
	Position  Position
	Dimension Dimension
}

func (be BaseElement) GetWHXY() (float64, float64, float64, float64) {
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

type textBox struct {
	BaseElement
	Text
	Stroke
}

var (
	WHITE    = "white"
	DARKGREY = "#555"
)

func NewTextBox(w, h, x, y float64, content string) *textBox {
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
	}
}

func (ele *textBox) Svg() svg.SvgElement {
	return svg.Group{
		Draggable: svg.Draggable{
			Draggable: true,
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
	BaseElement
	Text
	Stroke
}

func NewButton(w, h, x, y float64, content string) *button {
	return &button{
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
	return svg.Group{
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
	BaseElement
	Stroke
}

func NewBox(w, h, x, y float64) *box {
	return &box{
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
	}
}

type label struct {
	BaseElement
	Text
	Stroke
	Draggable bool
}

func NewLabel(w, h, x, y float64, content string, draggable bool) *label {
	return &label{
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
	}
}

type line struct {
	BaseElement
	Stroke
}

func NewLine(w, h, x, y float64) *line {
	return &line{
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
	}
}

type ScaleBox struct {
	MockupElement
}

func NewScaleBox(mockupElement MockupElement) *ScaleBox {
	return &ScaleBox{
		MockupElement: mockupElement,
	}
}

func (ele *ScaleBox) Move(x, y float64) {

}

func (ele *ScaleBox) Svg() *svg.Group {
	stroke_width := float64(2)
	square_height := float64(8)
	w, h, x, y := ele.MockupElement.GetWHXY()
	x = x - w/4
	y = y - w/4
	w = w * 1.25
	h = h * 1.25
	return &svg.Group{
		Content: []svg.SvgElement{
			svg.Path{
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
			},
			scaleboxRect(x-square_height/2, y-square_height/2, stroke_width, square_height),
			scaleboxRect(x+w/2-square_height/2, y-square_height/2, stroke_width, square_height),
			scaleboxRect(x-square_height/2, y+h/2-square_height/2, stroke_width, square_height),
			scaleboxRect(x-square_height/2, y+h-square_height/2, stroke_width, square_height),
			scaleboxRect(x+w-square_height/2, y-square_height/2, stroke_width, square_height),
			scaleboxRect(x+w/2-square_height/2, y+h-square_height/2, stroke_width, square_height),
			scaleboxRect(x+w-square_height/2, y+h/2-square_height/2, stroke_width, square_height),
			scaleboxRect(x+w-square_height/2, y+h-square_height/2, stroke_width, square_height),
			ele.MockupElement.Svg(),
		},
		Draggable: svg.Draggable{
			Draggable: true,
		},
		Fillable: svg.NewFillable(WHITE, 1),
	}
}

func scaleboxRect(x, y, stroke_width, square_height float64) svg.Rect {
	return svg.Rect{
		X:        x,
		Y:        y,
		Width:    square_height,
		Height:   square_height,
		Fillable: svg.NewFillable(WHITE, 1),
		Strokeable: svg.Strokeable{
			Stroke:      DARKGREY,
			StrokeWidth: stroke_width,
		},
	}
}
