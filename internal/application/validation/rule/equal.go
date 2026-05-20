package rule

import "fmt"

func Equal(field string, value string, expected string) Rule {
	return func() error {
		if value != expected {
			return fmt.Errorf("%s is not equal", field)
		}

		return nil
	}
}
