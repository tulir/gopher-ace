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
TODO direct event bindings
*/

// On binds the given function to the given Ace event.
func (session EditSession) On(on string, f interface{}) {
	session.Call("on", on, f)
}

// AddDynamicMarker adds a dynamic marker to the session.
func (session EditSession) AddDynamicMarker(marker interface{}, inFront bool) *js.Object {
	return session.Call("addDynamicMarker", marker, inFront)
}

// AddGutterDecoration adds `className` to the `row`, to be used for CSS stylings and whatnot.
func (session EditSession) AddGutterDecoration(row int, className string) {
	session.Call("addGutterDecoration", row, className)
}

// AddMarker adds a new marker to the given `Range`.
// If `inFront` is `true`, a front marker is defined, and the `changeFrontMarker` event fires;
// otherwise, the `changeBackMarker` event fires.
func (session EditSession) AddMarker(rangee interface{}, class string, typee interface{}, inFront bool) int {
	return session.Call("addMarker", rangee, class, typee, inFront).Int()
}

// ClearAnnotations clears all the annotations for this session.
// This function also triggers the `changeAnnotation` event.
func (session EditSession) ClearAnnotations() {
	session.Call("clearAnnotations")
}

// ClearBreakpoint removes a breakpoint on the row number given by rows.
// This function also emites the `changeBreakpoint` event.
func (session EditSession) ClearBreakpoint(row int) {
	session.Call("clearBreakpoint", row)
}

// ClearBreakpoints removes all breakpoints on the rows.
// This function also emites the `changeBreakpoint` event.
func (session EditSession) ClearBreakpoints() {
	session.Call("clearBreakpoints")
}

// DocumentToScreenColumn for the given document row and column, returns the screen column.
func (session EditSession) DocumentToScreenColumn(row, docColumn int) int {
	return session.Call("documentToScreenColumn", row, docColumn).Int()
}

// DocumentToScrenPosition converts document coordinates to screen coordinates.
// This takes into account code folding, word wrap, tab size, and any other visual modifications.
func (session EditSession) DocumentToScrenPosition(docRow, docColumn int) *js.Object {
	return session.Call("documentToScrenPosition", docRow, docColumn)
}

// DocumentToScreenRow for the given document row and column, returns the screen row.
func (session EditSession) DocumentToScreenRow(docRow, docColumn int) {
	session.Call("documentToScreenRow")
}

// DuplicateLines duplicates all the text between `firstRow` and `lastRow`.
func (session EditSession) DuplicateLines(firstRow, lastRow int) int {
	return session.Call("duplicateLines", firstRow, lastRow).Int()
}

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
