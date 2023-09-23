package guru_contact_details

import (
	"github.com/google/uuid"
)

// var ContactDetailsId = 3

type ContactDetails struct {
	contactDetailsId uuid.UUID
	typeName         string
	typeValue        string
	isActive         bool
}

func NewContactDetails(typeName string, typeValue string) *ContactDetails {

	var newObjectOfContactDetails = &ContactDetails{
		contactDetailsId: uuid.New(),
		typeName:         typeName,
		typeValue:        typeValue,
		isActive:         true,
	}
	return newObjectOfContactDetails
}

func CreateContactDetails(typeName, typeValue string) *ContactDetails {
	return NewContactDetails(typeName, typeValue)
}

func (cd *ContactDetails) UpdateContactDetails(keyName string, keyValue string) *ContactDetails {
	cd.typeName = keyName
	cd.typeValue = keyValue
	return cd
}

// func (cd *ContactDetails) ReadContactDetails() {
// 	fmt.Printf("\nContact Details")
// 	fmt.Printf("\nType : %s and Type Value : %s", cd.typeName, cd.typeValue)
// }

// func (cd *ContactDetails) ReadContactDetails() (readContactDetails string) {
// 	readContactDetails += "Contact Details" +
// 		"\nContact Details Id: " + cd.contactDetailsId.String() +
// 		"\nType : " + cd.typeName +
// 		" and Type Value : " + cd.typeValue +
// 		"\nisActive : " + strconv.FormatBool(cd.isActive)
// 	return readContactDetails
// }

func (cd *ContactDetails) ReadContactDetails() (bool, *ContactDetails) {
	if cd.isActive {
		return true, cd
	}
	return false, cd
}

func (cd *ContactDetails) DeleteContactDetails() *ContactDetails {
	cd.typeName = ""
	cd.typeValue = ""
	cd.isActive = false
	return cd
}

func (cd *ContactDetails) GetType() string {
	return cd.typeName
}

func (cd *ContactDetails) GetContactDetailsId() uuid.UUID {
	return cd.contactDetailsId
}
