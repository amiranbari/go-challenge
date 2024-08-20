package uservalidator

import (
	"errors"
	userparam "github.com/amiranbari/challenge/param/user"
	errmsg "github.com/amiranbari/challenge/pkg/err_msg"
	richerror "github.com/amiranbari/challenge/rich_error"
	"github.com/amiranbari/challenge/validator"
	validation "github.com/go-ozzo/ozzo-validation"
)

func (v Validator) ValidateGetAll(req userparam.GetAllRequest) (map[string]string, error) {
	const op = "uservalidator.ValidateGetAll"

	validFields := []string{
		"first_name", "last_name",
	}
	if err := validation.ValidateStruct(&req,
		validation.Field(&req.Filter, validation.By(validator.AreFilterFieldsValid(validFields))),
	); err != nil {

		fieldErrors := make(map[string]string)

		var errV validation.Errors
		if errors.As(err, &errV) {
			for key, value := range errV {
				if value != nil {
					fieldErrors[key] = value.Error()
				}
			}
		}

		return fieldErrors, richerror.New(op).
			WithMessage(errmsg.ErrorMsgInvalidInput).
			WithKind(richerror.KindInvalid).
			WithMeta(map[string]interface{}{"req": req}).
			WithErr(err)
	}

	return map[string]string{}, nil
}
