package blocks

// Arr ...
type Arr struct {
	buf  []int
	rows int
	cols int
}

func (a *Arr) get(r int, c int) int {
	return a.buf[r*a.cols+c]
}

func (a *Arr) set(r int, c int, val int) {
	a.buf[r*a.cols+c] = val
}

// NewArr ...
func NewArr(rows int, cols int) *Arr {
	return &Arr{
		rows: rows,
		cols: cols,
		buf:  make([]int, rows*cols),
	}
}

// ArrFromTemplate ..
func ArrFromTemplate(template []int, rows int) *Arr {
	clone := make([]int, len(template))
	copy(clone, template)
	return &Arr{
		rows: rows,
		cols: len(template) / rows,
		buf:  clone,
	}
}

// Clone ...
func (a Arr) Clone() *Arr {
	return ArrFromTemplate(a.buf, a.rows)
}

// CanPlace ...
func (a *Arr) CanPlace(piece *Arr, r int, c int) bool {
	for y := 0; y < piece.rows; y++ {
		for x := 0; x < piece.cols; x++ {
			p := piece.get(y, x)
			if p == 0 {
				continue
			}
			br := r + y
			bc := c + x
			if br < 0 || br >= a.rows {
				return false
			}
			if bc < 0 || bc >= a.cols {
				return false
			}
			if a.get(br, bc) != 0 {
				return false
			}
		}
	}
	return true
}

// RotateClockwise ...
func (a *Arr) RotateClockwise() *Arr {
	dst := NewArr(a.cols, a.rows)
	for r := 0; r < a.rows; r++ {
		for c := 0; c < a.cols; c++ {
			v := a.get(r, c)
			dr := c
			dc := a.rows - 1 - c
			dst.set(dr, dc, v)
		}
	}
	return dst
}

// RotateCounterClockwise ...
func (a *Arr) RotateCounterClockwise() *Arr {
	dst := NewArr(a.cols, a.rows)
	for r := 0; r < a.rows; r++ {
		for c := 0; c < a.cols; c++ {
			v := a.get(r, c)
			dr := a.cols - 1 - c
			dc := r
			dst.set(dr, dc, v)
		}
	}
	return dst
}

// Place ...
func (a *Arr) Place(piece *Arr, r int, c int) {
	for y := 0; y < piece.rows; y++ {
		for x := 0; x < piece.cols; x++ {
			p := piece.get(y, x)
			if p == 0 {
				continue
			}
			br := r + y
			bc := c + x
			a.set(br, bc, p)
		}
	}
}

// Remove ...
func (a *Arr) Remove(piece *Arr, r int, c int) {
	for y := 0; y < piece.rows; y++ {
		for x := 0; x < piece.cols; x++ {
			p := piece.get(y, x)
			if p != 0 {
				br := r + y
				bc := c + x
				a.set(br, bc, 0)
			}
		}
	}
}

// IsRowFull ...
func (a *Arr) IsRowFull(r int) bool {
	for c := 0; c < a.cols; c++ {
		if a.get(r, c) == 0 {
			return false
		}
	}
	return true
}

// ClearRow ...
func (a *Arr) ClearRow(r int) {
	for c := 0; c < a.cols; c++ {
		a.set(r, c, 0)
	}
}

// ShiftRow ...
func (a *Arr) ShiftRow(r int) {
	for c := 0; c < a.cols; c++ {
		a.set(r+1, c, a.get(r, c))
	}
}

// RemoveRow ...
func (a *Arr) RemoveRow(row int) {
	for r := row; r > 0; r-- {
		a.ShiftRow(r - 1)
	}
	a.ClearRow(0)
}

// RemoveFullRows ...
func (a *Arr) RemoveFullRows() int {
	removed := 0
	for row := a.rows - 1; row > 0; row-- {
		if a.IsRowFull(row) {
			a.RemoveRow(row)
			row++
			removed++
		}
	}
	return removed
}
