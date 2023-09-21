package guru_player

import (
	"blackjack/guru_card"
	"strconv"
)

type Player struct {
	playerName string
	hold       bool
	cards      []*guru_card.Card
}

func NewPlayer(player1Name string) *Player {
	var cards []*guru_card.Card
	// cards = append(cards, guru_card.NewCard())
	return &Player{
		playerName: player1Name,
		hold:       false,
		cards:      cards,
	}
}
func (p *Player) PrintCards() (cardsInHand string) {
	cardsInHand += "Cards in hand of Player " + p.playerName + " : "
	// fmt.Printf("Cards in hand of Player %s : ", p.playerName)
	for i := 0; i < len(p.cards); i++ {
		cardsInHand += " " + p.cards[i].GetCardNumber() + " "
		// fmt.Printf(" %s ", p.cards[i].GetCardNumber())
	}
	return cardsInHand
}

func (p *Player) SumOfCards() (sum int) {
	sum = 0
	for i := 0; i < len(p.cards); i++ {
		cardNumber := p.cards[i].GetCardNumber()
		var cardNumberInt int = 0
		switch cardNumber {

		case "A":
			cardNumberInt = 1
		case "J":
			cardNumberInt = 11
		case "Q":
			cardNumberInt = 12
		case "K":
			cardNumberInt = 13
		default:
			// fmt.Println("CN : " + cardNumber)
			rough, err := strconv.Atoi(cardNumber)
			if err == nil {
				cardNumberInt = rough
				// fmt.Printf("CNI INNER: %d", cardNumberInt)
			}
			// fmt.Printf("CNI : %d %d", rough, cardNumberInt)

		}
		// fmt.Println(cardNumberInt)

		sum = sum + cardNumberInt
	}
	return sum
}

func (p *Player) AddCardToDeck() {
	p.cards = append(p.cards, guru_card.NewCard())
}

func (p *Player) SetHold() {
	p.hold = true
}

func (p *Player) GetHold() bool {
	return p.hold
}
func (p *Player) GetPlayerName() string {
	return p.playerName
}
