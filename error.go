package base122

import "errors"

var (
	ErrInputInvalid = errors.New("input is not a valid base122 encoded string")
	ErrEndOfBytes   = errors.New("end of bytes")
)
