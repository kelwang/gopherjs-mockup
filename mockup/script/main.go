package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/jquery"
	"github.com/kelwang/gopherjs-mockup/mockup"
	"github.com/kelwang/gopherjs-mockup/mockup/svg"
)

var jQuery = jquery.NewJQuery
var document = js.Global.Get("document")
var editing_class = "editing"
var line_editing_class = "line_editing"

func main() {
	container := initPanel()

	label1 := mockup.NewLabel(180, 20, 400, 158, "big text a lal ha", "E1", svg.EDITABLE)
	textbox1 := mockup.NewTextBox(160, 40, 400, 300, "textbox 1", "E2", svg.EDITABLE)
	button1 := mockup.NewButton(160, 40, 400, 500, "button 1", "E3", svg.EDITABLE)
	box1 := mockup.NewBox(100, 100, 800, 158, "E4", svg.EDITABLE)
	line1 := mockup.NewLine(100, 10, 800, 400, "E5")

	container.Content = append(container.Content,
		label1.Svg(),
		textbox1.Svg(),
		button1.Svg(),
		box1.Svg(),
		line1.Svg(),
	)

	m := map[string]mockup.MockupElement{
		"E1": label1,
		"E2": textbox1,
		"E3": button1,
		"E4": box1,
		"E5": line1,
	}

	container.Content = initToolBar(container, m)

	enableControl(m)
	js.Global.Get("document").Call("write", container.String())
	println("here")
}

func enableControl(m map[string]mockup.MockupElement) {
	jQuery(document).On(jquery.CLICK, svg.EDITABLE.JqSelector(), func(e jquery.Event) {
		wrapEditable(e, m)
	})

	jQuery(document).On(jquery.CLICK, svg.LINABLE.JqSelector(), func(e jquery.Event) {
		wrapLinable(e, m)
	})

	jQuery(document).On(jquery.CLICK, "."+editing_class, func(e jquery.Event) {
		unwrapEditable(e, m)
	})

	jQuery(document).On(jquery.CLICK, "."+line_editing_class, func(e jquery.Event) {
		unwrapLinable(e, m)
	})

	mockup.NewControlEditable(260, 5, 1260, 805).BindEvents(m)
}

func wrapLinable(e jquery.Event, m map[string]mockup.MockupElement) {
	id := jQuery(e.CurrentTarget).Attr("id")
	if mockupE, ok := m[id]; ok {
		border := mockup.NewScaleLine(mockupE)
		m[mockup.EditablePrefix+id] = border
		jQuery("#" + id).ReplaceWith(border.Svg().Jq())
		jQuery("#" + mockup.EditablePrefix + id).AddClass(line_editing_class)
	}
}

func unwrapLinable(e jquery.Event, m map[string]mockup.MockupElement) {
	id := jQuery(e.CurrentTarget).Attr("id")
	if border, ok := m[id]; ok {
		jQuery("#" + id).ReplaceWith(border.(*mockup.ScaleLine).Line().Svg().Jq())
		delete(m, id)
	}
}

func wrapEditable(e jquery.Event, m map[string]mockup.MockupElement) {
	//container := jQuery("svg")
	id := jQuery(e.CurrentTarget).Attr("id")
	if mockupE, ok := m[id]; ok {
		border1 := mockup.NewScaleBox(mockupE)
		m[mockup.EditablePrefix+id] = border1
		jQuery("#" + id).ReplaceWith(border1.Svg().Jq())
		jQuery("#" + mockup.EditablePrefix + id).AddClass(editing_class)
	}
}

func unwrapEditable(e jquery.Event, m map[string]mockup.MockupElement) {

	//container := jQuery("svg")
	id := jQuery(e.CurrentTarget).Attr("id")
	if border, ok := m[id]; ok {
		jQuery("#" + id).ReplaceWith(border.(*mockup.ScaleBox).MockupElement.Svg().Jq())
		delete(m, id)
	}
}

func initToolBar(container svg.Svg, m map[string]mockup.MockupElement) []svg.SvgElement {
	textboxTool := mockup.NewTextBox(60, 20, 30, 20, "textbox", "T1", svg.CLONABLE)
	buttonTool := mockup.NewButton(60, 20, 150, 20, "button", "T2", svg.CLONABLE)
	boxTool := mockup.NewBox(60, 60, 30, 60, "T3", svg.CLONABLE)
	labelTool := mockup.NewLabel(60, 20, 160, 90, "label", "T4", svg.CLONABLE)
	lineTool := mockup.NewLine(60, 0, 30, 160, "T5")

	m["T1"] = textboxTool
	m["T2"] = buttonTool
	m["T3"] = boxTool
	m["T4"] = labelTool
	m["T5"] = lineTool

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
				Width:    250,
				Height:   800,
				X:        5,
				Y:        5,
				Fillable: svg.NewFillable("#F1F1F1", 1),
			},
			//main convas
			&svg.Rect{
				Width:    1000,
				Height:   800,
				X:        260,
				Y:        5,
				Fillable: svg.NewFillable("#F1F1F1", 1),
			},
		},
	}
}
