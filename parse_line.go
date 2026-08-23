package syslogrx

import "fmt"

// ParseLine parses a single syslog line as RFC 5424, falling back to RFC 3164.
// Any parse failure — including an empty line — is wrapped with Malformed so
// callers can detect it via errors.Is(err, ErrMalformed).
func ParseLine(line string) (*Message, error) {
	if line == "" {
		return nil, Malformed(fmt.Errorf("empty line"))
	}
	if m, err := ParseRFC5424(line); err == nil {
		return m, nil
	}
	m, err := ParseRFC3164(line)
	if err != nil {
		return nil, Malformed(err)
	}
	return m, nil
}
