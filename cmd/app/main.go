package main

import (
	"github.com/ariyn/calculator"
	"log"
	"sort"
)

func main() {
	elements := calculator.ParseString("1 + 2 * 3")
	for i := 0; i < len(elements); i++ {
		e := elements[i]
		if e.Type() == calculator.TypeOperator {
			// TODO: error handling
			operator := e.(calculator.Operator)

			// TODO: error handling
			result := operator.Do(elements[i-2].(calculator.Operand), elements[i-1].(calculator.Operand))
			elements = remove(elements, i-2, i-1, i)
			elements = push(elements, result, i-2)
			i -= 3
		}
	}

	// TODO: error handling
	log.Println(elements[0])
}

func remove(elements []calculator.Element, indexes ...int) []calculator.Element {
	if len(elements) < len(indexes) {
		return elements
	}

	sort.Slice(indexes, func(i, j int) bool {
		return indexes[i] > indexes[j]
	})
	for _, i := range indexes {
		elements = append(elements[:i], elements[i+1:]...)
	}

	return elements
}

func push(elements []calculator.Element, operand calculator.Operand, index int) []calculator.Element {
	elements = append(elements, operand)
	copy(elements[index+1:], elements[index:])
	elements[index] = operand
	return elements
}
