package rule

import (
	"fmt"
	"strings"
)

func RequiredString(field string, value string) Rule {
	return func() error {
		if strings.TrimSpace(value) == "" {
			return fmt.Errorf("%s is required", field)
		}

		return nil
	}
}

func RequiredInt64(field string, value int64) Rule {
	return func() error {
		if value == 0 {
			return fmt.Errorf("%s is required", field)
		}

		return nil
	}
}
