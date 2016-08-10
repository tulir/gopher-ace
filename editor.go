package ace

import (
	"github.com/gopherjs/gopherjs/js"
)

// Editor is a wrapper for the Ace Editor type.
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
func (edit Editor) OnChangeSession(f func(oldSession, session EditSession)) {
	edit.Call("on", "changeSession", func(dat *js.Object) {
		f(EditSession{dat.Get("oldSession")}, EditSession{dat.Get("session")})
	})
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
func (edit Editor) OnPaste(f func(pasted string) string) {
	edit.Call("on", "paste", func(dat *js.Object) {
		dat.Set("text", f(dat.Get("text").String()))
	})
}

// On binds the given function to the given Ace event.
func (edit Editor) On(on string, f interface{}) {
	edit.Call("on", on, f)
}

// AddSelectionMarker adds the selection and cursor.
func (edit Editor) AddSelectionMarker(obj Range) Range {
	return Range{edit.Call("addSelectionMarker", obj)}
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

// GetLastSearchOptions returns an object containing all the search options.
// For more information on `options`, see `Search`.
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

// GetReadOnly returns `true` if the editor is set to read-only mode.
func (edit Editor) GetReadOnly() bool {
	return edit.Call("getReadOnly").Bool()
}

// GetScrollSpeed returns the value indicating how fast the mouse scroll speed is (in milliseconds).
func (edit Editor) GetScrollSpeed() int {
	return edit.Call("getScrollSpeed").Int()
}

// GetSelection returns the selection object.
func (edit Editor) GetSelection() *js.Object {
	return edit.Call("getSelection")
}

// GetSelectionRange returns the `Range` for the selected text.
func (edit Editor) GetSelectionRange() Range {
	return Range{edit.Call("getSelectionRange")}
}

// GetSelectionStyle returns the current selection style.
func (edit Editor) GetSelectionStyle() string {
	return edit.Call("getSelectionStyle").String()
}

// GetShowFoldWidgets returns `true` if the fold widgets are shown.
func (edit Editor) GetShowFoldWidgets() bool {
	return edit.Call("getShowFoldWidgets").Bool()
}

// GetShowInvisibles returns `true` if invisible characters are being shown.
func (edit Editor) GetShowInvisibles() bool {
	return edit.Call("getShowInvisibles").Bool()
}

// GetShowPrintMargin returns `true` if the print margin is being shown.
func (edit Editor) GetShowPrintMargin() bool {
	return edit.Call("getShowPrintMargin").Bool()
}

// GetValue returns the path of the current theme.
func (edit Editor) GetValue() string {
	return edit.Call("getValue").String()
}

// GetTheme returns the content in the editor.
func (edit Editor) GetTheme() string {
	return edit.Call("getTheme").String()
}

// GetWrapBehavioursEnabled returns `true` if the wrapping behaviors are currently enabled.
func (edit Editor) GetWrapBehavioursEnabled() bool {
	return edit.Call("getWrapBehavioursEnabled").Bool()
}

// GotoLine moves the cursor to the specified line number, and also into the indiciated column.
func (edit Editor) GotoLine(lineNumber, column int, animate bool) {
	edit.Call("gotoLine", lineNumber, column, animate)
}

// GotoPageDown shifts the document to wherever "page down" is, as well as moving the cursor position.
func (edit Editor) GotoPageDown() {
	edit.Call("gotoPageDown")
}

// GotoPageUp shifts the document to wherever "page up" is, as well as moving the cursor position.
func (edit Editor) GotoPageUp() {
	edit.Call("gotoPageUp")
}

// Indent indents the current line,
func (edit Editor) Indent() {
	edit.Call("indent")
}

// Insert inserts `text` into wherever the cursor is pointing.
func (edit Editor) Insert(text string) {
	edit.Call("insert", text)
}

// IsFocused returns `true` if the current `textInput` is in focus.
func (edit Editor) IsFocused() bool {
	return edit.Call("isFocused").Bool()
}

// IsRowFullyVisible indicates if the entire row is currently visible on the screen.
func (edit Editor) IsRowFullyVisible(row int) bool {
	return edit.Call("isRowFullyVisible", row).Bool()
}

// IsRowVisible indicates if the row is currently visible on the screen.
func (edit Editor) IsRowVisible(row int) bool {
	return edit.Call("isRowVisible", row).Bool()
}

// JumpToMatching moves the cursor's row and column to the next matching bracket.
func (edit Editor) JumpToMatching(sel interface{}) {
	edit.Call("jumpToMatching", sel)
}

// ModifyNumber - change the value of the character before the cursor by `amount` if it is a number.
func (edit Editor) ModifyNumber(amount int) {
	edit.Call("modifyNumber", amount)
}

// MoveCursorTo moves the cursor to the specified row and column.
// Note that this does not de-select the current selection.
func (edit Editor) MoveCursorTo(row, column int) {
	edit.Call("moveCursorTo", row, column)
}

// MoveCursorToPosition moves the cursor to the position indicated by `pos.row` and `pos.column`.
func (edit Editor) MoveCursorToPosition(pos map[string]interface{}) {
	edit.Call("moveCursorToPosition", pos)
}

// MoveLinesDown shifts all the selected lines down one row.
func (edit Editor) MoveLinesDown() int {
	return edit.Call("moveLinesDown").Int()
}

// MoveLinesUp shifts all the selected lines up one row.
func (edit Editor) MoveLinesUp() int {
	return edit.Call("moveLinesUp").Int()
}

// MoveText is an undocumented Ace function.
func (edit Editor) MoveText(args ...interface{}) *js.Object {
	return edit.Call("moveText", args...)
}

// NavigateUp moves the cursor up in the document the specified number of times.
// Note that this does de-select the current selection.
func (edit Editor) NavigateUp(times int) {
	edit.Call("navigateUp", times)
}

// NavigateDown moves the cursor down in the document the specified number of times.
// Note that this does de-select the current selection.
func (edit Editor) NavigateDown(times int) {
	edit.Call("navigateDown", times)
}

// NavigateLeft moves the cursor left in the document the specified number of times.
// Note that this does de-select the current selection.
func (edit Editor) NavigateLeft(times int) {
	edit.Call("navigateLeft", times)
}

// NavigateRight moves the cursor right in the document the specified number of times.
// Note that this does de-select the current selection.
func (edit Editor) NavigateRight(times int) {
	edit.Call("navigateRight", times)
}

// NavigateTo moves the cursor to the specified row and column.
// Note that this does de-select the current selection.
func (edit Editor) NavigateTo(row, column int) {
	edit.Call("navigateTo", row, column)
}

// NavigateFileStart moves the cursor to the start of the current file.
// Note that this does de-select the current selection.
func (edit Editor) NavigateFileStart() {
	edit.Call("navigateFileStart")
}

// NavigateFileEnd moves the cursor to the end of the current file.
// Note that this does de-select the current selection.
func (edit Editor) NavigateFileEnd() {
	edit.Call("navigateFileEnd")
}

// NavigateLineStart moves the cursor to the start of the current line.
// Note that this does de-select the current selection.
func (edit Editor) NavigateLineStart() {
	edit.Call("navigateLineStart")
}

// NavigateLineEnd moves the cursor to the end of the current line.
// Note that this does de-select the current selection.
func (edit Editor) NavigateLineEnd() {
	edit.Call("navigateLineEnd")
}

// NavigateWordLeft moves the cursor to the word immediately to the left of the current position.
// Note that this does de-select the current selection.
func (edit Editor) NavigateWordLeft() {
	edit.Call("navigateLineLeft")
}

// NavigateWordRight moves the cursor to the word immediately to the right of the current position.
// Note that this does de-select the current selection.
func (edit Editor) NavigateWordRight() {
	edit.Call("navigateLineRight")
}

// Redo performs a redo operation on the document, reimplementing the last change.
func (edit Editor) Redo() {
	edit.Call("redo")
}

// Remove removes words of text from the editor.
// A "word" is defined as a string of characters bookended by whitespace.
func (edit Editor) Remove(dir string) {
	edit.Call("remove", dir)
}

// RemoveLines removes all the lines in the current selection.
func (edit Editor) RemoveLines() {
	edit.Call("removeLines")
}

// RemoveSelectionMarker removes the selection marker.
func (edit Editor) RemoveSelectionMarker(rangee Range) {
	edit.Call("removeSelectionMarker", rangee)
}

// RemoveToLineStart removes all the words to the left of the current selection,
// until the start of the line.
func (edit Editor) RemoveToLineStart() {
	edit.Call("removeToLineStart")
}

// RemoveToLineEnd removes all the words to the right of the current selection,
// until the end of the line.
func (edit Editor) RemoveToLineEnd() {
	edit.Call("removeToLineEnd")
}

// RemoveToLineLeft removes the word directly to the left of the current selection.
func (edit Editor) RemoveToLineLeft() {
	edit.Call("removeToLineLeft")
}

// RemoveToLineRight removes the word directly to the right of the current selection.
func (edit Editor) RemoveToLineRight() {
	edit.Call("removeToLineRight")
}

// Replace replaces the first occurance of `options.needle` with the value in `replacement`.
func (edit Editor) Replace(replacement string, options interface{}) {
	edit.Call("replace", replacement, options)
}

// ReplaceAll replaces all occurances of `options.needle` with the value in `replacement`.
func (edit Editor) ReplaceAll(replacement string, options interface{}) {
	edit.Call("replaceAll", replacement, options)
}

// RevealRange is an undocumented Ace function.
func (edit Editor) RevealRange(args ...interface{}) *js.Object {
	return edit.Call("revealRange", args...)
}

// ScrollPageUp scrolls the document to wherever "page up" is, without changing the cursor position.
func (edit Editor) ScrollPageUp() {
	edit.Call("scrollPageUp")
}

// ScrollPageDown scrolls the document to wherever "page down" is, without changing the cursor position.
func (edit Editor) ScrollPageDown() {
	edit.Call("scrollPageDown")
}

// ScrollToLine scrolls to a line.
// If center is true, it puts the line in middle of screen (or attempts to).
func (edit Editor) ScrollToLine(line int, center, animate bool, callback func()) {
	edit.Call("scrollToLine", line, center, animate, callback)
}

// ScrollToRow moves the editor to the specified row.
func (edit Editor) ScrollToRow(row interface{}) {
	edit.Call("scrollToRow", row)
}

// SelectAll Selects all the text in editor.
func (edit Editor) SelectAll() {
	edit.Call("selectAll")
}

// SelectMore finds the next occurence of text in an active selection and adds it to the selections.
func (edit Editor) SelectMore(dir int, skip bool) {
	edit.Call("selctMore", dir, skip)
}

// SelectMoreLines adds a cursor above or below the active cursor.
func (edit Editor) SelectMoreLines(dir int, skip bool) {
	edit.Call("selectMoreLines", dir, skip)
}

// SelectPageDown selects the text from the current position of the document until where a "page down" finishes.
func (edit Editor) SelectPageDown() {
	edit.Call("selectPageDown")
}

// SelectPageUp selects the text from the current position of the document until where a "page up" finishes.
func (edit Editor) SelectPageUp() {
	edit.Call("selectPageUp")
}

// SetAnimatedScroll is an undocumented Ace function.
func (edit Editor) SetAnimatedScroll(args ...interface{}) *js.Object {
	return edit.Call("setAnimatedScroll")
}

// SetBehavioursEnabled specifies whether to use behaviors or not.
// "Behaviors" in this case is the auto-pairing of special characters,
// like quotation marks, parenthesis, or brackets.
func (edit Editor) SetBehavioursEnabled(enabled bool) {
	edit.Call("setBehavioursEnabled", enabled)
}

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

// SetHighlightGutterLine is an undocumented Ace function.
func (edit Editor) SetHighlightGutterLine() {
	edit.Call("setHighlightGutterLine")
}

// SetHighlightSelectedWord determines if the currently selected word should be highlighted.
func (edit Editor) SetHighlightSelectedWord(shouldHighlight bool) {
	edit.Call("setHighlightSelectedWord", shouldHighlight)
}

// SetKeyboardHandler determines if the currently selected word should be highlighted.
func (edit Editor) SetKeyboardHandler(keyboardHandler string) {
	edit.Call("setKeyboardHandler", keyboardHandler)
}

// SetOverwrite - Pass in `true` to enable overwrites in your session, or `false` to disable.
// If overwrites is enabled, any text you enter will type over any text after it.
// If the value of `overwrite` changes, this function also emites the `changeOverwrite` event.
func (edit Editor) SetOverwrite(overwrite bool) {
	edit.Call("setOverwrite", overwrite)
}

// SetPrintMarginColumn sets the column defining where the print margin should be.
func (edit Editor) SetPrintMarginColumn(showPrintMargin int) {
	edit.Call("setPrintMarginColumn", showPrintMargin)
}

// SetReadOnly - If `readOnly` is true, then the editor is set to read-only mode,
// and none of the content can change.
func (edit Editor) SetReadOnly(readOnly bool) {
	edit.Call("setReadOnly", readOnly)
}

// SetScrollSpeed sets how fast the mouse scrolling should do.
func (edit Editor) SetScrollSpeed(speed int) {
	edit.Call("setScrollSpeed", speed)
}

// SetSelectionStyle indicates how selections should occur.
func (edit Editor) SetSelectionStyle(style string) {
	edit.Call("setSelectionStyle", style)
}

// SetSession sets a new editsession to use. This method also emits the `changeSession` event.
func (edit Editor) SetSession(session EditSession) {
	edit.Call("setSession", session)
}

// SetShowFoldWidgets indicates whether the fold widgets are shown or not.
func (edit Editor) SetShowFoldWidgets(show bool) {
	edit.Call("setShowFoldWidgets", show)
}

// SetShowInvisibles - If `showInvisibles` is set to `true`, invisible characters—like spaces or
// new lines—are show in the editor.
func (edit Editor) SetShowInvisibles(showInvisibles bool) {
	edit.Call("setShowInvisibles", showInvisibles)
}

// SetShowPrintMargin - If `showPrintMargin` is set to `true`, the print margin is shown in the editor.
func (edit Editor) SetShowPrintMargin(showPrintMargin bool) {
	edit.Call("setShowPrintMargin", showPrintMargin)
}

// SetStyle adds a new class, `style`, to the editor.
func (edit Editor) SetStyle(style string) {
	edit.Call("setStyle", style)
}

// SetTheme sets a new theme for the editor. `theme` should exist, and be a directory path,
// like `ace/theme/textmate`.
func (edit Editor) SetTheme(theme string) {
	edit.Call("setTheme", theme)
}

// SetValue sets the current document to `val`.
func (edit Editor) SetValue(val string) string {
	return edit.Call("setValue", val).String()
}

// SetValuePos sets the current document to `val`.
func (edit Editor) SetValuePos(val string, pos int) string {
	return edit.Call("setValue", val, pos).String()
}

// SetWrapBehavioursEnabled specifies whether to use wrapping behaviors or not, i.e. automatically
// wrapping the selection with characters such as brackets when such a character is typed in.
func (edit Editor) SetWrapBehavioursEnabled(enabled bool) {
	edit.Call("setWrapBehavioursEnabled", enabled)
}

// SortLines is an undocumented Ace function.
func (edit Editor) SortLines(args ...interface{}) *js.Object {
	return edit.Call("sortLines", args...)
}

// SplitLine splits the line at the current selection (by inserting an `\n`).
func (edit Editor) SplitLine() {
	edit.Call("splitLine")
}

// ToggleCommentLines either comments all the lines or uncomments all of them depending on the
// currently selected range.
func (edit Editor) ToggleCommentLines() {
	edit.Call("toggleCommentLines")
}

// ToggleOverwrite sets the value of overwrite to the opposite of whatever it currently is.
func (edit Editor) ToggleOverwrite() {
	edit.Call("toggleOverwrite")
}

// ToLowerCase converts the current selection entirely into lowercase.
func (edit Editor) ToLowerCase() {
	edit.Call("toLowerCase")
}

// ToUpperCase converts the current selection entirely into uppercase.
func (edit Editor) ToUpperCase() {
	edit.Call("toUpperCase")
}

// TransposeLetters transposes the current line.
func (edit Editor) TransposeLetters() {
	edit.Call("transposeLetters")
}

// TransposeSelections transposes the selected ranges.
func (edit Editor) TransposeSelections(dir int) {
	edit.Call("transposeSelections", dir)
}

// Undo performs an undo operation on the document, reverting the last change.
func (edit Editor) Undo() {
	edit.Call("undo")
}

// UnsetStyle removes the class `style` from the editor.
func (edit Editor) UnsetStyle(style interface{}) {
	edit.Call("unsetStyle", style)
}

// UpdateSelectionMarkers updates the cursor and marker layers.
func (edit Editor) UpdateSelectionMarkers() {
	edit.Call("updateSelectionMarkers")
}
