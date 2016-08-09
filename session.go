package ace

import (
	"github.com/gopherjs/gopherjs/js"
)

// An EditSession is a wrapper for the Ace EditSession type.
type EditSession struct {
	*js.Object
}

// GetSession returns the EditSession object for the Editor
func (edit Editor) GetSession() EditSession {
	return EditSession{edit.Call("getSession")}
}

/*
TODO most functions missing
*/

// SetTabSize sets the number of spaces that define a soft tab; for example,
// passing in 4 transforms the soft tabs to be equivalent to four spaces.
// This function also emits the `changeTabSize` event.
func (session EditSession) SetTabSize(size int) {
	session.Call("SetTabSize", 4)
}

// GetLength returns the number of rows in the document
func (session EditSession) GetLength() int {
	return session.Call("getLength").Int()
}

// SetUseSoftTabs - pass `true` to enable the use of soft tabs.
// Soft tabs means you're using spaces instead of the tab character (`\t`)
func (session EditSession) SetUseSoftTabs(useSoftTabs bool) {
	session.Call("setUseSoftTabs", useSoftTabs)
}

// SetUseWrapMode sets whether or not line wrapping is enabled.
// If `useWrapMode` is different than the current value, the `changeWrapMode` event is emitted.
func (session EditSession) SetUseWrapMode(useWrapMode bool) {
	session.Call("setUseWrapMode", useWrapMode)
}
