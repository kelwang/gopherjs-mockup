package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/jquery"
	"github.com/kelwang/gopherjs-mockup/mockup"
	"github.com/kelwang/gopherjs-mockup/mockup/svg"
)

var jQuery = jquery.NewJQuery

func main() {
	container := initPanel()
	container.Content = initToolBar(container)

	label1 := mockup.NewLabel(60, 60, 400, 150, "big text a lal ha", true)
	label2 := mockup.NewLabel(60, 60, 400, 450, "label2", true)
	container.Content = append(container.Content,
		label1.Svg(),
		label2.Svg())

	enableControl()
	js.Global.Get("document").Call("write", container.String())
	println("here")
}

func enableControl() {
	mockup.NewDraggable(jQuery(nil), 260, 5, 1260, 805).BindEvents()
}

func initToolBar(container svg.Svg) []svg.SvgElement {
	textboxTool := mockup.NewTextBox(60, 20, 30, 20, "textbox")
	buttonTool := mockup.NewButton(60, 20, 150, 20, "button")
	boxTool := mockup.NewBox(60, 60, 30, 60)
	labelTool := mockup.NewLabel(60, 20, 160, 90, "label", false)
	lineTool := mockup.NewLine(60, 0, 30, 160)
	return append(container.Content,
		textboxTool.Svg(),
		buttonTool.Svg(),
		boxTool.Svg(),
		labelTool.Svg(),
		lineTool.Svg(),
	)
}

func initPanel() svg.Svg {
	return svg.Svg{
		Width:  1300,
		Height: 800,
		Content: []svg.SvgElement{
			//left toolbar
			&svg.Rect{
				Width:  250,
				Height: 800,
				X:      5,
				Y:      5,
				Fillable: svg.Fillable{
					Fill: "#F1F1F1",
				},
			},
			//main convas
			&svg.Rect{
				Width:  1000,
				Height: 800,
				X:      260,
				Y:      5,
				Fillable: svg.Fillable{
					Fill: "#F1F1F1",
				},
			},
		},
	}
}
