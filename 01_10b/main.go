package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
)

var (
	errEmptyOperation            = errors.New("please insert an operation using -expr")
	errInvalidOperation          = errors.New("this operation is not available")
	errCannotParseFloat          = errors.New("cannot parse value")
	errInvalidAmountOfParameters = errors.New("invalid amount of parameters, should be 3")
)

// operators is the map of legal operators and their functions
var operators = map[string]func(x, y float64) float64{
	"+": func(x, y float64) float64 { return x + y },
	"-": func(x, y float64) float64 { return x - y },
	"*": func(x, y float64) float64 { return x * y },
	"/": func(x, y float64) float64 { return x / y },
}

// parseOperand parses a string to a float64
func parseOperand(op string) (float64, error) {
	parsedOp, err := strconv.ParseFloat(op, 64)

	if err != nil {
		return 0, errCannotParseFloat
	}

	return parsedOp, nil
}

// calculate returns the result of a 2 operand mathematical expression
func calculate(expr string) (float64, error) {
	ops := strings.Fields(expr)

	if len(ops) != 3 {
		return 0, errInvalidAmountOfParameters
	}

	left, lErr := parseOperand(ops[0])
	right, rErr := parseOperand(ops[2])

	if lErr != nil || rErr != nil {
		return 0, errCannotParseFloat
	}

	f := operators[ops[1]]

	if f == nil {
		return 0, errInvalidOperation
	}

	result := f(left, right)
	return result, nil
}

func main() {
	expr := flag.String("expr", "",
		"The expression to calculate on, separated by spaces.")

	flag.Parse()

	if expr == nil || *expr == "" {
		printError(errEmptyOperation)
		return
	}

	result, err := calculate(*expr)

	if err != nil {
		printError(err, *expr)
		return
	}

	log.Printf("%s = %.2f\n", *expr, result)
}

func printError(errorMessage error, extraData ...interface{}) {
	fmt.Printf("Error: %s", errorMessage)

	if len(extraData) > 0 {
		for _, value := range extraData {
			fmt.Printf(" %v", value)
		}
	}

	fmt.Printf("\n")
}
