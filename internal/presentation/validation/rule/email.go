package rule

import (
	"regexp"
	"strings"

	"github.com/arthurhzna/go-clean-architecture/internal/presentation/validation/core"
)

var emailRegex = regexp.MustCompile(
	`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`,
)

func Email(
	field string,
	value string,
) core.Rule {

	return func() core.FieldError {

		if strings.TrimSpace(value) == "" {
			return nil
		}

		if !emailRegex.MatchString(value) {

			return core.NewValidationError(
				field,
				"email",
			)
		}

		return nil
	}
}
