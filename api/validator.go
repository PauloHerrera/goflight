package api

import (
	"time"

	"gihub.com/pauloherrera/goflight/util"
	"github.com/go-playground/validator/v10"
)

// Checks if the request date is valid and is in the future
var validDate validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if dateArg, ok := fieldLevel.Field().Interface().(string); ok {
		date, err := time.Parse("2006-01-02", dateArg)
		if err != nil {
			return false
		}

		today := time.Now()

		return date.After(today)
	}
	return false
}

// Checks if its a valid brazilian airport
var validAirport validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if codeArg, ok := fieldLevel.Field().Interface().(string); ok {
		return util.IsValidAirport(codeArg)
	}
	return false
}
