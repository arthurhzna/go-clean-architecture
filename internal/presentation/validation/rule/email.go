package rule

import (
	"regexp"
	"strings"

	"github.com/arthurhzna/go-clean-architecture/internal/presentation/validation"
)

var emailRegex = regexp.MustCompile(
	`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`,
)

func Email(
	field string,
	value string,
) Rule {

	return func() validation.FieldError {

		if strings.TrimSpace(value) == "" {
			return nil
		}

		if !emailRegex.MatchString(value) {

			return validation.NewValidationError(
				field,
				"email",
			)
		}

		return nil
	}
}
