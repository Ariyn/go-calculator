package main

import (
	"bufio"
	"fmt"
	"github.com/ariyn/calculator"
	"os"
	"sort"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	for {
		line, _, err := r.ReadLine()
		if err != nil {
			break
		}

		result := parse(string(line))
		fmt.Println(fmt.Sprintf("%s = %s", string(line), result.String()))
	}
}

func parse(line string) calculator.Operand {
	elements := calculator.ParseString(line)
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

	//TODO: error handling
	return elements[0].(calculator.Operand)
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
