package cell

type Cell struct {
	Mark string
}

func NewCell() *Cell {
	// fmt.Println("Cell started")
	return &Cell{
		Mark: "E",
	}
}

func (c *Cell) IsEmpty() bool {
	return c.Mark == "E"
}

func (c *Cell) MarkCell(symbol string) bool {
	c.Mark = symbol
	return true
}
func (c *Cell) GetMark() string {
	return c.Mark
}
