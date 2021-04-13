package errors_test

import (
	"testing"

	"github.com/masihtehrani/alien_invasion/internal/adaptors/errors"
)

func TestError_Error(t *testing.T) {
	err := errors.Error{Err: "sample error message", Code: 10}
	msg := "sample error message"
	// err := errors.Error(msg)

	if msg != err.Error() {
		t.Error("message & error message not equal together.")
	}
}
