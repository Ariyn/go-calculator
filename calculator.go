package calculator

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

var (
	TypeOperand  = "operand"
	TypeOperator = "operator"
)

type Element interface {
	Type() string
}

type Operator interface {
	Type() string
	Priority() int
	SizeOfOperands() int
	Do(values ...Operand) Operand
	String() string
}

type TwoOperandOperator struct {
	signal   string
	priority int
	do       func(a, b float64) float64
}

func NewTwoOperandOperator(signal string, priority int, do func(a, b float64) float64) TwoOperandOperator {
	return TwoOperandOperator{
		signal:   signal,
		priority: priority,
		do:       do,
	}
}

func (t TwoOperandOperator) Type() string {
	return TypeOperator
}

func (t TwoOperandOperator) SizeOfOperands() int {
	return 2
}

func (t TwoOperandOperator) String() string {
	return t.signal
}

func (t TwoOperandOperator) Priority() int {
	return t.priority
}

func (t TwoOperandOperator) Do(values ...Operand) Operand {
	if len(values) != 2 {
		panic(errors.New("incorrect operand size"))
	}

	result := t.do(values[0].Float(), values[1].Float())
	return Operand{
		Typ:   reflect.Float64,
		value: result,
	}
}

func NewOperatorByString(value rune) (o Operator, err error) {
	switch value {
	case '+':
		o = NewTwoOperandOperator("+", 1, func(a, b float64) float64 {
			return a + b
		})
	case '-':
		o = NewTwoOperandOperator("-", 1, func(a, b float64) float64 {
			return a - b
		})
	case '*':
		o = NewTwoOperandOperator("x", 2, func(a, b float64) float64 {
			return a * b
		})
	case '/':
		o = NewTwoOperandOperator("/", 2, func(a, b float64) float64 {
			return a / b
		})
	}

	return
}

type Operand struct {
	Typ   reflect.Kind
	value interface{}
}

func (o Operand) Type() string {
	return TypeOperand
}

func (o Operand) String() string {
	if int(o.Float() * 100)%100 == 0 {
		return fmt.Sprintf("%d", o.Int())
	}
	return fmt.Sprintf("%0.1f", o.Float())
}

func (o Operand) Float() float64 {
	var v interface{}
	switch o.Typ {
	case reflect.Int:
		v = float64(o.value.(int))
	case reflect.Int64:
		v = float64(o.value.(int64))
	case reflect.Float64:
		v = o.value.(float64)
	}
	return v.(float64)
}

func (o Operand) Int() int {
	return int(o.Float())
}

func (o Operand) Int64() int64 {
	return int64(o.Float())
}

func NewOperandByString(value string) (o Operand, err error) {
	v, err := strconv.Atoi(value)
	if err != nil {
		return
	}

	return Operand{
		Typ:   reflect.Int,
		value: v,
	}, nil
}

// 0 = 48, 1 = 49, 9 = 57, . = 46
func IsOperand(value rune) bool {
	return ('0' <= value && value <= '9') || value == '.'
}

// / = 47, + = 43, - = 45, * = 42, ~ = 126, ^ = 94, & = 38
func IsOperator(value rune) bool {
	return value == '/' || value == '+' || value == '-' || value == '*' || value == '~' || value == '^' || value == '&'
}
