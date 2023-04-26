package validate

import "github.com/go-playground/validator/v10"

type ValidationErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

type Validate struct {
	validator *validator.Validate
}

func NewValidate() *Validate {
	return &Validate{
		validator: validator.New(),
	}
}

func (v *Validate) ValidateStruct(s any) []*ValidationErrorResponse {
	var errs []*ValidationErrorResponse
	err := v.validator.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var e ValidationErrorResponse
			e.FailedField = err.StructNamespace()
			e.Tag = err.Tag()
			e.Value = err.Param()
			errs = append(errs, &e)
		}
	}
	return errs
}
