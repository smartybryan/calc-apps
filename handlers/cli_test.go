package handlers

import (
	"bytes"
	"errors"
	"testing"

	calc "github.com/smartybryan/calc-lib"
)

func TestHandler_NilCalculator(t *testing.T) {
	handler := NewCLIHandler(nil, nil)
	err := handler.Handle(nil)
	assertError(t, err, errUnsupportedOperation)
}

func TestHandler_TwoArgsRequired(t *testing.T) {
	handler := NewCLIHandler(nil, &calc.Addition{})
	err := handler.Handle(nil)
	assertError(t, err, errWrongArgCount)
}

func TestHandler_FirstArgInvalid(t *testing.T) {
	handler := NewCLIHandler(nil, &calc.Addition{})
	err := handler.Handle([]string{"INVALID", "42"})
	assertError(t, err, errInvalidArgument)
}

func TestHandler_SecondArgInvalid(t *testing.T) {
	handler := NewCLIHandler(nil, &calc.Addition{})
	err := handler.Handle([]string{"42", "INVALID"})
	assertError(t, err, errInvalidArgument)
}

func TestHandler_ResultWrittenToOutput(t *testing.T) {
	outbuf := &bytes.Buffer{}
	handler := NewCLIHandler(outbuf, &calc.Addition{})
	err := handler.Handle([]string{"2", "3"})
	assertError(t, err, nil)
	if outbuf.String() != "5" {
		t.Errorf("want: %v, got: %v", "5", outbuf.String())
	}
}

func TestHandler_OutputWriterError(t *testing.T) {
	boink := errors.New("boink")
	stdout := &ErringWriter{boink}
	handler := NewCLIHandler(stdout, &calc.Addition{})
	err := handler.Handle([]string{"2", "3"})
	assertError(t, err, errOutputFailure, boink)
}

/////////////////////////////////////////////////////////

func assertError(t *testing.T, actual error, targets ...error) {
	for _, target := range targets {
		if !errors.Is(actual, target) {
			t.Helper()
			t.Errorf("want: %v, actual: %v", target, actual)
		}
	}
}

type ErringWriter struct {
	err error
}

func (this *ErringWriter) Write([]byte) (n int, err error) {
	return 0, this.err
}
