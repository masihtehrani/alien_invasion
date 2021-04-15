package errors

import (
	"fmt"
)

// define default error system.
var (
	ErrAlienCountEmpty    = err(1001, "alien count must be more than 0")
	ErrRoundCountEmpty    = err(1002, "round count must be more than 0")
	ErrWorldFileEmpty     = err(1003, "world file must be filled")
	ErrCantFindRandomCity = err(1004, "can't find random city")
)

// Error custom error structure.
type Error struct {
	Err  string
	Code int
}

func (e Error) Error() string {
	return fmt.Sprintf("%v", e.Err)
}

func err(code int, s string) Error {
	return Error{
		Err:  s,
		Code: code,
	}
}
