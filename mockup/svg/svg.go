package svg

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/jquery"
)

var jQuery = jquery.NewJQuery

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

func mergeAttr(m1, m2 js.M) js.M {
	for k, v := range m2 {
		m1[k] = v
	}
	return m1
}

func initJq(tag string) jquery.JQuery {
	return jQuery(js.Global.Get("document").Call("createElementNS", "http://www.w3.org/2000/svg", tag))
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
	Jq() jquery.JQuery
	MoveTo(x, y float64)
}

//SvgElement
type Rect struct {
	Width  float64 `svg:"width"`
	Height float64 `svg:"height"`
	X      float64 `svg:"x"`
	Y      float64 `svg:"y"`
	RX     float64 `svg:"rx"`
	RY     float64 `svg:"ry"`
	Idable
	Strokeable
	Draggable
	Fillable fillable
}

var unSupportMsg = "Sorry, your browser does not support inline SVG."

func (se *Rect) String() string {
	s := `<rect width="` + jsString(se.Width) + `" height="` + jsString(se.Height) + `" x="` + jsString(se.X) + `" y="` + jsString(se.Y) + `"`
	s += se.Idable.String()
	s += se.Fillable.String()
	s += se.Strokeable.String()
	s += se.Draggable.String()
	if se.RX != 0 {
		s += ` rx="` + jsString(se.RX) + `"`
	}
	if se.RY != 0 {
		s += ` ry="` + jsString(se.RY) + `"`
	}
	s += ` >` + unSupportMsg + `</rect>`
	return s
}

func (se *Rect) Jq() jquery.JQuery {
	attr := js.M{
		"width":  se.Width,
		"height": se.Height,
		"x":      se.X,
		"y":      se.Y,
	}
	attr = mergeAttr(attr, se.Idable.Attr())
	attr = mergeAttr(attr, se.Fillable.Attr())
	attr = mergeAttr(attr, se.Strokeable.Attr())
	attr = mergeAttr(attr, se.Draggable.Attr())
	if se.RX != 0 {
		attr["rx"] = se.RX
	}
	if se.RY != 0 {
		attr["ry"] = se.RY
	}
	return initJq("rect").SetAttr(attr)
}

func (se *Rect) MoveTo(x, y float64) {
	se.X = x
	se.Y = y
	jQuery("#" + se.ID).SetAttr(js.M{
		"x": se.X,
		"y": se.Y,
	})
}

type StrokeLineCap int

const (
	BUTT StrokeLineCap = iota
	ROUND
	SQUARE
)

var strokeLineCapString = []string{"butt", "round", "square"}

type Strokeable struct {
	StrokeWidth     float64       `svg:"stroke-width"`
	Stroke          string        `svg:"stroke"`
	StrokeDashArray []float64     `svg:"stroke-dasharray"`
	StrokeLineCap   StrokeLineCap `svg:"stroke-linecap"`
}

func (se Strokeable) String() string {
	s := ""
	if se.Stroke != "" {
		s += ` stroke="` + jsString(se.Stroke) + `"`
	}
	if se.StrokeWidth != 0 {
		s += ` stroke-width="` + jsString(se.StrokeWidth) + `"`
	}

	if se.StrokeDashArray != nil && len(se.StrokeDashArray) > 0 {
		s += ` stroke-dasharray="`
		for k, v := range se.StrokeDashArray {
			if k != 0 {
				s += ","
			}
			s += jsString(v)
		}
		s += `"`
	}

	if se.StrokeLineCap != BUTT {
		s += ` stroke-linecap="` + strokeLineCapString[se.StrokeLineCap] + `"`
	}

	return s
}

func (se Strokeable) Attr() js.M {
	attr := js.M{}
	if se.Stroke != "" {
		attr["stroke"] = se.Stroke
	}
	if se.StrokeWidth != 0 {
		attr["stroke-width"] = se.StrokeWidth
	}

	if se.StrokeDashArray != nil && len(se.StrokeDashArray) > 0 {
		s := `"`
		for k, v := range se.StrokeDashArray {
			if k != 0 {
				s += ","
			}
			s += jsString(v)
		}
		s += `"`
		attr["stroke-dasharray"] = s
	}

	if se.StrokeLineCap != BUTT {
		attr["stroke-linecap"] = strokeLineCapString[se.StrokeLineCap]
	}
	return attr
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

func (draggable Draggable) Attr() js.M {
	if draggable.Draggable {
		return js.M{"class": "draggable"}
	}
	return js.M{}
}

type Idable struct {
	ID string
}

func (id Idable) String() string {
	if id.ID == "" {
		return ""
	}
	return ` id="` + jsString(id.ID) + `"`
}

func (id Idable) Attr() js.M {
	if id.ID == "" {
		return js.M{}
	}

	return js.M{
		"id": id.ID,
	}
}

type fillable struct {
	Fill    string  `svg:"fill"`
	Opacity float64 `svg:"fill-opacity"`
}

func NewFillable(fill string, opacity float64) fillable {
	return fillable{
		Fill:    fill,
		Opacity: opacity,
	}
}

func (se fillable) String() string {
	s := ""

	if se.Fill != "" {
		s += ` fill="` + jsString(se.Fill) + `"`
	}

	if se.Opacity != float64(1) {
		s += ` fill-opacity="` + jsString(se.Opacity) + `"`
	}

	return s
}

func (se fillable) Attr() js.M {
	attr := js.M{}
	if se.Fill != "" {
		attr["fill"] = se.Fill
	}
	if se.Opacity != float64(1) {
		attr["fill-opacity"] = se.Opacity
	}
	return attr
}

//SvgElement
type Circle struct {
	X        float64 `svg:"cx"`
	Y        float64 `svg:"cy"`
	R        float64 `svg:"r"`
	Fillable fillable
	Strokeable
	Draggable
}

func (se Circle) String() string {
	s := `<circle r="` + jsString(se.R) + `" cx="` + jsString(se.X) + `" cy="` + jsString(se.Y) + `"`
	s += se.Fillable.String()
	s += se.Strokeable.String()
	s += se.Draggable.String()
	s += ` >` + unSupportMsg + `</circle>`
	return s
}

//SvgElement
type Ellipse struct {
	X        float64 `svg:"cx"`
	Y        float64 `svg:"cy"`
	RX       float64 `svg:"rx"`
	RY       float64 `svg:"ry"`
	Fillable fillable
	Strokeable
	Draggable
}

func (se Ellipse) String() string {
	s := `<ellipse rx="` + jsString(se.RX) + `" ry="` + jsString(se.RY) + `" cx="` + jsString(se.X) + `" cy="` + jsString(se.Y) + `"`
	s += se.Fillable.String()
	s += se.Strokeable.String()
	s += se.Draggable.String()
	s += ` >` + unSupportMsg + `</ellipse>`
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
	Idable
}

func (se *Line) String() string {
	s := `<line x1="` + jsString(se.X1) + `" y1="` + jsString(se.Y1) + `" x2="` + jsString(se.X2) + `" y2="` + jsString(se.Y2) + `"`
	s += se.Idable.String()
	s += se.Strokeable.String()
	s += se.Draggable.String()
	s += ` >` + unSupportMsg + `</line>`
	return s
}

func (se *Line) Jq() jquery.JQuery {
	attr := js.M{
		"x1": se.X1,
		"y1": se.Y1,
		"x2": se.X2,
		"y2": se.Y2,
	}
	attr = mergeAttr(attr, se.Idable.Attr())
	attr = mergeAttr(attr, se.Strokeable.Attr())
	attr = mergeAttr(attr, se.Draggable.Attr())
	return initJq("line").SetAttr(attr)
}

func (se *Line) MoveTo(x, y float64) {
	dx := x - se.X1
	dy := y - se.Y1
	se.X1 = x
	se.Y1 = y
	se.X2 += dx
	se.Y2 += dy
}

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
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
	Points   Points `svg:"points"`
	Fillable fillable
	Strokeable
	Draggable
}

func (se Polygon) String() string {
	s := `<polygon points="` + se.Points.String() + `"`
	s += se.Fillable.String()
	s += se.Strokeable.String()
	s += se.Draggable.String()
	s += ` >` + unSupportMsg + `</polygon>`
	return s
}

//SvgElement
type Polyline struct {
	Points   Points `svg:"points"`
	Fillable fillable
	Strokeable
	Draggable
}

func (se Polyline) String() string {
	s := `<polyline points="` + se.Points.String() + `"`
	s += se.Fillable.String()
	s += se.Strokeable.String()
	s += se.Draggable.String()
	s += ` >` + unSupportMsg + `</polyline>`
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
	D        PathItems `svg:"d"`
	Fillable fillable
	Strokeable
	Draggable
	Idable
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

func (se *Path) String() string {
	s := `<path d="` + se.D.String() + `"`
	s += se.Idable.String()
	s += se.Fillable.String()
	s += se.Strokeable.String()
	s += se.Draggable.String()
	s += ` >` + unSupportMsg + `</path>`
	return s
}

func (se *Path) Jq() jquery.JQuery {
	attr := js.M{
		"d": se.D.String(),
	}
	attr = mergeAttr(attr, se.Fillable.Attr())
	attr = mergeAttr(attr, se.Strokeable.Attr())
	attr = mergeAttr(attr, se.Draggable.Attr())

	return initJq("path").SetAttr(attr)
}

func (se *Path) MoveTo(x, y float64) {
	dx := x - se.D[0].Point.x
	dy := y - se.D[0].Point.y
	for k := range se.D {
		se.D[k].Point.x += dx
		se.D[k].Point.y += dy
	}
	jQuery("#"+se.ID).SetAttr("d", se.D.String())
}

type Text struct {
	Content  string  `svg:"content"`
	X        float64 `svg:"x"`
	Y        float64 `svg:"y"`
	Fillable fillable
	Strokeable
	Draggable
	Idable
}

func (se *Text) String() string {
	s := `<text x="` + jsString(se.X) + `" y="` + jsString(se.Y) + `"`
	s += se.Idable.String()
	s += se.Fillable.String()
	s += se.Strokeable.String()
	s += se.Draggable.String()
	s += ` >` + se.Content + `</text>`
	return s
}

func (se *Text) Jq() jquery.JQuery {
	attr := js.M{
		"x": se.X,
		"y": se.Y,
	}
	attr = mergeAttr(attr, se.Idable.Attr())
	attr = mergeAttr(attr, se.Fillable.Attr())
	attr = mergeAttr(attr, se.Strokeable.Attr())
	attr = mergeAttr(attr, se.Draggable.Attr())
	return initJq("text").SetAttr(attr).SetHtml(se.Content)
}

func (se *Text) MoveTo(x, y float64) {
	se.X = x
	se.Y = y
	jQuery("#" + se.ID).SetAttr(js.M{
		"x": se.X,
		"y": se.Y,
	})

}

type Group struct {
	Content []SvgElement `svg:"content"`
	Idable
	Fillable fillable
	Strokeable
	Draggable
}

func (se *Group) String() string {
	s := `<g `
	s += se.Idable.String()
	s += se.Fillable.String()
	s += se.Strokeable.String()
	s += se.Draggable.String()
	s += ` >`
	for _, v := range se.Content {
		s += v.String()
	}
	s += `</g>`
	return s
}

func (se *Group) Jq() jquery.JQuery {
	attr := js.M{}
	attr = mergeAttr(attr, se.Idable.Attr())
	attr = mergeAttr(attr, se.Fillable.Attr())
	attr = mergeAttr(attr, se.Strokeable.Attr())
	attr = mergeAttr(attr, se.Draggable.Attr())
	s := ""
	for _, v := range se.Content {
		s += v.String()
	}
	return initJq("g").SetAttr(attr).SetHtml(s)
}

func (se *Group) MoveTo(x, y float64) {
	//do nothing
	//for k := range se.Content {
	//	se.Content[k].MoveTo(x, y)
	//}
}
