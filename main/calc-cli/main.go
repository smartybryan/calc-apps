package main

import (
	"os"

	"github.com/smartybryan/calc-apps/handlers"
	calc "github.com/smartybryan/calc-lib"
)

func main() {
	handler := handlers.NewCLIHandler(os.Stdout, &calc.Addition{})
	err := handler.Handle(os.Args[1:])
	if err != nil {
		panic(err)
	}
}
