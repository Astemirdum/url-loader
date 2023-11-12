package pkg

import (
	"errors"
	"fmt"
	"net/url"
)

var ErrInvalidUrl = errors.New("invalid urls")

func ValidateURL(u string) error {
	_, err := url.ParseRequestURI(u)
	if err != nil {
		return fmt.Errorf("%s: %w", err.Error(), ErrInvalidUrl)
	}
	return nil
}
