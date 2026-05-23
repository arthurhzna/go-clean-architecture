package rule

import (
	"github.com/arthurhzna/go-clean-architecture/internal/presentation/validation"
)

func Equal(
	field string,
	value string,
	compare string,
) Rule {

	return func() validation.FieldError {

		if value != compare {

			return validation.NewValidationError(
				field,
				"equal",
			)
		}

		return nil
	}
}
