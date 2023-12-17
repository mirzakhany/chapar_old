package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// RequestListObject is a widget that represents a request in the sidebar
type RequestListObject struct {
	widget.BaseWidget

	Type *CustomLabel
	Name *widget.Label

	hovered    bool
	menuButton *widget.Button
}

// SetType sets the type of the request
func (reqList *RequestListObject) SetType(t string) {
	reqList.Type.SetText(t)
}

// SetName sets the name of the request
func (reqList *RequestListObject) SetName(n string) {
	reqList.Name.SetText(n)
}

// NewRequestListObject creates a new RequestListObject widget
func NewRequestListObject(requestType, name string) *RequestListObject {
	e := widget.NewEntry()
	e.SetText(name)

	item := &RequestListObject{
		Type: NewCustomLabel(requestType, fyne.NewSize(50, 0), fyne.TextAlignTrailing),
		Name: widget.NewLabel(name), // e, //NewEditableLabel(name),
	}

	menu := fyne.NewMenu("",
		fyne.NewMenuItem("Delete", func() {}),
		fyne.NewMenuItem("Duplicate", func() {}),
	)

	item.menuButton = widget.NewButtonWithIcon("", theme.MoreHorizontalIcon(), func() {
		position := fyne.CurrentApp().Driver().AbsolutePositionForObject(item.menuButton)
		position.Y += item.menuButton.Size().Height

		widget.ShowPopUpMenuAtPosition(menu, fyne.CurrentApp().Driver().CanvasForObject(item.menuButton), position)
	})

	item.menuButton.Importance = widget.LowImportance
	item.menuButton.Hide()

	item.ExtendBaseWidget(item)
	return item
}

// MouseIn is called when a desktop pointer enters the widget
func (reqList *RequestListObject) MouseIn(*desktop.MouseEvent) {
	reqList.hovered = true
	reqList.Refresh()
}

// MouseMoved is called when a desktop pointer hovers over the widget
func (reqList *RequestListObject) MouseMoved(*desktop.MouseEvent) {
}

// MouseOut is called when a desktop pointer exits the widget
func (reqList *RequestListObject) MouseOut() {
	reqList.hovered = false
	reqList.Refresh()
}

// Refresh is called when the widget should be redrawn
func (reqList *RequestListObject) Refresh() {
	if reqList.hovered {
		reqList.menuButton.Show()
	} else {
		reqList.menuButton.Hide()
	}
	reqList.BaseWidget.Refresh()
}

// CreateRenderer is a private method to Fyne which links this widget to its renderer
func (reqList *RequestListObject) CreateRenderer() fyne.WidgetRenderer {
	c := container.NewBorder(nil, nil, reqList.Type, reqList.menuButton, reqList.Name)
	return widget.NewSimpleRenderer(c)
}

type CustomLabel struct {
	*widget.Label

	minSize   fyne.Size
	alignment fyne.TextAlign
}

func NewCustomLabel(text string, minSize fyne.Size, alignment fyne.TextAlign) *CustomLabel {
	label := widget.NewLabelWithStyle(text, alignment, fyne.TextStyle{Monospace: false, Bold: true})

	return &CustomLabel{
		Label:   label,
		minSize: minSize,
	}
}

func (l *CustomLabel) MinSize() fyne.Size {
	return l.minSize
}
