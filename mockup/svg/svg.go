package svg

import (
	"github.com/gopherjs/gopherjs/js"
)

type Svg struct {
	Width   float64
	Height  float64
	Content []SvgElement
}

func jsString(a interface{}) string {
	if a != nil {
		s := js.Global.Call("String", a).String()
		if s != "undefined" && s != "null" {
			return s
		}
	}
	return ""
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

//SvgElement
type Rect struct {
	Width  float64 `svg:"width"`
	Height float64 `svg:"height"`
	X      float64 `svg:"x"`
	Y      float64 `svg:"y"`
	RX     float64 `svg:"rx"`
	RY     float64 `svg:"ry"`
	Strokeable
	Draggable
	Fillable
}

var unSupportMsg = "Sorry, your browser does not support inline SVG."

func (se *Rect) String() string {
	s := `<rect width="` + jsString(se.Width) + `" height="` + jsString(se.Height) + `" x="` + jsString(se.X) + `" y="` + jsString(se.Y) + `"`
	s += se.Fillable.String()
	s += se.Strokeable.String()
	s += se.Draggable.String()
	if se.RX != 0 {
		s += ` rx="` + jsString(se.RX) + `"`
	}
	if se.RY != 0 {
		s += ` ry="` + jsString(se.RY) + `"`
	}
	s += `" >` + unSupportMsg + `</rect>`
	return s
}

type Strokeable struct {
	StrokeWidth float64 `svg:"stroke-width"`
	Stroke      string  `svg:"stroke"`
}

func (se Strokeable) String() string {
	s := ""
	if se.Stroke != "" {
		s += ` stroke="` + jsString(se.Stroke) + `"`
	}
	if se.StrokeWidth != 0 {
		s += ` stroke-width="` + jsString(se.StrokeWidth) + `"`
	}
	return s
}

type Draggable struct {
	Draggable bool
}

func (draggable Draggable) String() string {
	if draggable.Draggable {
		return ` class="draggable"`
	}
	return ""
}

type Fillable struct {
	Fill string `svg:"fill"`
}

func (se Fillable) String() string {
	s := ""
	if se.Fill != "" {
		s += ` fill="` + jsString(se.Fill) + `"`
	}
	return s
}

//SvgElement
type Circle struct {
	X float64 `svg:"cx"`
	Y float64 `svg:"cy"`
	R float64 `svg:"r"`
	Fillable
	Strokeable
	Draggable
}

func (se Circle) String() string {
	s := `<circle r="` + jsString(se.R) + `" cx="` + jsString(se.X) + `" cy="` + jsString(se.Y) + `"`
	s += se.Fillable.String()
	s += se.Strokeable.String()
	s += se.Draggable.String()
	s += `" >` + unSupportMsg + `</circle>`
	return s
}

//SvgElement
type Ellipse struct {
	X  float64 `svg:"cx"`
	Y  float64 `svg:"cy"`
	RX float64 `svg:"rx"`
	RY float64 `svg:"ry"`
	Fillable
	Strokeable
	Draggable
}

func (se Ellipse) String() string {
	s := `<ellipse rx="` + jsString(se.RX) + `" ry="` + jsString(se.RY) + `" cx="` + jsString(se.X) + `" cy="` + jsString(se.Y) + `"`
	s += se.Fillable.String()
	s += se.Strokeable.String()
	s += se.Draggable.String()
	s += `" >` + unSupportMsg + `</ellipse>`
	return s
}

//SvgElement
type Line struct {
	X1 float64 `svg:"x1"`
	Y1 float64 `svg:"y1"`
	X2 float64 `svg:"x2"`
	Y2 float64 `svg:"y2"`
	Strokeable
	Draggable
}

func (se Line) String() string {
	s := `<line x1="` + jsString(se.X1) + `" y1="` + jsString(se.Y1) + `" x2="` + jsString(se.X2) + `" y2="` + jsString(se.Y2) + `"`
	s += se.Strokeable.String()
	s += se.Draggable.String()
	s += `" >` + unSupportMsg + `</line>`
	return s
}

type Point struct {
	x float64
	y float64
}

type Points []Point

func (p Points) String() string {
	s := ""
	for i := 0; i < len(p); i++ {
		s += jsString(p[i].x) + "," + jsString(p[i].y)
		if i != len(p)-1 {
			s += " "
		}
	}
	return s
}

//SvgElement
type Polygon struct {
	Points Points `svg:"points"`
	Fillable
	Strokeable
	Draggable
}

func (se Polygon) String() string {
	s := `<polygon points="` + se.Points.String() + `"`
	s += se.Fillable.String()
	s += se.Strokeable.String()
	s += se.Draggable.String()
	s += `" >` + unSupportMsg + `</polygon>`
	return s
}

//SvgElement
type Polyline struct {
	Points Points `svg:"points"`
	Fillable
	Strokeable
	Draggable
}

func (se Polyline) String() string {
	s := `<polyline points="` + se.Points.String() + `"`
	s += se.Fillable.String()
	s += se.Strokeable.String()
	s += se.Draggable.String()
	s += `" >` + unSupportMsg + `</polyline>`
	return s
}

type pathAction int

const (
	MOVETO pathAction = iota
	LINETO
	HORIZONTAL_LINETO
	VERTICAL_LINETO
	CURVETO
	SMOOTH_CURVETO
	QUADRATIC_BEZIER_CURVE
	SMOOTH_QUADRATIC_BEZIER_CURVETO
	ELLIPTICAL_ARC
	CLOSEPATH
)

var pathActionString = "MLHVCSQTAZ"

func (pa pathAction) String() string {
	return pathActionString[pa : pa+1]
}

type PathItem struct {
	Action pathAction
	Point  Point
}

type PathItems []PathItem

//SvgElement
type Path struct {
	D PathItems `svg:"d"`
	Fillable
	Strokeable
	Draggable
}

func (ps PathItems) String() string {
	s := ""
	for i := 0; i < len(ps); i++ {
		s += ps[i].Action.String()
		if ps[i].Action != CLOSEPATH {
			s += " " + jsString(ps[i].Point.x) + " " + jsString(ps[i].Point.y)
		}
		if i != len(ps)-1 {
			s += " "
		}
	}
	return s
}

func (se Path) String() string {
	s := `<path d="` + se.D.String() + `"`
	s += se.Fillable.String()
	s += se.Strokeable.String()
	s += se.Draggable.String()
	s += `" >` + unSupportMsg + `</path>`
	return s
}

type Text struct {
	Content string  `svg:"content"`
	X       float64 `svg:"x"`
	Y       float64 `svg:"y"`
	Fillable
	Strokeable
	Draggable
}

func (se Text) String() string {
	s := `<text x="` + jsString(se.X) + `" y="` + jsString(se.Y) + `"`
	s += se.Fillable.String()
	s += se.Strokeable.String()
	s += se.Draggable.String()
	s += `" >` + se.Content + `</text>`
	return s
}

type Group struct {
	Content []SvgElement `svg:"content"`
	Fillable
	Strokeable
	Draggable
}

func (se Group) String() string {
	s := `<g `
	s += se.Fillable.String()
	s += se.Strokeable.String()
	s += se.Draggable.String()
	s += `" >`
	for _, v := range se.Content {
		s += v.String()
	}
	s += `</g>`
	return s
}