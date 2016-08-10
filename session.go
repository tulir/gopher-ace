package ace

import (
	"github.com/gopherjs/gopherjs/js"
)

// EditSession is a wrapper for the Ace EditSession type.
type EditSession struct {
	*js.Object
}

// GetSession returns the EditSession object for the Editor
func (edit Editor) GetSession() EditSession {
	return EditSession{edit.Call("getSession")}
}

// OnChange binds the given function to the Ace change event.
func (session EditSession) OnChange(f func(obj *js.Object)) {
	session.Call("on", "change", f)
}

// OnChangeAnnotation binds the given function to the Ace changeAnnotation event.
func (session EditSession) OnChangeAnnotation(f func()) {
	session.Call("on", "changeAnnotation", f)
}

// OnChangeBackMarker binds the given function to the Ace changeBackMarker event.
func (session EditSession) OnChangeBackMarker(f func()) {
	session.Call("on", "changeBackMarker", f)
}

// OnChangeBreakpoint binds the given function to the Ace changeBreakpoint event.
func (session EditSession) OnChangeBreakpoint(f func()) {
	session.Call("on", "changeBreakpoint", f)
}

// OnChangeFold binds the given function to the Ace changeFold event.
func (session EditSession) OnChangeFold(f func()) {
	session.Call("on", "changeFold", f)
}

// OnChangeFrontMarker binds the given function to the Ace changeFrontMarker event.
func (session EditSession) OnChangeFrontMarker(f func()) {
	session.Call("on", "changeFrontMarker", f)
}

// OnChangeMode binds the given function to the Ace changeMode event.
func (session EditSession) OnChangeMode(f func()) {
	session.Call("on", "changeMode", f)
}

// OnChangeOverwrite binds the given function to the Ace changeOverwrite event.
func (session EditSession) OnChangeOverwrite(f func()) {
	session.Call("on", "changeOverwrite", f)
}

// OnChangeScrollLeft binds the given function to the Ace changeScrollLeft event.
func (session EditSession) OnChangeScrollLeft(f func(scrollTop int)) {
	session.Call("on", "changeScrollLeft", func(obj *js.Object) {
		f(obj.Int())
	})
}

// OnChangeScrollTop binds the given function to the Ace changeScrollTop event.
func (session EditSession) OnChangeScrollTop(f func(scrollTop int)) {
	session.Call("on", "changeScrollTop", func(obj *js.Object) {
		f(obj.Int())
	})
}

// OnChangeTabSize binds the given function to the Ace changeTabSize event.
func (session EditSession) OnChangeTabSize(f func()) {
	session.Call("on", "changeTabSize", f)
}

// OnChangeWrapLimit binds the given function to the Ace changeWrapLimit event.
func (session EditSession) OnChangeWrapLimit(f func()) {
	session.Call("on", "changeWrapLimit", f)
}

// OnChangeWrapMode binds the given function to the Ace changeWrapMode event.
func (session EditSession) OnChangeWrapMode(f func()) {
	session.Call("on", "changeWrapMode", f)
}

// OnTokenizerUpdate binds the given function to the Ace tokenizerUpdate event.
func (session EditSession) OnTokenizerUpdate(f func(obj *js.Object)) {
	session.Call("on", "tokenizerUpdate", f)
}

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
func (session EditSession) AddMarker(rangee Range, class string, typee interface{}, inFront bool) int {
	return session.Call("addMarker", rangee, class, typee, inFront).Int()
}

// ClearAnnotations clears all the annotations for this session.
// This function also triggers the `changeAnnotation` event.
func (session EditSession) ClearAnnotations() {
	session.Call("clearAnnotations")
}

// ClearBreakpoint removes a breakpoint on the row number given by rows.
// This function also emits the `changeBreakpoint` event.
func (session EditSession) ClearBreakpoint(row int) {
	session.Call("clearBreakpoint", row)
}

// ClearBreakpoints removes all breakpoints on the rows.
// This function also emits the `changeBreakpoint` event.
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

// GetAnnotations returns the annotations for the `EditSession`.
func (session EditSession) GetAnnotations() *js.Object {
	return session.Call("getAnnotations")
}

// GetAWordRange gets the range of a word, including its right whitespace.
func (session EditSession) GetAWordRange(row, column int) Range {
	return Range{session.Call("getAWordRange", row, column)}
}

// GetBreakpoints returns an array of numbers, indicating which rows have breakpoints.
func (session EditSession) GetBreakpoints() int {
	return session.Call("getBreakpoints").Int()
}

// GetDocument returns the `Document` associated with this session.
func (session EditSession) GetDocument() *js.Object {
	return session.Call("getDocument")
}

// GetDocumentLastRowColumn returns the column position of the last screen row for the given document row and column.
func (session EditSession) GetDocumentLastRowColumn(docRow, docColumn int) *js.Object {
	return session.Call("getDocumentLastRowColumn", docRow, docColumn)
}

// GetDocumentLastRowColumnPosition returns the document position of the last row for the given document row and column.
func (session EditSession) GetDocumentLastRowColumnPosition(docRow, docColumn int) *js.Object {
	return session.Call("getDocumentLastRowColumnPosition", docRow, docColumn)
}

// GetLength returns the number of rows in the document.
func (session EditSession) GetLength() int {
	return session.Call("getLength").Int()
}

// GetLine returns a verbatim copy of the given line as it is in the document.
func (session EditSession) GetLine(row int) string {
	return session.Call("getLine", row).String()
}

// GetLines returns an array of strings of the rows between `firstRow` and `lastRow`.
// This function is inclusive of `lastRow`.
func (session EditSession) GetLines(firstRow, lastRow int) string {
	return session.Call("getLines", firstRow, lastRow).String()
}

// GetMarkers returns an array containing the IDs of all the markers, either front or back.
func (session EditSession) GetMarkers(inFront bool) *js.Object {
	return session.Call("getMarkers", inFront)
}

// GetMode returns the current text mode.
func (session EditSession) GetMode() *js.Object {
	return session.Call("getMode")
}

// GetNewLineMode returns the current new line mode.
func (session EditSession) GetNewLineMode() string {
	return session.Call("getNewLineMode").String()
}

// GetOverwrite returns `true` if overwrites are enabled; `false` otherwise.
func (session EditSession) GetOverwrite() bool {
	return session.Call("getOverwrite").Bool()
}

// GetRowLength returns number of screenrows in a wrapped line.
func (session EditSession) GetRowLength(row int) int {
	return session.Call("getRowLength", row).Int()
}

// GetRowSplitData returns the split data for the given row.
func (session EditSession) GetRowSplitData(row interface{}) string {
	return session.Call("getRowSplitData", row).String()
}

// GetScreenLastRowColumn returns the position (on screen) for the last character in the provided screen row.
func (session EditSession) GetScreenLastRowColumn(screenRow int) int {
	return session.Call("getScreenLastRowColumn").Int()
}

// GetScreenLength returns the length of the screen.
func (session EditSession) GetScreenLength() int {
	return session.Call("getScreenLength").Int()
}

// GetScreenTabSize returns the distance to the next tab stop at the specified screen column.
func (session EditSession) GetScreenTabSize(screenColumn int) int {
	return session.Call("getScreenTabSize", screenColumn).Int()
}

// GetScreenWidth returns the width of the screen.
func (session EditSession) GetScreenWidth() int {
	return session.Call("getScreenWidth").Int()
}

// GetScrollLeft returns the value of the distance between the left of the editor and the leftmost part of the visible content.
func (session EditSession) GetScrollLeft() int {
	return session.Call("getScrollLeft").Int()
}

// GetScrollTop returns the value of the distance between the top of the editor and the topmost part of the visible content.
func (session EditSession) GetScrollTop() int {
	return session.Call("getScrollTop").Int()
}

// GetSelection returns the selection object.
func (session EditSession) GetSelection() *js.Object {
	return session.Call("getSelection")
}

// GetState returns the state of tokenization at the end of a row.
func (session EditSession) GetState(row int) *js.Object {
	return session.Call("getState", row)
}

// GetTabSize returns the current tab size.
func (session EditSession) GetTabSize() int {
	return session.Call("getTabSize").Int()
}

// GetTabString returns the current value for tabs.
// If the user is using soft tabs, this will be a series of spaces (defined by getTabSize()); otherwise it's simply '\t'.
func (session EditSession) GetTabString() string {
	return session.Call("getTabString").String()
}

// GetTextRange returns all the text within the given range as a single string.
func (session EditSession) GetTextRange(rangee Range) string {
	return session.Call("getTextRange", rangee).String()
}

// GetTokenAt returns an object indicating the token at the current row.
// The object has two properties: `index` and `start`.
func (session EditSession) GetTokenAt(row, column int) *js.Object {
	return session.Call("getTokenAt", row, column)
}

// GetTokens starts tokenizing at the row indicated.
// Returns a list of objects of the tokenized rows.
func (session EditSession) GetTokens(row int) *js.Object {
	return session.Call("getTokens", row)
}

// GetUndoManager returns the current undo manager.
func (session EditSession) GetUndoManager() *js.Object {
	return session.Call("getUndoManager")
}

// GetUseSoftTabs returns `true` if soft tabs are being used, `false` otherwise.
func (session EditSession) GetUseSoftTabs() bool {
	return session.Call("getUseSoftTabs").Bool()
}

// GetUseWorker returns `true` if workers are being used.
func (session EditSession) GetUseWorker() bool {
	return session.Call("getUseWorker").Bool()
}

// GetUseWrapMode returns `true` if wrap mode is being used; `false` otherwise.
func (session EditSession) GetUseWrapMode() bool {
	return session.Call("getUseWrapMode").Bool()
}

// GetValue returns the current `Document` as a string.
func (session EditSession) GetValue() string {
	return session.Call("getValue").String()
}

// GetWordRange returns the `Range` of the first word boundary it finds after the starting row and column.
func (session EditSession) GetWordRange(row, column int) *js.Object {
	return session.Call("getWordRange", row, column)
}

// GetWrapLimit returns the value of wrap limit.
func (session EditSession) GetWrapLimit() int {
	return session.Call("getWrapLimit").Int()
}

// GetWrapLimitRange returns the minimum and maximum of the wrap limit
func (session EditSession) GetWrapLimitRange() (int, int) {
	obj := session.Call("getWrapLimitRange")
	return obj.Get("min").Int(), obj.Get("max").Int()
}

// Highlight is an undocumented Ace function.
func (session EditSession) Highlight(args ...interface{}) *js.Object {
	return session.Call("highlight", args...)
}

// HighlightLines is an undocumented Ace function.
func (session EditSession) HighlightLines(args ...interface{}) *js.Object {
	return session.Call("highlightLines", args...)
}

// IndentRows indents all the rows, from `startRow` to `endRow` (inclusive),
// by prefixing each row with the token in indentString.
func (session EditSession) IndentRows(startRow, endRow int, indentString string) {
	session.Call("indentRows", startRow, endRow, indentString)
}

// Insert inserts a block of `text` and the indicated position.
func (session EditSession) Insert(row, column int, text string) *js.Object {
	return session.Call("insert", map[string]interface{}{
		"row":    row,
		"column": column,
	}, text)
}

// IsTabStop returns `true` if the character at the position is a soft tab.
func (session EditSession) IsTabStop(row, column int) bool {
	return session.Call("isTabStop", map[string]interface{}{
		"row":    row,
		"column": column,
	}).Bool()
}

// MoveLinesDown shifts all the lines in the document down one,
// starting from `firstRow` and ending at `lastRow`.
func (session EditSession) MoveLinesDown(firstRow, lastRow int) int {
	return session.Call("moveLinesDown", firstRow, lastRow).Int()
}

// MoveLinesUp shifts all the lines in the document up one,
// starting from `firstRow` and ending at `lastRow`.
func (session EditSession) MoveLinesUp(firstRow, lastRow int) int {
	return session.Call("moveLinesUp", firstRow, lastRow).Int()
}

// MoveText moves a range of text from the given range to the given position.
func (session EditSession) MoveText(fromRange Range, toRow, toColumn int) Range {
	return Range{session.Call("moveText", fromRange, map[string]interface{}{
		"row":    toRow,
		"column": toColumn,
	})}
}

// OutdentRows outdents all the rows defined by the `start` and `end` properties of `range`.
func (session EditSession) OutdentRows(rangee Range) {
	session.Call("outdentRows", rangee)
}

// Redo is an undocumented Ace function.
func (session EditSession) Redo(args ...interface{}) *js.Object {
	return session.Call("redo", args...)
}

// RedoChanges re-implements a previously undone change to your document.
func (session EditSession) RedoChanges(deltas []interface{}, dontSelect bool) Range {
	return Range{session.Call("redoChanges", deltas, dontSelect)}
}

// Remove removes the `range` from the document.
func (session EditSession) Remove(rangee Range) *js.Object {
	return session.Call("remove", rangee)
}

// RemoveGutterDecoration removes `className` from the `row`.
func (session EditSession) RemoveGutterDecoration(row int, className string) {
	session.Call("removeGutterDecoration", row, className)
}

// RemoveMarker removes the marker with the specified ID.
// If this marker was in front, the `changeFrontMarker` event is emitted.
// If the marker was in the back, the `changeBackMarker` event is emitted.
func (session EditSession) RemoveMarker(markerID int) {
	session.Call("removeMarker", markerID)
}

// Replace replaces a range in the document with the new `text`.
func (session EditSession) Replace(rangee Range, text string) *js.Object {
	return session.Call("replace", rangee, text)
}

// Reset is an undocumented Ace function.
func (session EditSession) Reset(args ...interface{}) *js.Object {
	return session.Call("reset", args...)
}

// ResetCaches is an undocumented Ace function.
func (session EditSession) ResetCaches(args ...interface{}) *js.Object {
	return session.Call("resetCaches", args...)
}

// ScreenToDocumentColumn is an undocumented Ace function.
func (session EditSession) ScreenToDocumentColumn(args ...interface{}) *js.Object {
	return session.Call("screenToDocumentColumn", args...)
}

// ScreenToDocumentPosition converts characters coordinates on the screen to characters coordinates
// within the document.
// This takes into account code folding, word wrap, tab size, and any other visual modifications.
func (session EditSession) ScreenToDocumentPosition(screenRow, screenColumn int) *js.Object {
	return session.Call("screenToDocumentPosition", screenRow, screenColumn)
}

// ScreenToDocumentRow is an undocumented Ace function.
func (session EditSession) ScreenToDocumentRow(args ...interface{}) *js.Object {
	return session.Call("screenToDocumentRow", args...)
}

// SetAnnotations sets annotations for the `EditSession`.
// This functions emits the `changeAnnotation` event.
func (session EditSession) SetAnnotations(annotations []interface{}) {
	session.Call("setAnnotations", annotations)
}

// SetBreakpoint sets a breakpoint on the row number given by `rows`.
// This function also emits the `changeBreakpoint` event.
func (session EditSession) SetBreakpoint(row int, className string) {
	session.Call("setBreakpoint", row, className)
}

// SetBreakpoints sets a breakpoint on every row given by `rows`.
// This function also emits the `changeBreakpoint` event.
func (session EditSession) SetBreakpoints(rows []interface{}) {
	session.Call("setBreakpoints", rows)
}

// SetDocument sets the `EditSession` to point to a new `Document`. If a `BackgroundTokenizer` exists, it also points to `doc`.
func (session EditSession) SetDocument(doc interface{}) {
	session.Call("setDocument", doc)
}

// SetMode is an undocumented Ace function.
func (session EditSession) SetMode(args ...interface{}) *js.Object {
	return session.Call("setMode", args...)
}

// SetNewLineMode sets the new line mode.
func (session EditSession) SetNewLineMode(newLineMode string) {
	session.Call("setNewLineMode", newLineMode)
}

// SetOverwrite - Pass in `true` to enable overwrites in your session, or `false` to disable
func (session EditSession) SetOverwrite(overwrite bool) {
	session.Call("setOverwrite", overwrite)
}

// SetScrollLeft sets the value of the distance between the left of the editor and the leftmost part of the visible content.
func (session EditSession) SetScrollLeft(scrollLeft interface{}) {
	session.Call("setScrollLeft", scrollLeft)
}

// SetScrollTop sets the scroll top value.
// This function also emits the `changeScrollTop` event.
func (session EditSession) SetScrollTop(scrollTop int) {
	session.Call("setScrollTop", scrollTop)
}

// SetTabSize set the number of spaces that define a soft tab;
// for example, passing in `4` transforms the soft tabs to be equivalent to four spaces.
// This function also emits the `changeTabSize` event.
func (session EditSession) SetTabSize(tabSize int) {
	session.Call("setTabSize", tabSize)
}

// SetUndoManager sets the undo manager.
func (session EditSession) SetUndoManager(undoManager interface{}) {
	session.Call("setUndoManager", undoManager)
}

// SetUndoSelect enables or disables highlighting of the range where an undo occured.
func (session EditSession) SetUndoSelect(enable bool) {
	session.Call("setUndoSelect", enable)
}

// SetUseSoftTabs - pass `true` to enable the use of soft tabs.
// Soft tabs means you're using spaces instead of the tab character (`\t`).
func (session EditSession) SetUseSoftTabs(useSoftTabs bool) {
	session.Call("setUseSoftTabs", useSoftTabs)
}

// SetUseWorker identifies if you want to use a worker for the `EditSession`.
func (session EditSession) SetUseWorker(useWorker bool) {
	session.Call("setUseWorker", useWorker)
}

// SetUseWrapMode sets whether or not line wrapping is enabled.
// If `useWrapMode` is different than the current value, the `changeWrapMode` event is emitted.
func (session EditSession) SetUseWrapMode(useWrapMode bool) {
	session.Call("setUseWrapMode", useWrapMode)
}

// SetValue sets the session text.
func (session EditSession) SetValue(text string) {
	session.Call("setValue", text)
}

// SetWrapLimitRange sets the boundaries of wrap.
// Either value can be null to have an unconstrained wrap, or, they can be the same number to pin the limit.
// If the wrap limits for min or max are different, this method also emits the `changeWrapMode` event.
func (session EditSession) SetWrapLimitRange(min, max int) {
	session.Call("setWrapLimitRange", min, max)
}

// ToggleOverwrite sets the value of overwrite to the opposite of whatever it currently is.
func (session EditSession) ToggleOverwrite() {
	session.Call("toggleOverwrite")
}

// ToString returns the current `Document` as a string.
func (session EditSession) ToString() string {
	return session.Call("toString").String()
}

// Undo is an undocumented Ace function.
func (session EditSession) Undo(args ...interface{}) *js.Object {
	return session.Call("undo", args...)
}

// UndoChanges reverts previous changes to your document.
func (session EditSession) UndoChanges(deltas []interface{}, dontSelect bool) Range {
	return Range{session.Call("undoChanges", deltas, dontSelect)}
}

// SetOption sets an option.
func (session EditSession) SetOption(optionName string, optionValue interface{}) {
	session.Call("setOption", optionName, optionValue)
}

// SetOptions sets all options.
func (session EditSession) SetOptions(options map[string]interface{}) {
	session.Call("setOptions", options)
}

// GetOption gets the option with teh given name.
func (session EditSession) GetOption(optionName string) *js.Object {
	return session.Call("getOption", optionName)
}

// GetOptions gets all options
func (session EditSession) GetOptions() *js.Object {
	return session.Call("getOptions")
}
