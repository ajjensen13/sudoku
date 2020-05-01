package sudoku

import (
	"reflect"
	"testing"
)

func TestRowTraverser_Next(t *testing.T) {

	called7Times := newRowTraverser(4)
	called8Times := newRowTraverser(4)
	called9Times := newRowTraverser(4)
	called10Times := newRowTraverser(4)
	for i := 0; i < 6; i++ {
		called7Times.next()
		called8Times.next()
		called9Times.next()
		called10Times.next()
	}
	called8Times.next()
	called9Times.next()
	called10Times.next()
	called9Times.next()
	called10Times.next()
	called10Times.next()

	tests := []struct {
		name  string
		r     traverser
		want  Point
		want1 bool
	}{
		{
			name:  "one",
			r:     newRowTraverser(0),
			want:  Point{Row: 0, Col: 0},
			want1: true,
		},
		{
			name:  "seven",
			r:     called7Times,
			want:  Point{Row: 4, Col: 6},
			want1: true,
		},
		{
			name:  "eight",
			r:     called8Times,
			want:  Point{Row: 4, Col: 7},
			want1: true,
		},
		{
			name:  "nine",
			r:     called9Times,
			want:  Point{Row: 4, Col: 8},
			want1: true,
		},
		{
			name:  "ten",
			r:     called10Times,
			want:  newInvalidPoint(),
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.r.next()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RowTraverser.Next() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("RowTraverser.Next() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestColTraverser_Next(t *testing.T) {

	called7Times := newColTraverser(4)
	called8Times := newColTraverser(4)
	called9Times := newColTraverser(4)
	called10Times := newColTraverser(4)
	for i := 0; i < 6; i++ {
		called7Times.next()
		called8Times.next()
		called9Times.next()
		called10Times.next()
	}
	called8Times.next()
	called9Times.next()
	called10Times.next()
	called9Times.next()
	called10Times.next()
	called10Times.next()

	tests := []struct {
		name  string
		r     traverser
		want  Point
		want1 bool
	}{
		{
			name:  "one",
			r:     newColTraverser(0),
			want:  Point{Row: 0, Col: 0},
			want1: true,
		},
		{
			name:  "seven",
			r:     called7Times,
			want:  Point{Row: 6, Col: 4},
			want1: true,
		},
		{
			name:  "eight",
			r:     called8Times,
			want:  Point{Row: 7, Col: 4},
			want1: true,
		},
		{
			name:  "nine",
			r:     called9Times,
			want:  Point{Row: 8, Col: 4},
			want1: true,
		},
		{
			name:  "ten",
			r:     called10Times,
			want:  newInvalidPoint(),
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.r.next()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ColTraverser.Next() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ColTraverser.Next() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestQuadTraverser_Next(t *testing.T) {

	called7Times := newQuadrantTraverser(Point{Row: 7, Col: 7})
	called8Times := newQuadrantTraverser(Point{Row: 7, Col: 7})
	called9Times := newQuadrantTraverser(Point{Row: 7, Col: 7})
	called10Times := newQuadrantTraverser(Point{Row: 7, Col: 7})
	for i := 0; i < 6; i++ {
		called7Times.next()
		called8Times.next()
		called9Times.next()
		called10Times.next()
	}
	called8Times.next()
	called9Times.next()
	called10Times.next()
	called9Times.next()
	called10Times.next()
	called10Times.next()

	tests := []struct {
		name  string
		r     traverser
		want  Point
		want1 bool
	}{
		{
			name:  "one",
			r:     newQuadrantTraverser(Point{Row: 1, Col: 1}),
			want:  Point{Row: 0, Col: 0},
			want1: true,
		},
		{
			name:  "seven",
			r:     called7Times,
			want:  Point{Row: 8, Col: 6},
			want1: true,
		},
		{
			name:  "eight",
			r:     called8Times,
			want:  Point{Row: 8, Col: 7},
			want1: true,
		},
		{
			name:  "nine",
			r:     called9Times,
			want:  Point{Row: 8, Col: 8},
			want1: true,
		},
		{
			name:  "ten",
			r:     called10Times,
			want:  newInvalidPoint(),
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.r.next()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ColTraverser.Next() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ColTraverser.Next() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPuzzle_Value(t *testing.T) {
	type args struct {
		p Point
	}
	tests := []struct {
		name string
		puz  *Puzzle
		args args
		want uint8
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.puz.Value(tt.args.p); got != tt.want {
				t.Errorf("Puzzle.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValue(t *testing.T) {
	type args struct {
		v uint8
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValue(tt.args.v); got != tt.want {
				t.Errorf("isValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPuzzle_HasValue(t *testing.T) {
	type args struct {
		p Point
	}
	tests := []struct {
		name string
		puz  *Puzzle
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.puz.HasValue(tt.args.p); got != tt.want {
				t.Errorf("Puzzle.HasValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPuzzle_SetValue(t *testing.T) {
	type args struct {
		p Point
		v uint8
	}
	tests := []struct {
		name string
		puz  *Puzzle
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.puz.SetValue(tt.args.p, tt.args.v)
		})
	}
}

func TestPuzzle_Solve(t *testing.T) {

	var puzzle Puzzle = [9][9]uint8{
		{0, 0, 3, 0, 2, 0, 6, 0, 0},
		{9, 0, 0, 3, 0, 5, 0, 0, 1},
		{0, 0, 1, 8, 0, 6, 4, 0, 0},
		{0, 0, 8, 1, 0, 2, 9, 0, 0},
		{7, 0, 0, 0, 0, 0, 0, 0, 8},
		{0, 0, 6, 7, 0, 8, 2, 0, 0},
		{0, 0, 2, 6, 0, 9, 5, 0, 0},
		{8, 0, 0, 2, 0, 3, 0, 0, 9},
		{0, 0, 5, 0, 1, 0, 3, 0, 0},
	}

	tests := []struct {
		name    string
		p       *Puzzle
		wantErr bool
	}{
		{
			name:    "basic",
			p:       &puzzle,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.Solve(); (err != nil) != tt.wantErr {
				t.Errorf("Puzzle.Solve() error = %v, wantErr %v\n%v", err, tt.wantErr, &puzzle)
			}
		})
	}
}

func TestPuzzle_doSolve(t *testing.T) {
	type args struct {
		t *areaTraverser
	}
	tests := []struct {
		name    string
		p       *Puzzle
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.solveOnePoint(tt.args.t); (err != nil) != tt.wantErr {
				t.Errorf("Puzzle.doSolve() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPuzzle_IsComplete(t *testing.T) {
	tests := []struct {
		name string
		p    *Puzzle
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.isComplete(); got != tt.want {
				t.Errorf("Puzzle.IsComplete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPuzzle_IsValid(t *testing.T) {
	tests := []struct {
		name string
		p    *Puzzle
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.isValid(); got != tt.want {
				t.Errorf("Puzzle.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPuzzle_isTraverserComplete(t *testing.T) {
	type args struct {
		t traverser
	}
	var puzzle Puzzle = [9][9]uint8{
		{0, 0, 3, 0, 2, 0, 6, 0, 0},
		{9, 0, 0, 3, 0, 5, 0, 0, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 0},
		{1, 2, 3, 4, 5, 6, 7, 8, 8},
		{0, 0, 6, 7, 0, 8, 2, 0, 0},
		{0, 0, 2, 6, 0, 9, 5, 0, 0},
		{8, 0, 0, 2, 0, 3, 0, 0, 9},
		{0, 0, 5, 0, 1, 0, 3, 0, 0},
	}
	tests := []struct {
		name string
		puz  *Puzzle
		args args
		want bool
	}{
		{
			name: "row 0",
			puz:  &puzzle,
			args: args{newRowTraverser(0)},
			want: false,
		},
		{
			name: "row 1",
			puz:  &puzzle,
			args: args{newRowTraverser(1)},
			want: false,
		},
		{
			name: "row 2",
			puz:  &puzzle,
			args: args{newRowTraverser(2)},
			want: true,
		},
		{
			name: "row 3",
			puz:  &puzzle,
			args: args{newRowTraverser(3)},
			want: false,
		},
		{
			name: "row 4",
			puz:  &puzzle,
			args: args{newRowTraverser(4)},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.puz.isTraverserComplete(tt.args.t); got != tt.want {
				t.Errorf("Puzzle.isTraverserComplete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPuzzle_isTraverserValid(t *testing.T) {
	type args struct {
		t traverser
	}
	var puzzle Puzzle = [9][9]uint8{
		{0, 0, 3, 0, 2, 0, 6, 0, 0},
		{9, 0, 0, 3, 0, 5, 0, 0, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 0},
		{1, 2, 3, 4, 5, 6, 7, 8, 8},
		{0, 0, 6, 7, 0, 8, 2, 0, 0},
		{0, 0, 2, 6, 0, 9, 5, 0, 0},
		{8, 0, 0, 2, 0, 3, 0, 0, 9},
		{0, 0, 5, 0, 1, 0, 3, 0, 0},
	}
	tests := []struct {
		name string
		puz  *Puzzle
		args args
		want bool
	}{
		{
			name: "row 0",
			puz:  &puzzle,
			args: args{newRowTraverser(0)},
			want: true,
		},
		{
			name: "row 1",
			puz:  &puzzle,
			args: args{newRowTraverser(1)},
			want: false,
		},
		{
			name: "row 2",
			puz:  &puzzle,
			args: args{newRowTraverser(2)},
			want: true,
		},
		{
			name: "row 3",
			puz:  &puzzle,
			args: args{newRowTraverser(3)},
			want: true,
		},
		{
			name: "row 4",
			puz:  &puzzle,
			args: args{newRowTraverser(4)},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.puz.isTraverserValid(tt.args.t); got != tt.want {
				t.Errorf("Puzzle.isTraverserValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newInvalidPoint(t *testing.T) {
	tests := []struct {
		name string
		want Point
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newInvalidPoint(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newInvalidPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPuzzleTraverser(t *testing.T) {
	tests := []struct {
		name string
		want traverser
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newPuzzleTraverser(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPuzzleTraverser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewQuadrantTraverser(t *testing.T) {
	type args struct {
		p Point
	}
	tests := []struct {
		name string
		args args
		want traverser
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newQuadrantTraverser(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewQuadrantTraverser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRowTraverser(t *testing.T) {
	type args struct {
		row uint8
	}
	tests := []struct {
		name string
		args args
		want traverser
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newRowTraverser(tt.args.row); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRowTraverser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewColTraverser(t *testing.T) {
	type args struct {
		col uint8
	}
	tests := []struct {
		name string
		args args
		want traverser
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newColTraverser(tt.args.col); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewColTraverser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAreaTraverser_Next(t *testing.T) {
	tests := []struct {
		name  string
		r     *areaTraverser
		want  Point
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.r.next()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AreaTraverser.Next() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("AreaTraverser.Next() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
