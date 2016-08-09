package ace

import (
	"github.com/gopherjs/gopherjs/js"
)

// An Editor is a wrapper for the Ace Editor type.
type Editor struct {
	*js.Object
}

// OnBlur binds the given function to the Ace blur event.
func (edit Editor) OnBlur(f func()) {
	edit.Call("on", "blur", f)
}

// OnChange binds the given function to the Ace change event.
func (edit Editor) OnChange(f func(e *js.Object)) {
	edit.Call("on", "change", f)
}

// OnChangeSelectionStyle binds the given function to the Ace changeSelectionStyle event.
func (edit Editor) OnChangeSelectionStyle(f func(data *js.Object)) {
	edit.Call("on", "changeSelectionStyle", f)
}

// OnChangeSession binds the given function to the Ace changeSession event.
func (edit Editor) OnChangeSession(f func(e *js.Object)) {
	edit.Call("on", "changeSession", f)
}

// OnCopy binds the given function to the Ace copy event.
func (edit Editor) OnCopy(f func(text string)) {
	edit.Call("on", "copy", func(dat *js.Object) {
		f(dat.String())
	})
}

// OnFocus binds the given function to the Ace focus event.
func (edit Editor) OnFocus(f func()) {
	edit.Call("on", "focus", f)
}

// OnPaste binds the given function to the Ace paste event.
func (edit Editor) OnPaste(f func(e *js.Object)) {
	edit.Call("on", "paste", f)
}

// On binds the given function to the given Ace event.
func (edit Editor) On(on string, f interface{}) {
	edit.Call("on", on, f)
}

// AddSelectionMarker adds the selection and cursor.
func (edit Editor) AddSelectionMarker(obj interface{}) *js.Object {
	return edit.Call("addSelectionMarker", obj)
}

// AlignCursors aligns the cursors or selected text.
func (edit Editor) AlignCursors() {
	edit.Call("alignCursors")
}

// BlockOutdent outdents the current line.
func (edit Editor) BlockOutdent() {
	edit.Call("blockOutdent")
}

// Blur blurs the current `textInput`
func (edit Editor) Blur() {
	edit.Call("blur")
}

// CenterSelection attempts to center the current selection on the screen.
func (edit Editor) CenterSelection() {
	edit.Call("centerSelection")
}

// ClearSelection empties the selection (by de-selecting it).
// This function also emits the `change-selection` event
func (edit Editor) ClearSelection() {
	edit.Call("clearSelection")
}

// CopyLinesDown copies all the selected lines down one row.
func (edit Editor) CopyLinesDown() int {
	return edit.Call("copyLinesDown").Int()
}

// CopyLinesUp copies all the selected lines up one row.
func (edit Editor) CopyLinesUp() int {
	return edit.Call("copyLinesUp").Int()
}

// Destroy cleans up the entire editor.
func (edit Editor) Destroy() {
	edit.Call("destroy")
}

// DuplicateSelection is an undocumented Ace function.
func (edit Editor) DuplicateSelection(args ...interface{}) *js.Object {
	return edit.Call("duplicateSelection", args...)
}

// ExecCommand is an undocumented Ace function.
func (edit Editor) ExecCommand(args ...interface{}) *js.Object {
	return edit.Call("execCommand", args...)
}

// ExitMultiSelectMode removes all selections except the last added one.
func (edit Editor) ExitMultiSelectMode() {
	edit.Call("exitMultiSelectMode")
}

// Find attempts to find `needle` within the document.
// For more information on `options`, see `Search`
func (edit Editor) Find(needle string, options interface{}, animate bool) {
	edit.Call("find", needle, options, animate)
}

// FindAll finds and selects all the occurences of `needle`
func (edit Editor) FindAll(needle string, the interface{}, keeps bool) int {
	return edit.Call("findAll", needle, the, keeps).Int()
}

// FindNext performs another search for `needle` in the document.
// For more information on `options`, see `Search`.
func (edit Editor) FindNext(options interface{}, animate bool) {
	edit.Call("findNext", options, animate)
}

// FindPrevious performs a search for `needle` backwards.
// For more information on `options`, see `Search`.
func (edit Editor) FindPrevious(options interface{}, animate bool) {
	edit.Call("findPrevious", options, animate)
}

// Focus brings the current `textInput` into focus.
func (edit Editor) Focus() {
	edit.Call("focus")
}

// ForEachSelection executes a command for each selection range
func (edit Editor) ForEachSelection(cmd, args string) {
	edit.Call("forEachSelection", cmd, args)
}

// GetAnimatedScroll is an undocumented Ace function.
func (edit Editor) GetAnimatedScroll(args ...interface{}) *js.Object {
	return edit.Call("getAnimatedScroll", args...)
}

// GetBehavioursEnabled returns `true` if the behaviors are currently enabled.
// "Behaviors" in this case is the auto-pairing of special characters,
// like quotation marks, parenthesis, or brackets.
func (edit Editor) GetBehavioursEnabled() bool {
	return edit.Call("getBehavioursEnabled").Bool()
}

// GetCopyText returns the string of text currently highlighted.
func (edit Editor) GetCopyText() string {
	return edit.Call("getCopyText").String()
}

// GetCursorPosition get sthe current position of the cursor
func (edit Editor) GetCursorPosition() *js.Object {
	return edit.Call("getCursorPosition")
}

// GetCursorPositionScreen returns the screen position of the cursor
func (edit Editor) GetCursorPositionScreen() int {
	return edit.Call("getCursorPositionScreen").Int()
}

// GetDisplayIndentGuides is an undocumented Ace function.
func (edit Editor) GetDisplayIndentGuides(args ...interface{}) *js.Object {
	return edit.Call("getDisplayIndentGuides", args...)
}

// GetDragDelay returns the current mouse drag delay.
func (edit Editor) GetDragDelay() int {
	return edit.Call("getDragDelay").Int()
}

// GetFadeFoldWidgets is an undocumented Ace function.
func (edit Editor) GetFadeFoldWidgets(args ...interface{}) *js.Object {
	return edit.Call("getFadeFoldWidgets", args...)
}

// GetFirstVisibleRow returns the index of the first visible row.
func (edit Editor) GetFirstVisibleRow() int {
	return edit.Call("getFirstVisibleRow").Int()
}

// GetHighlightActiveLine returns `true` if the current lines are always highlighted.
func (edit Editor) GetHighlightActiveLine() bool {
	return edit.Call("getHighlightActiveLine").Bool()
}

// GetHighlightGutterLine is an undocumented Ace function.
func (edit Editor) GetHighlightGutterLine(args ...interface{}) *js.Object {
	return edit.Call("getHighlightGutterLine", args...)
}

// GetHighlightSelectedWord returns `true` if currently highlighted words are to be highlighted.
func (edit Editor) GetHighlightSelectedWord() bool {
	return edit.Call("getHighlightSelectedWord").Bool()
}

// GetKeyboardHandler returns the keyboard handler, such as "vim" or "windows"
func (edit Editor) GetKeyboardHandler() string {
	return edit.Call("getKeyboardHandler").String()
}

// GetLastSearchOptions returns an object containing all the search options. For more information on `options`, see `Search`.
func (edit Editor) GetLastSearchOptions() *js.Object {
	return edit.Call("GetLastSearchOptions")
}

// GetLastVisibleRow returns the index of the last visible row.
func (edit Editor) GetLastVisibleRow() int {
	return edit.Call("getLastVisibleRow").Int()
}

// GetNumberAt works like `EditSession.GetTokenAt()`, except it returns an integer.
func (edit Editor) GetNumberAt(row, column interface{}) int {
	return edit.Call("getNumberAt", row, column).Int()
}

// GetOverwrite returns `true` if overwrites are enabled; `false` otherwise.
func (edit Editor) GetOverwrite() bool {
	return edit.Call("getOverwrite").Bool()
}

// GetPrintMarginColumn returns the column number of where the print margin is.
func (edit Editor) GetPrintMarginColumn() int {
	return edit.Call("getPrintMarginColumn").Int()
}

/*
TODO Missing functions GetPrintMarginColumn and SetDisplayIndentGuides
*/

// SetDisplayIndentGuides is an undocumented Ace function.
func (edit Editor) SetDisplayIndentGuides(args ...interface{}) *js.Object {
	return edit.Call("setDisplayIndentGuides", args...)
}

// SetDragDelay sets the delay (in milliseconds) of the mouse drag.
func (edit Editor) SetDragDelay(dragDelay int) {
	edit.Call("setDragDelay", dragDelay)
}

// SetFadeFoldWidgets is an undocumented Ace function.
func (edit Editor) SetFadeFoldWidgets(args ...interface{}) *js.Object {
	return edit.Call("setFadeFoldWidgets", args...)
}

// SetFontSize sets a new font size (in pixels) for the text editor
func (edit Editor) SetFontSize(fontSize int) {
	edit.Call("setFontSize", fontSize)
}

// SetHighlightActiveLine determines whether or not the current line should be highlighted.
func (edit Editor) SetHighlightActiveLine(shouldHighlight bool) {
	edit.Call("setHighlightActiveLine", shouldHighlight)
}

/*
TODO Missing functions after SetHighlightActiveLine
*/

// SetTheme changes the theme of the editor.
func (edit Editor) SetTheme(theme string) {
	edit.Call("setTheme", theme)
}

// GetValue returns the content in the editor.
func (edit Editor) GetValue() string {
	return edit.Call("getValue").String()
}

// SetValue sets the value of the editor to that of the given string.
func (edit Editor) SetValue(val string) {
	edit.Call("setValue", val)
}
