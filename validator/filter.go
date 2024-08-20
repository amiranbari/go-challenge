package validator

import (
	"fmt"
	params "github.com/amiranbari/challenge/param"
	errmsg "github.com/amiranbari/challenge/pkg/err_msg"
	validation "github.com/go-ozzo/ozzo-validation"
	"slices"
)

func AreFilterFieldsValid(validFilters []string) validation.RuleFunc {
	return func(value interface{}) error {
		filters, ok := value.(params.FilterRequest)
		if !ok {
			return fmt.Errorf(errmsg.ErrorMsgSomethingWentWrong)
		}
		for filter := range filters {
			fmt.Println(validFilters, filter, slices.Contains(validFilters, filter))
			if !slices.Contains(validFilters, filter) {
				return fmt.Errorf(errmsg.ErrorMsgFiltersAreNotValid)
			}
		}

		return nil
	}
}
