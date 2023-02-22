package main

import (
	"flag"
	"log"
)

// isBalanced returns whether the given expression
// has balanced brackets.
func isBalanced(expr string) bool {
	var closingComparator string

	for _, r := range expr {
		if o, c := isOpeningSymbol(r); o {
			closingComparator = string(c) + closingComparator
			continue
		}

		if r != rune(closingComparator[0]) {
			return false
		}

		if len(closingComparator) > 0 {
			closingComparator = closingComparator[1:]
		}

	}

	if len(closingComparator) > 0 {
		return false
	}

	return true
}

func isOpeningSymbol(symbol rune) (bool, rune) {
	switch symbol {
	case '(':
		return true, ')'
	case '[':
		return true, ']'
	case '{':
		return true, '}'
	default:
		return false, ' '
	}
}

// printResult prints whether the expression is balanced.
func printResult(expr string, balanced bool) {
	if balanced {
		log.Printf("%s is balanced.\n", expr)
		return
	}
	log.Printf("%s is not balanced.\n", expr)
}

func main() {
	expr := flag.String("expr", "", "The expression to validate brackets on.")
	flag.Parse()
	printResult(*expr, isBalanced(*expr))
}
