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
	if len(args) != 2 {
		return ErrWrongArgCount
	}

	a, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("%w: '%s'", ErrInvalidArgument, args[0])
	}
	b, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("%w: '%s'", ErrInvalidArgument, args[1])
	}
	result := this.calculator.Calculate(a, b)
	_, err = fmt.Fprint(this.stdout, result)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrOutputFailure, err)
	}
	return nil
}

var (
	ErrWrongArgCount   = errors.New("usage: calculator <a> <b>")
	ErrInvalidArgument = errors.New("invalid syntax")
	ErrOutputFailure   = errors.New("output failure")
)
