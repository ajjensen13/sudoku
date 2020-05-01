package sudoku

import (
	"bytes"
	"errors"
	"fmt"
)

const (
	// PuzzleSize is the size of a puzzle
	PuzzleSize = 9
	// MinValidValue is the minimum value of an assigned value in a puzzle
	MinValidValue = 1
	// MaxValidValue is the maximum value of an assigned value in a puzzle
	MaxValidValue = 9
	// UnassignedValue is a placeholder value for when a cell still needs a value designated
	UnassignedValue = 0

	quadrantSize = PuzzleSize / 3
)

// Puzzle is a single Sudoku puzzle. It is a 9 by 9
// square. Each cell should hold a value [MinValidValue, MaxValidValue],
// or UnassignedValue if a value has not yet been assigned
// to the cell.
type Puzzle [PuzzleSize][PuzzleSize]uint8

func (p *Puzzle) String() string {
	var buf bytes.Buffer
	for row := range p {
		fmt.Fprintf(&buf, "%v\n", p[row])
	}
	return buf.String()
}

// Value returns the value of a cell
func (p *Puzzle) Value(point Point) uint8 {
	return p[point.Row][point.Col]
}

func isValue(v uint8) bool {
	return v > 0 && v < 10
}

// HasValue returns true if a cell contains a valid value,
// otherwise false.
func (p *Puzzle) HasValue(point Point) bool {
	return isValue(p[point.Row][point.Col])
}

// SetValue sets a value in a puzzle to a valid value.
// Use ClearValue to remove a value (set it to UnassignedValue).
func (p *Puzzle) SetValue(point Point, v uint8) {
	if v > 9 || v < 1 {
		panic(fmt.Sprintf("invalid value: %v", v))
	}
	p[point.Row][point.Col] = v
}

// ClearValue removes a valid value from a cell, replacing
// it with UnassignedValue.
func (p *Puzzle) ClearValue(point Point) {
	p[point.Row][point.Col] = UnassignedValue
}

// Point represents a single coordinate in a puzzle
type Point struct {
	Row, Col uint8
}

func (p Point) String() string {
	return fmt.Sprintf("[%d,%d]", p.Row, p.Col)
}

// Solve attempts to solve a puzzle in place.
func (p *Puzzle) Solve() error {
	err := p.solveOnePoint(newPuzzleTraverser().(*areaTraverser))
	if err != nil {
		return err
	}

	if !p.isValid() {
		return errors.New("unexpected exception when solving puzzle")
	}
	if !p.isComplete() {
		return errors.New("unexpected exception when solving puzzle")
	}

	return nil
}

func (p Puzzle) traverserValuesString(n traverser) string {
	var buf bytes.Buffer
	for point, ok := n.next(); ok; point, ok = n.next() {
		fmt.Fprintf(&buf, "%d", p.Value(point))
	}
	return buf.String()
}

func (p *Puzzle) solveOnePoint(t *areaTraverser) error {
	point, ok := t.next()
	if !ok {
		return nil // solved!
	}

	if p.HasValue(point) {
		return p.solveOnePoint(t)
	}

	var seen [10]bool
	p.markSeenValues(newRowTraverser(point.Row), &seen)
	p.markSeenValues(newColTraverser(point.Col), &seen)
	p.markSeenValues(newQuadrantTraverser(point), &seen)

	for val := uint8(1); val < 10; val++ {
		if seen[val] {
			continue
		}

		p.SetValue(point, val)

		nextTraverser := *t // value copy
		err := p.solveOnePoint(&nextTraverser)
		if err == nil {
			return nil
		}
	}

	p.ClearValue(point)
	return errors.New("puzzle cannot be solved")
}

func (p *Puzzle) markSeenValues(t traverser, seen *[10]bool) {
	for point, ok := t.next(); ok; point, ok = t.next() {
		if !p.HasValue(point) {
			continue
		}
		v := p.Value(point)
		seen[v] = true
	}
}

func (p *Puzzle) isComplete() bool {
	for row := uint8(0); row < PuzzleSize; row++ {
		t := newRowTraverser(row)
		if !p.isTraverserComplete(t) {
			return false
		}
	}
	for col := uint8(0); col < PuzzleSize; col++ {
		t := newColTraverser(col)
		if !p.isTraverserComplete(t) {
			return false
		}
	}
	for row := uint8(0); row < PuzzleSize; row += quadrantSize {
		for col := uint8(0); col < PuzzleSize; col += quadrantSize {
			t := newQuadrantTraverser(Point{Row: row, Col: col})
			if !p.isTraverserComplete(t) {
				return false
			}
		}
	}
	return true
}

func (p *Puzzle) isValid() bool {
	for row := uint8(0); row < PuzzleSize; row++ {
		t := newRowTraverser(row)
		if !p.isTraverserValid(t) {
			return false
		}
	}
	for col := uint8(0); col < PuzzleSize; col++ {
		t := newColTraverser(col)
		if !p.isTraverserValid(t) {
			return false
		}
	}
	for row := uint8(0); row < PuzzleSize; row += quadrantSize {
		for col := uint8(0); col < PuzzleSize; col += quadrantSize {
			t := newQuadrantTraverser(Point{Row: row, Col: col})
			if !p.isTraverserValid(t) {
				return false
			}
		}
	}
	return true
}

func (p *Puzzle) isTraverserComplete(t traverser) bool {
	for point, ok := t.next(); ok; point, ok = t.next() {
		if !p.HasValue(point) {
			return false
		}
	}
	return true
}

func (p *Puzzle) isTraverserValid(t traverser) bool {
	var seen [10]bool
	for point, ok := t.next(); ok; point, ok = t.next() {
		if !p.HasValue(point) {
			continue
		}
		if seen[p.Value(point)] {
			return false
		}
		seen[p.Value(point)] = true
	}
	return true
}

func newInvalidPoint() Point {
	return Point{Row: 0xFF, Col: 0xFF}
}

type traverser interface {
	next() (Point, bool)
}

func newPuzzleTraverser() traverser {
	return &areaTraverser{width: PuzzleSize, height: PuzzleSize}
}

func newQuadrantTraverser(p Point) traverser {
	return &areaTraverser{Start: p, width: quadrantSize, height: quadrantSize}
}

func newRowTraverser(row uint8) traverser {
	return &areaTraverser{Start: Point{Row: row}, width: PuzzleSize, height: 1}
}

func newColTraverser(col uint8) traverser {
	return &areaTraverser{Start: Point{Col: col}, width: 1, height: PuzzleSize}
}

type areaTraverser struct {
	init, done    bool
	Start         Point
	current       Point
	width, height uint8
}

func (r *areaTraverser) next() (Point, bool) {
	if r.done {
		return newInvalidPoint(), false
	}

	if !r.init {
		r.init = true
		r.current = Point{
			Row: r.Start.Row / r.height * r.height,
			Col: r.Start.Col / r.width * r.width,
		}
		r.Start = r.current // copy by value
	}

	result := r.current // copy by value

	r.current.Col++
	if r.current.Col >= r.Start.Col+r.width {
		r.current.Col = r.Start.Col
		r.current.Row++
		if r.current.Row >= r.Start.Row+r.height {
			r.done = true
		}
	}

	return result, true
}
