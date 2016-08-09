package ace

import (
	"github.com/gopherjs/gopherjs/js"
)

// Range is a wrapper for the Ace Range type.
type Range struct {
	*js.Object
}

// StartRow returns the starting row of this Range object.
func (r Range) StartRow() int {
	return r.Get("start").Get("row").Int()
}

// StartColumn returns the starting column of this Range object.
func (r Range) StartColumn() int {
	return r.Get("start").Get("column").Int()
}

// EndRow returns the ending row of this Range object.
func (r Range) EndRow() int {
	return r.Get("end").Get("row").Int()
}

// EndColumn returns the ending column of this Range object.
func (r Range) EndColumn() int {
	return r.Get("end").Get("column").Int()
}

// ClipRows returns the part of the current `Range` that occurs within the boundaries of `firstRow` and `lastRow` as a new `Range` object.
func (r Range) ClipRows(firstRow, lastRow int) Range {
	return Range{r.Call("clipRows", firstRow, lastRow)}
}

// Clone returns a duplicate of the calling range.
func (r Range) Clone() Range {
	return Range{r.Call("clone")}
}

// CollapseRows returns a range containing the starting and ending rows of the original range, but with a column value of `0`.
func (r Range) CollapseRows() Range {
	return Range{r.Call("collapseRows")}
}

// Compare checks the row and column points with the row and column points of the calling range.
func (r Range) Compare(row, column int) int {
	return r.Call("compare", row, column).Int()
}

// CompareStart checks the row and column points with the row and column points of the calling range.
func (r Range) CompareStart(row, column int) int {
	return r.Call("compareStart", row, column).Int()
}

// CompareEnd checks the row and column points with the row and column points of the calling range.
func (r Range) CompareEnd(row, column int) int {
	return r.Call("compareEnd", row, column).Int()
}

// CompareInside checks the row and column points with the row and column points of the calling range.
func (r Range) CompareInside(row, column int) int {
	return r.Call("compareInside", row, column).Int()
}

// ComparePoint checks the row and column points of `p` with the row and column points of the calling range.
func (r Range) ComparePoint(p Range) int {
	return r.Call("comparePoint", p).Int()
}

// CompareRange compares this range with the given range.
func (r Range) CompareRange(r2 Range) int {
	return r.Call("compareRange", r2).Int()
}

// Contains returns `true` if the `row` and `column` provided are within the given range.
// This can better be expressed as returning `true` if:
//     range.StartRow()    <= row    <= range.EndRow() &&
//     range.StartColumn() <= column <= range.EndColumn()
func (r Range) Contains(row, column int) bool {
	return r.Call("contains", row, column).Bool()
}

// ContainsRange checks the start and end points of `range` and compares them to the calling range.
// Returns `true` if the `range` is contained within the caller's range.
func (r Range) ContainsRange(r2 Range) bool {
	return r.Call("containsRange", r2).Bool()
}

// Extend changes the row and column points for the calling range for both the starting and ending points.
func (r Range) Extend(row, column int) Range {
	return Range{r.Call("extend", row, column)}
}
