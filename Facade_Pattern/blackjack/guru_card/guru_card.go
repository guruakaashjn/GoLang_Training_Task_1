package guru_card

import (
	"math/rand"
	"strconv"
)

type Card struct {
	cardNumber string
}

func NewCard() *Card {
	cardNumber := drawNewCard()
	return &Card{
		cardNumber: cardNumber,
	}
}

func (c *Card) GetCardNumber() string {
	return c.cardNumber
}

func drawNewCard() string {
	cardNumberInt := rand.Intn(12) + 1
	var cardNumber string
	switch {
	case cardNumberInt > 1 && cardNumberInt <= 10:
		cardNumber = strconv.Itoa(cardNumberInt)
	case cardNumberInt == 1:
		cardNumber = "A"
	case cardNumberInt == 11:
		cardNumber = "J"
	case cardNumberInt == 12:
		cardNumber = "Q"
	case cardNumberInt == 13:
		cardNumber = "K"
	}
	return cardNumber
}
