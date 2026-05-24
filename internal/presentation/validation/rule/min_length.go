package rule

import (
	"github.com/arthurhzna/go-clean-architecture/internal/presentation/validation/core"
)

func MinLength(
	field string,
	value string,
	min int,
) core.Rule {

	return func() core.FieldError {

		if len(value) < min {

			return core.NewValidationError(
				field,
				"min_length",
			)
		}

		return nil
	}
}
