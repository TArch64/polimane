package base

import "fmt"

func TagError(tag string, err error) error {
	if err == nil {
		return nil
	}

	return fmt.Errorf("%s: %w", tag, err)
}
