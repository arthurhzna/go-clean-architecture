package rule

import (
	"time"

	"github.com/arthurhzna/go-clean-architecture/internal/presentation/validation"
)

func TimeFormat(
	field string,
	value string,
	layout string,
) Rule {

	return func() validation.FieldError {

		if value == "" {
			return nil
		}

		_, err := time.Parse(layout, value)

		if err != nil {

			return validation.NewValidationError(
				field,
				TagTimeFormat,
			)
		}

		return nil
	}
}
