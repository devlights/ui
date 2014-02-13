// 12 february 2014
package main

import (
	"fmt"
	"sync"
)

// A Button represents a clickable button with some text.
type Button struct {
	// This channel gets a message when the button is clicked. Unlike other channels in this package, this channel is initialized to non-nil when creating a new button, and cannot be set to nil later.
	Clicked	chan struct{}

	lock		sync.Mutex
	created	bool
	parent	Control
	sysData	*sysData
	initText	string
}

// NewButton creates a new button with the specified text.
func NewButton(text string) (b *Button) {
	return &Button{
		sysData:	&sysData{
			cSysData:		cSysData{
				ctype:	c_button,
			},
		},
		initText:	text,
		Clicked:	make(chan struct{}),
	}
}

// SetText sets the button's text.
func (b *Button) SetText(text string) (err error) {
	b.lock.Lock()
	defer b.lock.Unlock()

	if b.pWin != nil && b.pWin.created {
		panic("TODO")
	}
	b.initText = text
	return nil
}

func (b *Button) apply(window *sysData) error {
	b.lock.Lock()
	defer b.lock.Unlock()

	b.sysData.clicked = b.Clicked
	return b.sysData.make(b.initText, 300, 300, window)
	// TODO size to parent size
}

func (b *Button) setParent(c Control) {
	b.parent = c
}
