package rule

import "fmt"

func MinLength(field string, value string, min int) Rule {
	return func() error {
		if len(value) < min {
			return fmt.Errorf(
				"%s minimum length is %d",
				field,
				min,
			)
		}

		return nil
	}
}
