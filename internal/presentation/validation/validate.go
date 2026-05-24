package validation

import "github.com/arthurhzna/go-clean-architecture/internal/presentation/validation/core"

func Validate(
	rules []core.Rule,
) error {

	return core.Execute(rules)
}
