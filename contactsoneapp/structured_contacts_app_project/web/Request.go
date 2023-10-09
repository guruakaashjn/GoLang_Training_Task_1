package web

import (
	"contactsoneapp/errors"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func UnmarshalJSON(request *http.Request, out interface{}) error {
	if request.Body == nil {
		fmt.Println("request.Body == nil error")
		errors.NewHTTPError(errors.ErrorCodeEmptyRequestBody, http.StatusBadRequest)
	}

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		fmt.Println("ioutil.ReadAll error")
		return errors.NewHTTPError(err.Error(), http.StatusBadRequest)
	}
	if len(body) == 0 {
		fmt.Println("len(body) == 0 error")
		return errors.NewHTTPError(errors.ErrorCodeEmptyRequestBody, http.StatusBadRequest)
	}

	err = json.Unmarshal(body, out)
	if err != nil {
		fmt.Println("json.UnMarshal error")
		fmt.Println(body)
		fmt.Println(out)
		return errors.NewHTTPError(err.Error(), http.StatusBadRequest)
	}
	return nil

}

func ParseLimitAndOffset(request *http.Request) (limit int, offset int, err error) {

	limitGiven := request.URL.Query().Get("limit")
	offsetGiven := request.URL.Query().Get("offset")

	limit, err1 := parseLimit(limitGiven)
	offset, err2 := parseOffset(offsetGiven)

	if err1 != nil || err2 != nil {
		return 0, 0, errors.NewValidationError(err1.Error() + " " + err2.Error())
	}
	return limit, offset, nil

}

func parseLimit(limitGiven string) (limit int, err error) {
	if limitGiven == "" {
		limit = 5
		return limit, nil
	}

	limit, err = strconv.Atoi(limitGiven)
	if err != nil {
		// fmt.Println("Invalid limit, default limit chosen")
		// limit = 5
		return 0, errors.NewValidationError("Invalid limit provided")
	}
	if limit < 0 {
		return 0, errors.NewValidationError("Limit cannot be negative")
	}
	if limit > 10 {
		return 0, errors.NewValidationError("Limit chosen is too large")
	}
	return limit, nil
}

func parseOffset(offsetGiven string) (offset int, err error) {
	if offsetGiven == "" {
		offset = 0
		return offset, nil
	}
	offset, err1 := strconv.Atoi(offsetGiven)
	if err1 != nil {
		// fmt.Println("Invalid offset, default offset chosen")
		// offset = 0
		return 0, errors.NewValidationError("Invalid offset provided")
	}
	if offset < 0 {
		return 0, errors.NewValidationError("Offset cannot be negative")
	}
	return offset, nil

}

func ParsePreloading(request *http.Request) (preload []string) {
	includesRough := request.URL.Query().Get("includes")
	// fmt.Println("includes rough : ", includesRough)
	var includes []string
	if includesRough == "" {
		return nil
	}
	includes = strings.Split(includesRough, ",")
	// fmt.Println("includes : ", includes)
	return includes

}
