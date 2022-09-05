package config

import "fmt"

var _ error = configReadError{}
var _ error = configParseError{}

type configReadError struct {
	err error
}

type configParseError struct {
	err error
}

func makeConfigReadError(err error) error {
	return &configReadError{err: err}
}

func makeConfigParseError(err error) error {
	return &configParseError{err: err}
}

func (c configReadError) Error() string {
	return fmt.Sprint("config read error:", c.err)
}

func (c configParseError) Error() string {
	return fmt.Sprint("config parse error:", c.err)
}
