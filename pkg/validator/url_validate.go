package validator

import "net/url"

// Validates URL string. Takes URL string as parameter. Returns nil if URL is valid
// else returns an error.
func Validate(URL string) error {

	_, err := url.ParseRequestURI(URL)

	if err != nil {
		return err
	}

	return nil
}
