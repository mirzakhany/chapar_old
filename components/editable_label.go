package components

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// EditableLabel combines a label and an entry widget to allow editing of the label's text.
type EditableLabel struct {
	widget.BaseWidget // Embedding BaseWidget for custom widget functionality.
	label             *widget.Label
	editor            *widget.Entry
	editing           bool
}

// NewEditableLabel creates a new EditableLabel with the specified text.
func NewEditableLabel(text string) *EditableLabel {
	el := &EditableLabel{}
	el.ExtendBaseWidget(el) // Important for custom widget lifecycle.

	el.label = widget.NewLabel(text)
	el.editor = widget.NewEntry()
	el.editor.SetText(text)
	el.editor.Hide()
	el.editor.OnChanged = func(s string) {
		el.label.SetText(s)
	}
	el.editor.OnSubmitted = func(s string) {
		el.editing = false
		el.toggleEditing()
	}

	//el.editor.OnFocusLost = func() {
	//	el.editing = false
	//	el.toggleEditing()
	//}

	return el
}

// Tapped toggles the editing mode on tap events.
func (e *EditableLabel) Tapped(_ *fyne.PointEvent) {
	e.editing = !e.editing
	e.toggleEditing()
}

// toggleEditing switches between label and editor visibility.
func (e *EditableLabel) toggleEditing() {
	if e.editing {
		e.editor.Show()
		e.label.Hide()
	} else {
		e.editor.Hide()
		e.label.Show()
	}
	e.Refresh()
}

// CreateRenderer returns a new renderer for the EditableLabel.
func (e *EditableLabel) CreateRenderer() fyne.WidgetRenderer {
	return &editableLabelRenderer{
		label:  e.label,
		editor: e.editor,
		el:     e,
	}
}

// Refresh is used to update the state of the widget.
func (e *EditableLabel) Refresh() {
	e.BaseWidget.Refresh()
}

// SetText updates the text of the EditableLabel.
func (e *EditableLabel) SetText(text string) {
	e.label.SetText(text)
	e.editor.SetText(text)
}

// editableLabelRenderer defines the rendering logic of the EditableLabel.
type editableLabelRenderer struct {
	label  *widget.Label
	editor *widget.Entry
	el     *EditableLabel
}

func (r *editableLabelRenderer) MinSize() fyne.Size {
	if r.el.editing {
		return r.editor.MinSize()
	}
	return r.label.MinSize()
}

func (r *editableLabelRenderer) Layout(size fyne.Size) {
	r.label.Resize(size)
	r.editor.Resize(size)
}

func (r *editableLabelRenderer) BackgroundColor() color.Color {
	return theme.BackgroundColor()
}

func (r *editableLabelRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.label, r.editor}
}

func (r *editableLabelRenderer) Refresh() {
	r.label.Refresh()
	r.editor.Refresh()
}

func (r *editableLabelRenderer) Destroy() {}
