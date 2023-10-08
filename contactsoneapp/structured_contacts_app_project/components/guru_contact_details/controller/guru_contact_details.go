package controller

import (
	"contactsoneapp/errors"
	"contactsoneapp/models/contactinfo"
	"contactsoneapp/web"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (controller *ContactInfoController) CreateContactDetails(w http.ResponseWriter, r *http.Request) {
	newContactDetails := contactinfo.ContactInfo{}
	err := web.UnmarshalJSON(r, &newContactDetails)
	if err != nil {
		controller.log.Print(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}
	slugs := mux.Vars(r)
	idTemp, _ := strconv.Atoi(slugs["contact-id"])
	newContactDetails.ContactRefer = uint(idTemp)
	err = controller.service.CreateContactInfo(&newContactDetails)
	if err != nil {
		controller.log.Print(err.Error())
		web.RespondError(w, err)
	}
	web.RespondJSON(w, http.StatusCreated, newContactDetails)

}

func (controller *ContactInfoController) GetAllContactDetails(w http.ResponseWriter, r *http.Request) {
	allContactDetails := &[]contactinfo.ContactInfo{}
	slugs := mux.Vars(r)
	contactIdTemp, err := strconv.Atoi(slugs["contact-id"])
	if err != nil {
		controller.log.Print(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}

	var totalCount int
	limit, offset := web.ParseLimitAndOffset(r)
	givenAssociations := web.ParsePreloading(r)

	err = controller.service.GetAllContactDetails(allContactDetails, uint(contactIdTemp), &totalCount, limit, offset, givenAssociations)
	if err != nil {
		controller.log.Print(err.Error())
		web.RespondError(w, err)
		return
	}
	web.RespondJSONWithXTotalCount(w, http.StatusOK, totalCount, allContactDetails)

}
func (controller *ContactInfoController) GetContactDetailsById(w http.ResponseWriter, r *http.Request) {
	requiredContactDetail := contactinfo.ContactInfo{}
	slugs := mux.Vars(r)
	contactIdTemp, err := strconv.Atoi(slugs["contact-id"])
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
	err = controller.service.GetContactDetailsById(&requiredContactDetail, uint(contactIdTemp), idTemp, givenAssociations)
	if err != nil {
		controller.log.Print(err.Error())
		web.RespondError(w, err)
		return
	}
	web.RespondJSONWithXTotalCount(w, http.StatusOK, 1, requiredContactDetail)
}
func (controller *ContactInfoController) UpdateContactDetails(w http.ResponseWriter, r *http.Request) {

	fmt.Println("contact details to update")
	contactDetailsToUpdate := contactinfo.ContactInfo{}
	// fmt.Println(r.Body)
	err := web.UnmarshalJSON(r, &contactDetailsToUpdate)
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
	contactDetailsToUpdate.ID = uint(intId)
	intId, _ = strconv.Atoi(slugs["contact-id"])
	contactDetailsToUpdate.ContactRefer = uint(intId)

	fmt.Println("contact details to update")
	fmt.Println(&contactDetailsToUpdate)
	err = controller.service.UpdateContactDetails(&contactDetailsToUpdate)
	if err != nil {
		controller.log.Print(err.Error())
		web.RespondError(w, err)
		return
	}
	web.RespondJSON(w, http.StatusOK, contactDetailsToUpdate)

}
func (controller *ContactInfoController) DeleteContactDetails(w http.ResponseWriter, r *http.Request) {
	controller.log.Print("delete contact details call")
	contactDetailsToDelete := contactinfo.ContactInfo{}
	var err error
	slugs := mux.Vars(r)
	intID, err := strconv.Atoi(slugs["id"])
	if err != nil {
		controller.log.Print(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}
	contactDetailsToDelete.ID = uint(intID)
	err = controller.service.DeleteContactDetails(&contactDetailsToDelete)
	if err != nil {
		controller.log.Print(err.Error())
		web.RespondError(w, err)
		return
	}
	web.RespondJSON(w, http.StatusOK, "Delete contact details successfull.")
}

// package controller

// import (
// 	contact_details_service "contactsoneapp/components/guru_contact_details/service"
// 	"contactsoneapp/guru_errors"

// 	user_service "contactsoneapp/components/guru_user/service"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"

// 	"github.com/google/uuid"
// 	"github.com/gorilla/mux"
// )

// func CreateContactDetails(w http.ResponseWriter, r *http.Request) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()
// 	fmt.Println("Inside Create Contact Details Controller Function")
// 	slugs := mux.Vars(r)
// 	var newContactDetailsTemp *contact_details_service.ContactDetails
// 	err := json.NewDecoder(r.Body).Decode(&newContactDetailsTemp)
// 	if err != nil {
// 		json.NewEncoder(w).Encode(err)
// 		panic(guru_errors.NewContactDetailsError(guru_errors.CreateContactDetailsFailed).GetSpecificMessage())
// 		// panic(err)
// 	}
// 	var currentUser *user_service.User = user_service.ReadUserById(uuid.MustParse(slugs["user-id"]))
// 	if currentUser == nil {
// 		json.NewEncoder(w).Encode(guru_errors.CreateContactDetailsFailed)
// 		panic(guru_errors.NewContactDetailsError(guru_errors.CreateContactDetailsFailed).GetSpecificMessage())

// 		// panic(guru_errors.NewUserError(guru_errors.UserDeletedStatus).GetSpecificMessage())
// 	}

// 	newContactDetails := currentUser.CreateContactDetails(
// 		uuid.MustParse(slugs["contact-id"]),
// 		newContactDetailsTemp.GetType(),
// 		newContactDetailsTemp.GetTypeValue(),
// 	)
// 	json.NewEncoder(w).Encode(newContactDetails)
// 	panic(guru_errors.NewContactDetailsError(guru_errors.CreateContactDetailsSuccess).GetSpecificMessage())

// 	// fmt.Println("Inside Create Contact Details Controller Function post request done")

// }

// func ReadContactDetailsById(w http.ResponseWriter, r *http.Request) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()
// 	fmt.Println("Inside Read Contact Details By Id Controller Function")
// 	slugs := mux.Vars(r)

// 	var currentUser *user_service.User = user_service.ReadUserById(uuid.MustParse(slugs["user-id"]))
// 	if currentUser == nil {
// 		json.NewEncoder(w).Encode(guru_errors.ReadContactDetailsFailed)
// 		panic(guru_errors.NewContactDetailsError(guru_errors.ReadContactDetailsFailed).GetSpecificMessage())
// 	}
// 	requiredContactDetails := currentUser.ReadContactDetailsById(
// 		uuid.MustParse(slugs["contact-id"]),
// 		uuid.MustParse(slugs["contact-details-id"]),
// 	)
// 	json.NewEncoder(w).Encode(requiredContactDetails)
// 	panic(guru_errors.NewContactDetailsError(guru_errors.ReadContactDetailsSuccess).GetSpecificMessage())

// 	// fmt.Println("Inside Read Contact Details By Id Controller Function get request done")
// }
// func ReadContactDetailsAll(w http.ResponseWriter, r *http.Request) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()
// 	fmt.Println("Inside Read Contact Details All Controller Function")
// 	slugs := mux.Vars(r)
// 	var currentUser *user_service.User = user_service.ReadUserById(uuid.MustParse(slugs["user-id"]))
// 	if currentUser == nil {
// 		json.NewEncoder(w).Encode(guru_errors.ReadContactDetailsFailed)
// 		panic(guru_errors.NewContactDetailsError(guru_errors.ReadContactDetailsFailed).GetSpecificMessage())
// 	}
// 	requiredContactDetailsAll := currentUser.ReadAllContactDetails(uuid.MustParse(slugs["contact-id"]))

// 	json.NewEncoder(w).Encode(requiredContactDetailsAll)
// 	panic(guru_errors.NewContactDetailsError(guru_errors.ReadContactDetailsSuccess).GetSpecificMessage())

// 	// fmt.Println("Inside Read Contact Details By Id Controller Function get request done")

// }
// func UpdateContactDetailsById(w http.ResponseWriter, r *http.Request) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()
// 	fmt.Println("Inside Update Contact Details By Id Controller Function")
// 	slugs := mux.Vars(r)
// 	var newContactDetailsTemp *contact_details_service.ContactDetails
// 	err := json.NewDecoder(r.Body).Decode(&newContactDetailsTemp)
// 	if err != nil {
// 		json.NewEncoder(w).Encode(err)
// 		panic(guru_errors.NewContactDetailsError(guru_errors.UpdateContactDetailsFailed).GetSpecificMessage())
// 		// panic(err)
// 	}
// 	var currentUser *user_service.User = user_service.ReadUserById(uuid.MustParse(slugs["user-id"]))
// 	if currentUser == nil {
// 		json.NewEncoder(w).Encode(guru_errors.UpdateContactDetailsFailed)
// 		panic(guru_errors.NewContactDetailsError(guru_errors.UpdateContactDetailsFailed).GetSpecificMessage())
// 	}
// 	updatedContactDetails := currentUser.UpdateContactDetailsObject(
// 		uuid.MustParse(slugs["contact-id"]),
// 		uuid.MustParse(slugs["contact-details-id"]),
// 		newContactDetailsTemp,
// 	)

// 	json.NewEncoder(w).Encode(updatedContactDetails)
// 	panic(guru_errors.NewContactDetailsError(guru_errors.UpdateContactDetailsSuccess).GetSpecificMessage())
// 	// fmt.Println("Inside Update Contact Details By Id Controller Function put request done")

// }

// func DeleteContactDetailsById(w http.ResponseWriter, r *http.Request) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()
// 	fmt.Println("Inside Delete Contact Details By Id Controller Function")
// 	slugs := mux.Vars(r)
// 	var currentUser *user_service.User = user_service.ReadUserById(uuid.MustParse((slugs["user-id"])))
// 	if currentUser == nil {
// 		json.NewEncoder(w).Encode(guru_errors.DeleteContactDetailsFailed)
// 		panic(guru_errors.NewContactDetailsError(guru_errors.DeleteContactDetailsFailed).GetSpecificMessage())
// 	}
// 	deletedContactDetails := currentUser.DeleteContactDetails(
// 		uuid.MustParse(slugs["contact-id"]),
// 		uuid.MustParse(slugs["contact-details-id"]),
// 	)
// 	json.NewEncoder(w).Encode(deletedContactDetails)
// 	panic(guru_errors.NewContactDetailsError(guru_errors.DeleteContactDetailsSuccess).GetSpecificMessage())
// 	// fmt.Println("Inside Delete Contact Details By Id Controller Function delete request done")
// }
