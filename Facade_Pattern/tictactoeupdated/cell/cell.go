package cell

type Cell struct {
	mark string
}

func NewCell() *Cell {
	// fmt.Println("Cell started")
	return &Cell{
		mark: "E",
	}
}

func (c *Cell) IsEmpty() bool {
	return c.mark == "E"
}

func (c *Cell) MarkCell(symbol string) bool {
	c.mark = symbol
	return true
}
func (c *Cell) GetMark() string {
	return c.mark
}
