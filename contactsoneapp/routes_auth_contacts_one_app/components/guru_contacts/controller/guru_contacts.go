package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	contactservice "contactsoneapp/components/guru_contacts/service"
	userservice "contactsoneapp/components/guru_user/service"
	"contactsoneapp/guru_errors"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func CreateContact(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("Inside Create Contact Controller Function")
	slugs := mux.Vars(r)
	var newContactTemp *contactservice.Contact
	err := json.NewDecoder(r.Body).Decode(&newContactTemp)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		panic(guru_errors.NewContactError(guru_errors.CreateContactFailed).GetSpecificMessage())
	}

	var currentUser *userservice.User = userservice.ReadUserById(uuid.MustParse(slugs["user-id"]))
	if currentUser == nil {
		json.NewEncoder(w).Encode(guru_errors.CreateContactFailed)
		panic(guru_errors.NewContactError(guru_errors.CreateContactFailed).GetSpecificMessage())
	}

	newContact := currentUser.CreateContact(
		newContactTemp.GetFirstName(),
		newContactTemp.GetLastName(),
	)
	json.NewEncoder(w).Encode(newContact)
	panic(guru_errors.NewContactError(guru_errors.CreateContactSuccess).GetSpecificMessage())

	// fmt.Println("Inside Create Contact Controller Function post request done")
}

func ReadContactById(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("Inside Read Contact By Id Controller Function")
	slugs := mux.Vars(r)
	var currentUser *userservice.User = userservice.ReadUserById(uuid.MustParse(slugs["user-id"]))
	if currentUser == nil {
		json.NewEncoder(w).Encode(guru_errors.ReadContactFailed)
		panic(guru_errors.NewContactError(guru_errors.ReadContactFailed).GetSpecificMessage())
	}

	var requiredContact *contactservice.Contact = currentUser.ReadContactById(uuid.MustParse(slugs["contact-id"]))
	if requiredContact == nil {
		json.NewEncoder(w).Encode(guru_errors.ReadContactFailed)
		panic(guru_errors.NewContactError(guru_errors.ReadContactFailed).GetSpecificMessage())
	}
	json.NewEncoder(w).Encode(requiredContact)
	panic(guru_errors.NewContactError(guru_errors.ReadContactSuccess).GetSpecificMessage())
	// fmt.Println("Inside Read Contact By Id Controller Function get request done")
}

func ReadContactsAll(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("Inside Read Contacts All Controller Function")
	slugs := mux.Vars(r)
	var currentUser *userservice.User = userservice.ReadUserById(uuid.MustParse(slugs["user-id"]))
	if currentUser == nil {
		json.NewEncoder(w).Encode(guru_errors.ReadContactFailed)
		panic(guru_errors.NewContactError(guru_errors.ReadContactFailed).GetSpecificMessage())
	}
	var requiredContacts []*contactservice.Contact = currentUser.ReadAllContact()
	if requiredContacts == nil {
		json.NewEncoder(w).Encode(guru_errors.ReadContactFailed)
		panic(guru_errors.NewContactError(guru_errors.ReadContactFailed).GetSpecificMessage())
	}
	json.NewEncoder(w).Encode(requiredContacts)
	panic(guru_errors.NewContactError(guru_errors.ReadContactSuccess).GetSpecificMessage())
	// fmt.Println("Inside Read Contacts All Controller Function get request done")
}

func UpdateContactById(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("Inside Update Contact By Id Controller Function")
	slugs := mux.Vars(r)
	var newContactTemp *contactservice.Contact
	err := json.NewDecoder(r.Body).Decode(&newContactTemp)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		panic(guru_errors.NewContactError(guru_errors.UpdateContactFailed).GetSpecificMessage())
	}
	var currentUser *userservice.User = userservice.ReadUserById(uuid.MustParse(slugs["user-id"]))
	if currentUser == nil {
		json.NewEncoder(w).Encode(guru_errors.UpdateContactFailed)
		panic(guru_errors.NewContactError(guru_errors.UpdateContactFailed).GetSpecificMessage())
	}
	updatedContact := currentUser.UpdateContactObject(
		uuid.MustParse(slugs["contact-id"]),
		newContactTemp,
	)
	json.NewEncoder(w).Encode(updatedContact)
	panic(guru_errors.NewContactError(guru_errors.UpdateContactSuccess).GetSpecificMessage())
	// fmt.Println("Inside Update Contact By Id Controller Function put request done")

}

func DeleteContactById(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("Inside Delete Contact By Id Controller Function")
	slugs := mux.Vars(r)
	var currentUser *userservice.User = userservice.ReadUserById(uuid.MustParse(slugs["user-id"]))
	if currentUser == nil {
		json.NewEncoder(w).Encode(guru_errors.DeleteContactFailed)
		panic(guru_errors.NewContactError(guru_errors.DeleteContactFailed).GetSpecificMessage())
	}
	deletedContact := currentUser.DeleteContact(uuid.MustParse(slugs["contact-id"]))
	json.NewEncoder(w).Encode(deletedContact)
	panic(guru_errors.NewContactError(guru_errors.DeleteContactSuccess).GetSpecificMessage())
	// fmt.Println("Inside Delete Contacts By Id Controller Function delete request done")
}
