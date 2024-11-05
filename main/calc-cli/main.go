package main

import (
	"flag"
	"os"

	"github.com/smartybryan/calc-apps/handlers"
	calc "github.com/smartybryan/calc-lib"
)

func main() {
	var op string
	flag.StringVar(&op, "op", "+", "Enter operation: +,-,*,/")
	flag.Parse()

	handler := handlers.NewCLIHandler(os.Stdout, calculators[op])
	err := handler.Handle(flag.Args())
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
