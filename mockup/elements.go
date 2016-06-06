package mockup

import ()

type Position struct {
	X float64
	Y float64
}

type Dimension struct {
	Width  float64
	Height float64
}

type BaseElement struct {
	Position  Position
	Dimension Dimension
}

type Color string

type Text struct {
	Color Color
	Size  int
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
	Color     Color
	Thickness Thickness
}
