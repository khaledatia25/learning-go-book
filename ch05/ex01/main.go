package main

import (
	"errors"
	"fmt"
	"strconv"
)

type SimpleCalculatorFunction = func(int, int) (int, error)

var opMap map[string]SimpleCalculatorFunction = map[string]SimpleCalculatorFunction{
	"+": func(i, j int) (int, error) {
		return i + j, nil
	},
	"-": func(i, j int) (int, error) {
		return i - j, nil
	},
	"*": func(i, j int) (int, error) {
		return i * j, nil
	},
	"/": func(i, j int) (int, error) {
		if j == 0 {
			return 0, errors.New("invalid division by zero")
		}
		return i / j, nil
	},
}

func main() {
	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "0"},
		{"2", "%", "1"},
		{"two", "+", "three"},
		{"5"},
	}
	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("invalid expression:", expression)
			continue
		}
		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		op := expression[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println("unsupported operator:", op)
			continue
		}
		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}
		result, err := opFunc(p1, p2)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(result)
		}
	}
}
