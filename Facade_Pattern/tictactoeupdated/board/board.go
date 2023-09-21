package board

import (
	"tictactoeupdated/cell"
)

type Board struct {
	Cells [9]*cell.Cell
}

func NewBoard() *Board {
	var cells [9]*cell.Cell

	for i := 0; i < 9; i++ {
		cells[i] = cell.NewCell()

	}
	// fmt.Println("Board started")

	return &Board{
		Cells: cells,
	}
}

func (b *Board) IsEmpty(cellNumber uint) bool {
	return b.Cells[cellNumber].IsEmpty()
}

func (b *Board) MarkCell(cellNumber uint, symbol string) {
	b.Cells[cellNumber].MarkCell(symbol)
}

func (b *Board) PrintBoard() (boardMarks string) {
	for i := 0; i < 9; i++ {
		if i%3 == 0 {
			boardMarks += "\n"
		}
		boardMarks += "  " + string(b.Cells[i].GetMark()) + "  "
		// fmt.Sprintf(" %s ", (*b.Cells[i]).GetMark())
		// fmt.Print("  ", (*b.Cells[i]).GetMark(), "  ")

	}
	// fmt.Println()

	return boardMarks
}

// 0 1 2
// 3 4 5
// 6 7 8

func (b *Board) checkDiagonal() bool {
	if !b.Cells[0].IsEmpty() && b.Cells[0].GetMark() == b.Cells[4].GetMark() && b.Cells[4].GetMark() == b.Cells[8].GetMark() ||
		!b.Cells[2].IsEmpty() && b.Cells[2].GetMark() == b.Cells[4].GetMark() && b.Cells[4].GetMark() == b.Cells[6].GetMark() {
		return true
	}
	return false

}

func (b *Board) checkRows() bool {
	if !b.Cells[0].IsEmpty() && b.Cells[0].GetMark() == b.Cells[1].GetMark() && b.Cells[1].GetMark() == b.Cells[2].GetMark() ||
		!b.Cells[3].IsEmpty() && b.Cells[3].GetMark() == b.Cells[4].GetMark() && b.Cells[4].GetMark() == b.Cells[5].GetMark() ||
		!b.Cells[6].IsEmpty() && b.Cells[6].GetMark() == b.Cells[7].GetMark() && b.Cells[7].GetMark() == b.Cells[8].GetMark() {
		return true
	}
	return false

}
func (b *Board) checkColumns() bool {
	if !b.Cells[0].IsEmpty() && b.Cells[0].GetMark() == b.Cells[3].GetMark() && b.Cells[3].GetMark() == b.Cells[6].GetMark() ||
		!b.Cells[1].IsEmpty() && b.Cells[1].GetMark() == b.Cells[4].GetMark() && b.Cells[4].GetMark() == b.Cells[7].GetMark() ||
		!b.Cells[2].IsEmpty() && b.Cells[2].GetMark() == b.Cells[5].GetMark() && b.Cells[5].GetMark() == b.Cells[8].GetMark() {
		return true
	}
	return false

}
func (b *Board) CheckWin() bool {
	if b.checkDiagonal() || b.checkRows() || b.checkColumns() {
		return true

	}
	return false
}

func (b *Board) CheckDraw() (flag bool) {
	flag = true
	for i := 0; i < len(b.Cells); i++ {
		if b.Cells[i].IsEmpty() {
			flag = false
			break
		}
	}
	return flag
}
