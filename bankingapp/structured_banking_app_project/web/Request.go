package web

import (
	"bankingapp/errors"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func UnmarshalJSON(request *http.Request, out interface{}) error {
	if request.Body == nil {
		fmt.Println("request.Body == nil error")
		errors.NewHTTPError(errors.ErrorCodeEmptyRequestBody, http.StatusBadRequest)
	}

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		fmt.Println("ioutil.ReadAll error")
		return errors.NewHTTPError(errors.ErrorCodeEmptyRequestBody, http.StatusBadRequest)
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

func ParseStartDateEndDate(r *http.Request) (startDateTimeDotTime, endDateTimeDotTime time.Time) {

	format := "2006-01-02"
	startDate := r.URL.Query().Get("start-date")
	endDate := r.URL.Query().Get("end-date")

	startDateTimeDotTime, err := time.Parse(format, startDate)
	if err != nil {
		fmt.Println("date not in proper format, default start date chosen")

		startDate = "1970-01-01"
		startDateTimeDotTime, _ = time.Parse(format, startDate)

	}

	endDateTimeDotTime, err = time.Parse(format, endDate)
	if err != nil {
		fmt.Println("date not in proper format, dafault end date chosen")
		endDate = time.Now().Format(format)
		endDateTimeDotTime, _ = time.Parse(format, endDate)
	}

	return startDateTimeDotTime, endDateTimeDotTime
}

func ParseLimitAndOffset(request *http.Request) (limit int, offset int) {
	limitGiven := request.URL.Query().Get("limit")
	offsetGiven := request.URL.Query().Get("offset")

	limit, err := strconv.Atoi(limitGiven)
	if err != nil {
		fmt.Println("Invalid limit, default limit chosen")
		limit = 5

	}
	offset, err1 := strconv.Atoi(offsetGiven)
	if err1 != nil {
		fmt.Println("Invalid offset, default offset chosen")
		offset = 0
	}

	if limit < 0 {
		fmt.Println("Invalid limit, default limit chosen")
		limit = 5
	}
	if offset < 0 {
		fmt.Println("Invalid offset, default offset chosen")
		offset = 0
	}
	return limit, offset

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
