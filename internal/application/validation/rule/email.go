package rule

import (
	"fmt"
	"net/mail"
)

func Email(field string, value string) Rule {
	return func() error {
		_, err := mail.ParseAddress(value)
		if err != nil {
			return fmt.Errorf("%s is invalid", field)
		}

		return nil
	}
}
