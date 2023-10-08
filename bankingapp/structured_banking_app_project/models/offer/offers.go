package offer

import "github.com/jinzhu/gorm"

type Offer struct {
	gorm.Model
	OfferName string
	BankID    uint
}

func NewOffer(OfferName string) *Offer {

	return &Offer{
		OfferName: OfferName,
	}

}
