package controller

import (
	"contactsoneapp/components/guru_contacts/service"
	"contactsoneapp/components/log"
	"contactsoneapp/errors"
	"contactsoneapp/middleware/auth"
	useridtokenuseridverification "contactsoneapp/middleware/userid_token_userid_verification"
	"contactsoneapp/models/contact"
	"contactsoneapp/web"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ContactController struct {
	log     log.Log
	service *service.ContactService
}

func NewContactController(contactService *service.ContactService, log log.Log) *ContactController {
	return &ContactController{
		service: contactService,
		log:     log,
	}
}

func (controller *ContactController) RegisterRoutes(router *mux.Router) {
	contactRouter := router.PathPrefix("/user/{user-id}/contact").Subrouter()
	contactRouter.HandleFunc("/", controller.CreateContact).Methods(http.MethodPost)
	contactRouter.HandleFunc("/", controller.GetAllContacts).Methods(http.MethodGet)
	contactRouter.HandleFunc("/{id}", controller.GetContactById).Methods(http.MethodGet)
	contactRouter.HandleFunc("/{id}", controller.UpdateContact).Methods(http.MethodPut)
	contactRouter.HandleFunc("/{id}", controller.DeleteContact).Methods(http.MethodDelete)
	contactRouter.Use(auth.IsUser)
	contactRouter.Use(useridtokenuseridverification.CheckUserVerify)
	fmt.Println("[Contact register routes]")
}

func (controller *ContactController) CreateContact(w http.ResponseWriter, r *http.Request) {
	newContact := contact.Contact{}
	err := web.UnmarshalJSON(r, &newContact)
	if err != nil {
		controller.log.Print(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}
	slugs := mux.Vars(r)
	idTemp, _ := strconv.Atoi(slugs["user-id"])

	newContact.UserID = uint(idTemp)
	err = controller.service.CreateContact(&newContact)
	if err != nil {
		controller.log.Print(err.Error())
		web.RespondError(w, err)
	}
	web.RespondJSON(w, http.StatusCreated, newContact)
}
func (controller *ContactController) GetAllContacts(w http.ResponseWriter, r *http.Request) {

	slugs := mux.Vars(r)
	idTemp, err := strconv.Atoi(slugs["user-id"])
	if err != nil {
		controller.log.Print(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}

	allContacts := &[]contact.Contact{}
	var totalCount int
	limit, offset := web.ParseLimitAndOffset(r)
	givenAssociations := web.ParsePreloading(r)

	err = controller.service.GetAllContacts(allContacts, uint(idTemp), &totalCount, limit, offset, givenAssociations)
	if err != nil {
		controller.log.Print(err.Error())
		web.RespondError(w, err)
		return
	}
	web.RespondJSONWithXTotalCount(w, http.StatusOK, totalCount, allContacts)

}

func (controller *ContactController) GetContactById(w http.ResponseWriter, r *http.Request) {
	requiredContact := contact.Contact{}
	slugs := mux.Vars(r)

	userIdTemp, err := strconv.Atoi(slugs["customer-id"])
	if err != nil {
		controller.log.Print(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}

	idTemp, err := strconv.Atoi(slugs["id"])
	if err != nil {
		controller.log.Print(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}
	givenAssociations := web.ParsePreloading(r)

	err = controller.service.GetContactById(&requiredContact, uint(userIdTemp), idTemp, givenAssociations)
	if err != nil {
		controller.log.Print(err.Error())
		web.RespondError(w, err)
		return
	}
	web.RespondJSONWithXTotalCount(w, http.StatusOK, 1, requiredContact)
}
func (controller *ContactController) UpdateContact(w http.ResponseWriter, r *http.Request) {
	fmt.Println("contact to update")
	contactToUpdate := contact.Contact{}
	fmt.Println(r.Body)
	err := web.UnmarshalJSON(r, &contactToUpdate)
	if err != nil {
		fmt.Println("error from UnMarshalJSON")
		controller.log.Print(err.Error())
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}
	slugs := mux.Vars(r)
	intId, err := strconv.Atoi(slugs["id"])

	if err != nil {
		controller.log.Print(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}
	contactToUpdate.ID = uint(intId)
	intId, _ = strconv.Atoi(slugs["user-id"])

	contactToUpdate.UserID = uint(intId)
	fmt.Println("contact to update")
	fmt.Println(&contactToUpdate)
	err = controller.service.UpdateContact(&contactToUpdate)
	if err != nil {
		controller.log.Print(err.Error())
		web.RespondError(w, err)
		return
	}
	web.RespondJSON(w, http.StatusOK, contactToUpdate)
}
func (controller *ContactController) DeleteContact(w http.ResponseWriter, r *http.Request) {
	controller.log.Print("delete contact call")
	contactToDelete := contact.Contact{}
	var err error
	slugs := mux.Vars(r)
	intID, err := strconv.Atoi(slugs["id"])
	if err != nil {
		controller.log.Print(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}
	// fmt.Println(intID)
	contactToDelete.ID = uint(intID)
	err = controller.service.DeleteContact(&contactToDelete)
	if err != nil {
		controller.log.Print(err.Error())
		web.RespondError(w, err)
		return
	}
	web.RespondJSON(w, http.StatusOK, "Delete contact successfull.")

}

// package controller

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"

// 	contactservice "contactsoneapp/components/guru_contacts/service"
// 	userservice "contactsoneapp/components/guru_user/service"
// 	"contactsoneapp/guru_errors"

// 	"github.com/google/uuid"
// 	"github.com/gorilla/mux"
// )

// func CreateContact(w http.ResponseWriter, r *http.Request) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()
// 	fmt.Println("Inside Create Contact Controller Function")
// 	slugs := mux.Vars(r)
// 	var newContactTemp *contactservice.Contact
// 	err := json.NewDecoder(r.Body).Decode(&newContactTemp)
// 	if err != nil {
// 		json.NewEncoder(w).Encode(err)
// 		panic(guru_errors.NewContactError(guru_errors.CreateContactFailed).GetSpecificMessage())
// 	}

// 	var currentUser *userservice.User = userservice.ReadUserById(uuid.MustParse(slugs["user-id"]))
// 	if currentUser == nil {
// 		json.NewEncoder(w).Encode(guru_errors.CreateContactFailed)
// 		panic(guru_errors.NewContactError(guru_errors.CreateContactFailed).GetSpecificMessage())
// 	}

// 	newContact := currentUser.CreateContact(
// 		newContactTemp.GetFirstName(),
// 		newContactTemp.GetLastName(),
// 	)
// 	json.NewEncoder(w).Encode(newContact)
// 	panic(guru_errors.NewContactError(guru_errors.CreateContactSuccess).GetSpecificMessage())

// 	// fmt.Println("Inside Create Contact Controller Function post request done")
// }

// func ReadContactById(w http.ResponseWriter, r *http.Request) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()
// 	fmt.Println("Inside Read Contact By Id Controller Function")
// 	slugs := mux.Vars(r)
// 	var currentUser *userservice.User = userservice.ReadUserById(uuid.MustParse(slugs["user-id"]))
// 	if currentUser == nil {
// 		json.NewEncoder(w).Encode(guru_errors.ReadContactFailed)
// 		panic(guru_errors.NewContactError(guru_errors.ReadContactFailed).GetSpecificMessage())
// 	}

// 	var requiredContact *contactservice.Contact = currentUser.ReadContactById(uuid.MustParse(slugs["contact-id"]))
// 	if requiredContact == nil {
// 		json.NewEncoder(w).Encode(guru_errors.ReadContactFailed)
// 		panic(guru_errors.NewContactError(guru_errors.ReadContactFailed).GetSpecificMessage())
// 	}
// 	json.NewEncoder(w).Encode(requiredContact)
// 	panic(guru_errors.NewContactError(guru_errors.ReadContactSuccess).GetSpecificMessage())
// 	// fmt.Println("Inside Read Contact By Id Controller Function get request done")
// }

// func ReadContactsAll(w http.ResponseWriter, r *http.Request) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()
// 	fmt.Println("Inside Read Contacts All Controller Function")
// 	slugs := mux.Vars(r)
// 	var currentUser *userservice.User = userservice.ReadUserById(uuid.MustParse(slugs["user-id"]))
// 	if currentUser == nil {
// 		json.NewEncoder(w).Encode(guru_errors.ReadContactFailed)
// 		panic(guru_errors.NewContactError(guru_errors.ReadContactFailed).GetSpecificMessage())
// 	}
// 	var requiredContacts []*contactservice.Contact = currentUser.ReadAllContact()
// 	if requiredContacts == nil {
// 		json.NewEncoder(w).Encode(guru_errors.ReadContactFailed)
// 		panic(guru_errors.NewContactError(guru_errors.ReadContactFailed).GetSpecificMessage())
// 	}
// 	json.NewEncoder(w).Encode(requiredContacts)
// 	panic(guru_errors.NewContactError(guru_errors.ReadContactSuccess).GetSpecificMessage())
// 	// fmt.Println("Inside Read Contacts All Controller Function get request done")
// }

// func UpdateContactById(w http.ResponseWriter, r *http.Request) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()
// 	fmt.Println("Inside Update Contact By Id Controller Function")
// 	slugs := mux.Vars(r)
// 	var newContactTemp *contactservice.Contact
// 	err := json.NewDecoder(r.Body).Decode(&newContactTemp)
// 	if err != nil {
// 		json.NewEncoder(w).Encode(err)
// 		panic(guru_errors.NewContactError(guru_errors.UpdateContactFailed).GetSpecificMessage())
// 	}
// 	var currentUser *userservice.User = userservice.ReadUserById(uuid.MustParse(slugs["user-id"]))
// 	if currentUser == nil {
// 		json.NewEncoder(w).Encode(guru_errors.UpdateContactFailed)
// 		panic(guru_errors.NewContactError(guru_errors.UpdateContactFailed).GetSpecificMessage())
// 	}
// 	updatedContact := currentUser.UpdateContactObject(
// 		uuid.MustParse(slugs["contact-id"]),
// 		newContactTemp,
// 	)
// 	json.NewEncoder(w).Encode(updatedContact)
// 	panic(guru_errors.NewContactError(guru_errors.UpdateContactSuccess).GetSpecificMessage())
// 	// fmt.Println("Inside Update Contact By Id Controller Function put request done")

// }

// func DeleteContactById(w http.ResponseWriter, r *http.Request) {
// 	defer func() {
// 		if err := recover(); err != nil
// 			fmt.Println(err)
// 		}
// 	}()
// 	fmt.Println("Inside Delete Contact By Id Controller Function")
// 	slugs := mux.Vars(r)
// 	var currentUser *userservice.User = userservice.ReadUserById(uuid.MustParse(slugs["user-id"]))
// 	if currentUser == nil {
// 		json.NewEncoder(w).Encode(guru_errors.DeleteContactFailed)
// 		panic(guru_errors.NewContactError(guru_errors.DeleteContactFailed).GetSpecificMessage())
// 	}
// 	deletedContact := currentUser.DeleteContact(uuid.MustParse(slugs["contact-id"]))
// 	json.NewEncoder(w).Encode(deletedContact)
// 	panic(guru_errors.NewContactError(guru_errors.DeleteContactSuccess).GetSpecificMessage())
// 	// fmt.Println("Inside Delete Contacts By Id Controller Function delete request done")
// }
