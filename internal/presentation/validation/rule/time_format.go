package rule

import (
	"time"

	"github.com/arthurhzna/go-clean-architecture/internal/presentation/validation/core"
	"github.com/arthurhzna/go-clean-architecture/internal/presentation/validation/utils"
)

func TimeFormat(
	field string,
	value string,
	layout string,
) core.Rule {

	return func() core.FieldError {

		if value == "" {
			return nil
		}

		_, err := time.Parse(layout, value)

		if err != nil {

			return core.NewValidationError(
				field,
				utils.TagTimeFormat,
			)
		}

		return nil
	}
}
