package rule

import (
	"github.com/arthurhzna/go-clean-architecture/internal/presentation/validation/constant"
	"github.com/arthurhzna/go-clean-architecture/internal/presentation/validation/core"
)

func Equal(
	field string,
	value string,
	compare string,
) core.Rule {

	return func() core.FieldError {

		if value != compare {

			return core.NewValidationError(
				field,
				constant.TagEqual,
				"",
			)
		}

		return nil
	}
}
