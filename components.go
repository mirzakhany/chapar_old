package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// RequestListObject is a widget that represents a request in the sidebar
type RequestListObject struct {
	widget.BaseWidget

	Type *canvas.Text
	Name *EditableLabel
}

// SetType sets the type of the request
func (reqList *RequestListObject) SetType(t string) {
	reqList.Type.Text = t
}

// SetName sets the name of the request
func (reqList *RequestListObject) SetName(n string) {
	reqList.Name.SetText(n)
}

// NewRequestListObject creates a new RequestListObject widget
func NewRequestListObject(requestType, name string) *RequestListObject {
	e := widget.NewEntry()
	e.SetText(name)

	t := canvas.NewText(requestType, theme.ButtonColor())
	t.TextStyle.Monospace = true
	t.TextStyle.Bold = true
	t.Alignment = fyne.TextAlignTrailing

	item := &RequestListObject{
		Type: t,
		Name: NewEditableLabel(name),
	}

	item.ExtendBaseWidget(item)

	return item
}

// CreateRenderer is a private method to Fyne which links this widget to its renderer
func (reqList *RequestListObject) CreateRenderer() fyne.WidgetRenderer {
	c := container.NewBorder(nil, nil, reqList.Type, nil, reqList.Name)
	return widget.NewSimpleRenderer(c)
}

type EditableLabel struct {
	*widget.Label
	Entry *widget.Entry

	editing bool
}

func NewEditableLabel(text string) *EditableLabel {
	label := widget.NewLabel(text)
	entry := widget.NewEntry()
	entry.SetText(text)

	entry.OnChanged = func(s string) {
		label.SetText(s)
	}

	editableLabel := &EditableLabel{
		Label: label,
		Entry: entry,
	}

	editableLabel.ExtendBaseWidget(editableLabel)
	return editableLabel
}

func (editLabel *EditableLabel) Tapped(_ *fyne.PointEvent) {
	fmt.Println("Tapped")
	editLabel.editing = true
	editLabel.Refresh()
}

func (editLabel *EditableLabel) TypedRune(r rune) {
	if r == '\r' {
		editLabel.editing = false
		editLabel.Refresh()
	}
}

// CreateRenderer is a private method to Fyne which links this widget to its renderer
func (editLabel *EditableLabel) CreateRenderer() fyne.WidgetRenderer {
	if editLabel.editing {
		return widget.NewSimpleRenderer(editLabel.Entry)
	}

	return widget.NewSimpleRenderer(editLabel.Label)
}
