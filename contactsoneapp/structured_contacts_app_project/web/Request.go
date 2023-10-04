package web

import (
	"contactsoneapp/errors"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
