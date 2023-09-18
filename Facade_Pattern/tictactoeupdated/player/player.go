package player

type Player struct {
	Name   string
	Symbol string
}

func NewPlayer(name, symbol string) *Player {
	// fmt.Println("Player start")
	return &Player{
		Name:   name,
		Symbol: symbol,
	}
}

func (p *Player) GetSymbol() string {
	return p.Symbol
}
func (p *Player) GetName() string {
	return p.Name
}
