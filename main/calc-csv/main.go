package main

import (
	"log"
	"os"

	"github.com/smartybryan/calc-apps/handlers"
	calc "github.com/smartybryan/calc-lib"
)

func main() {
	logger := log.New(os.Stderr, ">>>", 0)
	handler := handlers.NewCSVHandler(logger, os.Stdin, os.Stdout, calculators)
	err := handler.Handle()
	if err != nil {
		panic(err)
	}
}

var calculators = map[string]handlers.Calculator{
	"+": &calc.Addition{},
	"-": &calc.Subtraction{},
	"*": &calc.Multiplication{},
	"/": &calc.Division{},
}
