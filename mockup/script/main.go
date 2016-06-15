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

	label1 := mockup.NewLabel(60, 60, 400, 150, "label")
	container.Content = append(container.Content, label1.Svg())

	enableControl()
	js.Global.Get("document").Call("write", container.String())
	println("here")
}

func enableControl() {
	println("control")
	movable := false
	jQuery(js.Global.Get("document")).On(jquery.MOUSEDOWN, ".draggable", func(e jquery.Event) {
		movable = true
		current := jQuery(e.CurrentTarget)
		current.SetCss("cursor", "move")
		jQuery(js.Global.Get("document")).On(jquery.MOUSEMOVE, func(e jquery.Event) {
			if movable {
				current.SetAttr("x", e.Get("clientX").Float())
				current.SetAttr("y", e.Get("clientY").Float())
			}
		})
	})
	jQuery(js.Global.Get("document")).On(jquery.MOUSEUP, ".draggable", func(e jquery.Event) {
		movable = false
		jQuery(e.CurrentTarget).SetCss("cursor", "auto")

	})
}

func stopDrag(e jquery.Event) {

}

func initToolBar(container svg.Svg) []svg.SvgElement {
	textboxTool := mockup.NewTextBox(60, 20, 30, 20, "textbox")
	buttonTool := mockup.NewButton(60, 20, 150, 20, "button")
	boxTool := mockup.NewBox(60, 60, 30, 60)
	labelTool := mockup.NewLabel(60, 20, 160, 90, "label")
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
