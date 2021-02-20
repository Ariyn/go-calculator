package calculator

import "log"

/*
	1 + 2 * 3
	1      +2*3
	1   +   2*3
	12  +  *3
	12  +*   3
	123 +*
	123* +
	123*+

	2*(3+1)+1
	2      *(3+1)+1
	2   *     (3+1)+1
	2   *(    3+1)+1
	23  *(    +1)+1
	23   *(+   1)+1
	231   *(+  )+1
	231+  *    +1
	231+*  +   1
	231+*1 +
	231+*1+
*/

// TODO: error handling
func ParseString(v string) (elements []Element) {
	elements = make([]Element, 0)
	operators := make([]Operator, 0)

	previous := ""
	for i := 0; i < len(v); i++ {
		r := v[i]
		if IsOperand(rune(r)) {
			previous += string(r)
		} else if IsOperator(rune(r)) {
			appendOperandToElements(&elements, previous)
			previous = ""
			appendToOperators(&operators, string(r))

			if compareOperators(operators) {
				o := popFromOperators(&operators, len(operators)-2)
				appendOperatorToElements(&elements, o)
			}
		}
	}

	if previous != "" {
		appendOperandToElements(&elements, previous)
	}

	for i := len(operators) - 1; i >= 0; i-- {
		appendOperatorToElements(&elements, operators[i])
	}

	return
}

func appendOperandToElements(elements *[]Element, value string) {
	o, err := NewOperandByString(value)
	if err != nil {
		log.Fatal(err)
	}

	*elements = append(*elements, o)
}

func appendOperatorToElements(elements *[]Element, operator Operator) {
	*elements = append(*elements, operator)
}

func appendToOperators(operators *[]Operator, value string) {
	o, err := NewOperatorByString(rune(value[0]))
	if err != nil {
		log.Fatal(err)
	}

	*operators = append(*operators, o)
}

func popFromOperators(operators *[]Operator, index int) Operator {
	o := (*operators)[index]
	*operators = append((*operators)[:index], (*operators)[index+1:]...)
	return o
}

func compareOperators(operators []Operator) bool {
	length := len(operators)
	if length < 2 {
		return false
	}

	return operators[length-2].Priority() > operators[length-1].Priority()
}