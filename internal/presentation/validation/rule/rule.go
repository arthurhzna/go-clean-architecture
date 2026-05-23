package rule

import "github.com/arthurhzna/go-clean-architecture/internal/presentation/validation"

type Rule func() validation.FieldError

func Execute(
	rules []Rule,
) error {

	var errs validation.Errors

	for _, rule := range rules {

		if err := rule(); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}
