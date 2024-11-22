package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/yantong1775/vaultCore/util"
)

var validCurrency validator.Func = func(fl validator.FieldLevel) bool {
	if currency, ok := fl.Field().Interface().(string); ok {
		// check if the currency is valid
		return util.IsSupportedCurrency(currency)
	}
	return false
}
