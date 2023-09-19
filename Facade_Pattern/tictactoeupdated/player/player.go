package player

type Player struct {
	name   string
	Symbol string
}

func NewPlayer(name, symbol string) *Player {
	// fmt.Println("Player start")
	return &Player{
		name:   name,
		Symbol: symbol,
	}
}

func (p *Player) GetSymbol() string {
	return p.Symbol
}
func (p *Player) GetName() string {
	return p.name
}
