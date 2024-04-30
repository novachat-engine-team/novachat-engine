package errors

import "errors"

// New returns an error that formats as the given text.
// Each call to New returns a distinct error value even if the text is identical.
// because the former will succeed if err wraps an *fs.PathError.
func New(text string) error {
	return errors.New(text)
}
