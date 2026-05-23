package validation

type FieldError interface {
	error

	Field() string
	Tag() string
}
