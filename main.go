package main

import "go-fizzbuzz-report/fizzbuzz"

func main() {
	var numb int = 44
	newFizz := fizzbuzz.NewFizzBuzz(numb)
	fizzbuzz.CalcFizzBuzz(newFizz, numb)
	fizzbuzz.PrintFizzBuzz(newFizz)
	fizzbuzz.PrintFizzBuzzNumb((newFizz))
}
