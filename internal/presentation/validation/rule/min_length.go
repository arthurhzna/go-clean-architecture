package rule

import (
	"strconv"

	"github.com/arthurhzna/go-clean-architecture/internal/presentation/validation/constant"
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
				constant.TagMinLength,
				strconv.Itoa(min),
			)
		}

		return nil
	}
}
