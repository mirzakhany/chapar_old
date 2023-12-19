package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

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
