package guru_player

import (
	"blackjack/guru_card"
)

type Player struct {
	playerName string
	cards      []*guru_card.Card
}

func NewPlayer(player1Name string) *Player {
	var cards []*guru_card.Card
	for i := 0; i < 5; i++ {
		cards[i] = guru_card.NewCard()
	}
	return &Player{
		playerName: player1Name,
		cards:      cards,
	}
}
