package validation

import "github.com/arthurhzna/go-clean-architecture/internal/presentation/validation/core"

type Errors = core.Errors
type FieldError = core.FieldError
type Rule = core.Rule

func Validate(
	rules []Rule,
) error {

	return core.Execute(rules)
}
