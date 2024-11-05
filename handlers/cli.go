package handlers

import (
	"errors"
	"fmt"
	"io"
	"strconv"
)

type Calculator interface {
	Calculate(a, b int) int
}

type CLIHandler struct {
	stdout     io.Writer
	calculator Calculator
}

func NewCLIHandler(stdout io.Writer, calculator Calculator) *CLIHandler {
	return &CLIHandler{stdout, calculator}
}

func (this *CLIHandler) Handle(args []string) error {
	if this.calculator == nil {
		return errUnsupportedOperation
	}

	if len(args) != 2 {
		return errWrongArgCount
	}

	a, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("%w: '%s'", errInvalidArgument, args[0])
	}
	b, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("%w: '%s'", errInvalidArgument, args[1])
	}
	result := this.calculator.Calculate(a, b)
	_, err = fmt.Fprint(this.stdout, result)
	if err != nil {
		return fmt.Errorf("%w: %w", errOutputFailure, err)
	}
	return nil
}

var (
	errWrongArgCount        = errors.New("usage: calculator <a> <b>")
	errInvalidArgument      = errors.New("invalid syntax")
	errOutputFailure        = errors.New("output failure")
	errUnsupportedOperation = errors.New("unsupported operation")
)
