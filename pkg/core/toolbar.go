package core

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func NewToolbarLabel(label string) widget.ToolbarItem {
	return &ToolbarLabel{
		Label: label,
	}
}

type ToolbarLabel struct {
	widget.ToolbarAction
	Label string
}

func (t *ToolbarLabel) ToolbarObject() fyne.CanvasObject {
	label := widget.NewLabel(t.Label)
	return label
}
