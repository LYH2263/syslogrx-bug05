package syslogrx

import "fmt"

// Malformed wraps err as a malformed-syslog error that callers can identify
// via the ErrMalformed sentinel using errors.Is. A nil err yields a sentinel
// error describing an empty cause.
func Malformed(err error) error {
	if err == nil {
		return fmt.Errorf("%w: empty", ErrMalformed)
	}
	// %w wraps both ErrMalformed (so errors.Is(.., ErrMalformed) holds) and the
	// underlying cause (so its own sentinels/types remain extractable).
	return fmt.Errorf("%w: %w", ErrMalformed, err)
}
