package core

type ValidationError struct {
	field string
	tag   string
}

func NewValidationError(
	field string,
	tag string,
) ValidationError {

	return ValidationError{
		field: field,
		tag:   tag,
	}
}

func (v ValidationError) Error() string {
	return v.tag
}

func (v ValidationError) Field() string {
	return v.field
}

func (v ValidationError) Tag() string {
	return v.tag
}
