package pinger

import "fmt"

var _ error = PingError{}
var _ error = PingInitError{}

func MakePingError(err error) error {
	return &PingError{msg: err.Error()}
}

func MakePingInitError(err error) error {
	return &PingInitError{msg: err.Error()}
}

type PingError struct {
	msg string
}

func (err PingError) Error() string {
	return fmt.Sprintf("ping error: %s", err.msg)
}

type PingInitError struct {
	msg string
}

func (err PingInitError) Error() string {
	return fmt.Sprintf("ping init error: %s", err.msg)
}
