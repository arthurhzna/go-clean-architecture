package rule

type Rule func() error

func Execute(
	rules []Rule,
) error {

	for _, rule := range rules {

		if err := rule(); err != nil {
			return err
		}
	}

	return nil
}
