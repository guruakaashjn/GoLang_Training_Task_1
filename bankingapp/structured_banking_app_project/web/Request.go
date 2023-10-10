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

func ParseStartDateEndDate(r *http.Request) (startDateTimeDotTime, endDateTimeDotTime time.Time, err error) {

	format := "2006-01-02"
	startDate := r.URL.Query().Get("start-date")
	endDate := r.URL.Query().Get("end-date")

	startDateTimeDotTime, err1 := parseStartDate(startDate, format)
	endDateTimeDotTime, err2 := parseEndDate(endDate, format)
	if err1 != nil || err2 != nil {
		return time.Now(), time.Now(), errors.NewValidationError(err1.Error() + " " + err2.Error())
	}
	if startDateTimeDotTime.After(endDateTimeDotTime) {
		return time.Now(), time.Now(), errors.NewValidationError("start date cannot be greater then end date")
	}
	if startDateTimeDotTime.After(time.Now()) {
		return time.Now(), time.Now(), errors.NewValidationError("start date cannot be greater then current date time")
	}

	return startDateTimeDotTime, endDateTimeDotTime, nil
}

func parseStartDate(startDateGiven string, format string) (startDate time.Time, err error) {
	if startDateGiven == "" {
		startDateGiven = "1970-01-01"
		startDate, _ = time.Parse(format, startDateGiven)
		return startDate, nil
	}
	startDate, err = time.Parse(format, startDateGiven)
	if err != nil {
		// fmt.Println("date not in proper format, default start date chosen")
		// startDate = "1970-01-01"
		// startDateTimeDotTime, _ = time.Parse(format, startDate)

		return time.Now(), errors.NewValidationError("startdate provided is invalid")
	}
	return startDate, nil

}

func parseEndDate(endDateGiven string, format string) (endDate time.Time, err error) {
	if endDateGiven == "" {
		endDateGiven = time.Now().Format(format)
		endDate, _ = time.Parse(format, endDateGiven)
		return endDate, nil
	}
	endDate, err = time.Parse(format, endDateGiven)
	if err != nil {
		// fmt.Println("date not in proper format, default start date chosen")
		// startDate = "1970-01-01"
		// startDateTimeDotTime, _ = time.Parse(format, startDate)

		return time.Now(), errors.NewValidationError("enddate provided is invalid")
	}
	return endDate, nil

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

func ParseForLike(request *http.Request) (columnNames []string, conditions []string, operators []string, values []string) {

	columnNamesString := request.URL.Query().Get("column-name")
	if columnNamesString == "" {
		return
	}
	columnNames = strings.Split(columnNamesString, ",")
	conditions = []string{"LIKE ?"}
	operators = []string{"AND"}
	valueNamesString := request.URL.Query().Get("column-value")
	valueNames := strings.Split(valueNamesString, ",")

	values = valueNames

	for i := 0; i < len(values); i++ {
		values[i] = "%" + values[i] + "%"
	}
	fmt.Println(values)
	return
}

func ParseQueryParams(request *http.Request) (map[string]interface{}, error) {
	queryParams := strings.Split(request.URL.Query().Encode(), "&")

	queryParamsMap := make(map[string]interface{}, 0)
	for _, query := range queryParams {
		pair := strings.Split(query, "=")
		key := pair[0]
		var value string
		if len(pair) == 1 {
			value = ""
		} else {
			value = pair[1]
		}
		// queryParamsMap[key] = value

		// if key == "limit" {
		// 	limit, err := parseLimit(value)
		// 	if err != nil {
		// 		return nil, err
		// 	}
		// 	queryParamsMap["limit"] = limit
		// }
		// if key == "offset" {
		// 	offset, err := parseOffset(value)
		// 	if err != nil {
		// 		return nil, err
		// 	}
		// 	queryParamsMap["offset"] = offset
		// }
		// if key == "start-date" {
		// 	start_date, err := parseStartDate(value, "2006-01-02")
		// 	if err != nil {
		// 		return nil, err
		// 	}
		// 	queryParamsMap["start-date"] = start_date
		// }
		// if key == "end-date" {
		// 	end_date, err := parseStartDate(value, "2006-01-02")
		// 	if err != nil {
		// 		return nil, err
		// 	}
		// 	queryParamsMap["end-date"] = end_date
		// }
		// if key == "includes" {
		// 	preloads := ParsePreloading(request)
		// 	queryParamsMap["includes"] = preloads
		// }
		if key != "limit" &&
			key != "offset" &&
			key != "start-date" &&
			key != "end-date" &&
			key != "includes" && value != "" {
			queryParamsMap[key] = value
		}
		fmt.Println(queryParamsMap)

	}
	return queryParamsMap, nil

}

// func ParseStartDateEndDate(r *http.Request) (startDateTimeDotTime, endDateTimeDotTime time.Time) {

// 	format := "2006-01-02"
// 	startDate := r.URL.Query().Get("start-date")
// 	endDate := r.URL.Query().Get("end-date")

// 	startDateTimeDotTime, err := time.Parse(format, startDate)
// 	if err != nil {
// 		fmt.Println("date not in proper format, default start date chosen")

// 		startDate = "1970-01-01"
// 		startDateTimeDotTime, _ = time.Parse(format, startDate)

// 	}

// 	endDateTimeDotTime, err = time.Parse(format, endDate)
// 	if err != nil {
// 		fmt.Println("date not in proper format, dafault end date chosen")
// 		endDate = time.Now().Format(format)
// 		endDateTimeDotTime, _ = time.Parse(format, endDate)
// 	}

// 	return startDateTimeDotTime, endDateTimeDotTime
// }

// func ParseLimitAndOffset(request *http.Request) (limit int, offset int) {
// 	limitGiven := request.URL.Query().Get("limit")
// 	offsetGiven := request.URL.Query().Get("offset")

// 	limit, err := strconv.Atoi(limitGiven)
// 	if err != nil {
// 		fmt.Println("Invalid limit, default limit chosen")
// 		limit = 5

// 	}
// 	offset, err1 := strconv.Atoi(offsetGiven)
// 	if err1 != nil {
// 		fmt.Println("Invalid offset, default offset chosen")
// 		offset = 0
// 	}

// 	if limit < 0 {
// 		fmt.Println("Invalid limit, default limit chosen")
// 		limit = 5
// 	}
// 	if offset < 0 {
// 		fmt.Println("Invalid offset, default offset chosen")
// 		offset = 0
// 	}

// 	if limit > 10 {
// 		fmt.Println("Limit too large, default limit chosen")
// 		limit = 5
// 	}
// 	return limit, offset

// }

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
