package main

import (
	"fmt"
	"go-fizzbuzz-report/fizzbuzz"
	"go-fizzbuzz-report/input"
)

func main() {
	number := input.InputNumber()
	newFizz := fizzbuzz.NewFizzBuzz(number)
	resultFizzBuzz, err := fizzbuzz.CalcFizzBuzz(newFizz, number)
	if err != nil {
		fmt.Println("Ошибка создания FizzBuzz:", err)
		return
	}

	fizzbuzz.PrintFizzBuzz(&resultFizzBuzz)
	fizzbuzz.PrintFizzBuzzNumb((&resultFizzBuzz))
}
