package service

import (
	"bankingapp/models/offer"
	"bankingapp/repository"

	"github.com/jinzhu/gorm"
)

type OfferService struct {
	db           *gorm.DB
	repository   repository.Repository
	associations []string
}

func NewOfferService(db *gorm.DB, repo repository.Repository) *OfferService {
	return &OfferService{
		db:           db,
		repository:   repo,
		associations: []string{},
	}
}

// func (offerService *OfferService) doesOfferExist(Id uint) error {
// 	exists, err := repository.DoesRecordExist(offerService.db, int(Id), offer.Offer{}, repository.Filter("`id` = ?", Id))

// 	if !exists || err != nil {
// 		return errors.New("data id is invalid")
// 	}
// 	return nil

// }

func (offerService *OfferService) RegisterOffer(newOffer *offer.Offer) error {
	uow := repository.NewUnitOfWork(offerService.db, false)
	defer uow.RollBack()

	err := offerService.repository.Add(uow, newOffer)
	if err != nil {
		uow.RollBack()
		return err
	}
	uow.Commit()
	return nil
}

func (offerService *OfferService) GetAllOffers(allOffers *[]offer.Offer, bankIdTemp uint, totalCount *int, limit, offset int, givenAssociations []string) error {
	uow := repository.NewUnitOfWork(offerService.db, true)

	defer uow.RollBack()

	requiredAssociations := repository.FilterPreloading(offerService.associations, givenAssociations)

	// fmt.Println(requiredAssociations)
	err := offerService.repository.GetAll(uow, allOffers, repository.Filter("`bank_id` = ?", bankIdTemp), repository.Paginate(limit, offset, totalCount), repository.Preload(requiredAssociations))

	if err != nil {
		return err
	}

	// *totalCount = len(*allOffers)
	uow.Commit()
	return nil
}

func (offerService *OfferService) GetOfferById(requiredOffer *offer.Offer, givenAssociations []string) error {
	uow := repository.NewUnitOfWork(offerService.db, true)
	defer uow.RollBack()

	requiredAssociations := repository.FilterPreloading(offerService.associations, givenAssociations)
	err := offerService.repository.GetRecordForId(uow, requiredOffer.ID, requiredOffer, repository.Filter("`bank_id` = ?", requiredOffer.BankID), repository.Preload(requiredAssociations))

	if err != nil {
		return err
	}
	uow.Commit()
	return nil
}
