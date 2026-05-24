package rule

import (
	"strings"

	"github.com/arthurhzna/go-clean-architecture/internal/presentation/validation/core"
	"github.com/arthurhzna/go-clean-architecture/internal/presentation/validation/utils"
)

func RequiredString(
	field string,
	value string,
) core.Rule {

	return func() core.FieldError {

		if strings.TrimSpace(value) == "" {

			return core.NewValidationError(
				field,
				utils.TagRequired,
			)
		}

		return nil
	}
}

func RequiredInt64(
	field string,
	value int64,
) core.Rule {
	return func() core.FieldError {

		if value == 0 {

			return core.NewValidationError(
				field,
				utils.TagRequired,
			)
		}

		return nil
	}
}
