package guru_contact_details

import (
	"fmt"

	"github.com/google/uuid"
)

var ContactDetailsId = 3

type ContactDetails struct {
	contact_Details_ID uuid.UUID
	typeName           string
	typeValue          string
}

func NewContactDetails(typeName string, typeValue string) *ContactDetails {

	var newObjectOfContactDetails = &ContactDetails{
		contact_Details_ID: uuid.New(),
		typeName:           typeName,
		typeValue:          typeValue,
	}
	return newObjectOfContactDetails
}

func (cd *ContactDetails) UpdateContactDetails(keyName string, keyValue string) bool {
	cd.typeName = keyName
	cd.typeValue = keyValue
	return true
}

func (cd *ContactDetails) ReadContactDetails() {
	fmt.Printf("\nContact Details")
	fmt.Printf("\nType : %s and Type Value : %s", cd.typeName, cd.typeValue)
}

func (cd *ContactDetails) DeleteContactDetails() {
	cd.typeName = ""
	cd.typeValue = ""
}

func (cd *ContactDetails) GetType() string {
	return cd.typeName
}
