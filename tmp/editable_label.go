package tmp

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type EditableLabel struct {
	label   *widget.Label
	editor  *widget.Entry
	editing bool
}

func NewEditableLabel(text string) *EditableLabel {
	label := &EditableLabel{
		label:   widget.NewLabel(text),
		editing: false,
		editor:  widget.NewEntry(),
	}
	return label
}

func (e *EditableLabel) Move(position fyne.Position) {
	e.label.Move(position)
	e.editor.Move(position)
}

func (e *EditableLabel) Position() fyne.Position {
	return e.label.Position()
}

func (e *EditableLabel) Size() fyne.Size {
	return e.label.MinSize()
}

func (e *EditableLabel) Hide() {
	e.label.Hide()
	e.editor.Hide()
}

func (e *EditableLabel) Visible() bool {
	return e.label.Visible() || e.editor.Visible()
}

func (e *EditableLabel) Show() {
	e.label.Show()
}

func (e *EditableLabel) Tapped(_ *fyne.PointEvent) {
	if !e.editing {
		e.editing = true
	}
	e.toggleEditing()
}

func (e *EditableLabel) toggleEditing() {
	if e.editing {
		e.editor.SetText(e.label.Text)
		e.editor.Show()
		e.label.Hide()
	} else {
		e.editor.Hide()
		e.label.Show()
	}
	e.Refresh()
}

func (e *EditableLabel) CreateRenderer() fyne.WidgetRenderer {
	e.editor = widget.NewEntry()
	e.editor.Hide()
	e.editor.OnChanged = func(s string) {
		e.label.Text = s
	}

	e.editor.OnSubmitted = func(s string) {
		e.label.Text = s
		e.editing = false
		e.toggleEditing()
	}

	e.label.ExtendBaseWidget(e)
	return widget.NewSimpleRenderer(container.NewStack(e.label, e.editor))
}

func (e *EditableLabel) MinSize() fyne.Size {
	if e.editing {
		return e.editor.MinSize()
	}
	return e.label.MinSize()
}

func (e *EditableLabel) Refresh() {
	e.label.Refresh()
	e.editor.Refresh()
}

func (e *EditableLabel) Resize(s fyne.Size) {
	e.label.Resize(s)
	e.editor.Resize(s)
}

func (e *EditableLabel) SetText(text string) {
	e.label.SetText(text)
	e.editor.SetText(text)
}
