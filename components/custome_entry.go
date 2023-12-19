package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type CustomEntry struct {
	*widget.Entry

	OnFocused   func()
	OnFocusLost func()
}

func NewCustomEntry() *CustomEntry {
	return &CustomEntry{
		Entry: widget.NewEntry(),
	}
}

func (e *CustomEntry) FocusGained() {
	if e.OnFocused != nil {
		e.OnFocused()
	}
}

func (e *CustomEntry) FocusLost() {
	if e.OnFocusLost != nil {
		e.OnFocusLost()
	}
}

func (e *CustomEntry) TypedKey(key *fyne.KeyEvent) {
	e.Entry.TypedKey(key)
}

func (e *CustomEntry) TypedRune(r rune) {
	e.Entry.TypedRune(r)
}
