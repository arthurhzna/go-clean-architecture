package rule

import (
	"strings"

	"github.com/arthurhzna/go-clean-architecture/internal/presentation/validation"
)

func RequiredString(
	field string,
	value string,
) Rule {

	return func() validation.FieldError {

		if strings.TrimSpace(value) == "" {

			return validation.NewValidationError(
				field,
				TagRequired,
			)
		}

		return nil
	}
}

func RequiredInt64(
	field string,
	value int64,
) Rule {
	return func() validation.FieldError {

		if value == 0 {

			return validation.NewValidationError(
				field,
				TagRequired,
			)
		}

		return nil
	}
}
