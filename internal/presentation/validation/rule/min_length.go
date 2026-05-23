package rule

import (
	"github.com/arthurhzna/go-clean-architecture/internal/presentation/validation"
)

func MinLength(
	field string,
	value string,
	min int,
) Rule {

	return func() validation.FieldError {

		if len(value) < min {

			return validation.NewValidationError(
				field,
				"min_length",
			)
		}

		return nil
	}
}
