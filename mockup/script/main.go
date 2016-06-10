package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/kelwang/gopherjs-mockup/mockup"
	"github.com/kelwang/gopherjs-mockup/mockup/svg"
)

func main() {
	container := initPanel()

	textboxTool := mockup.NewTextBox(60, 20, 30, 20, "textbox")
	buttonTool := mockup.NewButton(60, 20, 150, 20, "button")
	boxTool := mockup.NewBox(60, 60, 30, 60)
	labelTool := mockup.NewLabel(60, 20, 160, 90, "label")
	lineTool := mockup.NewLine(60, 0, 30, 160)

	container.Content = append(container.Content,
		textboxTool.Svg(),
		buttonTool.Svg(),
		boxTool.Svg(),
		labelTool.Svg(),
		lineTool.Svg(),
	)

	js.Global.Get("document").Call("write", container.String())
	println("here")
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
