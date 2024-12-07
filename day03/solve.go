package day03

import (
	"fmt"
)

var debug bool

type state struct {
	previous           rune
	digitCount         int
	firstValue         int
	secondValue        int
	containsComma      bool
	valid              bool
	disabled           bool
	isConditionalState bool
	shouldDisable      bool
	content            []rune
}

func (s state) print() {
	if s.valid && s.previous == ')' {
		fmt.Printf(
			"State{\n\tprevious: '%c',\n\tdigitCount: %d\n\tfirstValue: %d\n\tsecondValue: %d\n\tcontainsComma: %t\n\tvalid: %t\n\tdisabled: %t\n\tisConditionalState: %t\n\tshouldDisable: %t\n\tcontent: %s\n}\n",
			s.previous, s.digitCount, s.firstValue, s.secondValue, s.containsComma, s.valid, s.disabled, s.isConditionalState, s.shouldDisable, string(s.content),
		)
	}
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func solve(content string, withConditionalStatements bool) int {
	var result int
	s := state{}
	for _, curr := range content {
		if debug {
			s.print()
			fmt.Printf("Reading rune '%c'\n", curr)
		}
		switch curr {
		case 'd':
			s = state{disabled: s.disabled}
			s.valid = true
			s.isConditionalState = true
		case 'o':
			if !s.valid {
				break
			}
			if s.previous != 'd' {
				s.valid = false
			}
		case 'n':
			if !s.valid {
				break
			}
			if s.previous != 'o' {
				s.valid = false
				break
			}
			s.shouldDisable = true
		case '\'':
			if !s.valid {
				break
			}
			if s.previous != 'n' {
				s.valid = false
			}
		case 't':
			if !s.valid {
				break
			}
			if s.previous != '\'' {
				s.valid = false
			}
		case 'm':
			s = state{disabled: s.disabled}
			s.valid = true
		case 'u':
			if !s.valid {
				break
			}
			if s.previous != 'm' {
				s.valid = false
			}
		case 'l':
			if !s.valid {
				break
			}
			if s.previous != 'u' {
				s.valid = false
			}
		case '(':
			if !s.valid {
				break
			}
			if s.previous != 'l' && s.previous != 't' && s.previous != 'o' {
				s.valid = false
			}
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			if !s.valid {
				break
			}
			if !isDigit(s.previous) && s.previous != '(' && s.previous != ',' && s.digitCount >= 3 {
				s.valid = false
				break
			}
			if !s.containsComma {
				s.firstValue = s.firstValue*10 + int(curr-'0')
			} else {
				s.secondValue = s.secondValue*10 + int(curr-'0')
			}
			s.digitCount++
		case ',':
			if !s.valid {
				break
			}
			if !isDigit(s.previous) {
				s.valid = false
				break
			}
			s.containsComma = true
			s.digitCount = 0
		case ')':
			if !s.valid {
				break
			}
			if isDigit(s.previous) {
				if !withConditionalStatements || !s.disabled {
					result += s.firstValue * s.secondValue
				}
			} else if s.previous == '(' && s.isConditionalState {
				s.disabled = s.shouldDisable
			} else {
				s.valid = false
				break
			}
		default:
			s.valid = false
		}
		s.content = append(s.content, curr)
		s.previous = curr
	}
	if debug {
		s.print()
	}
	return result
}

func Solve(inputFileName string, d bool) error {
	debug = d
	content, err := parse(inputFileName)
	if err != nil {
		return err
	}
	fmt.Println("Part 1:", solve(content, false))
	fmt.Println("Part 2:", solve(content, true))
	return nil
}
