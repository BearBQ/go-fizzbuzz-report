package main

import (
	"fmt"
	"strconv"
)

type Result struct {
	Number int    // само число, например 3
	Output string // что вывести: "Fizz", "Buzz", "FizzBuzz" или "3"
	Type   string // метка типа: "fizz", "buzz", "fizzbuzz", "number"
}

type FizzBuzz struct { //объявлю новую структуру, которая будет представлять из себя массив структур
	ResultSlice []Result
}

func NewFizzBuzz(numb int) *FizzBuzz { //создание новой физбаз предоставим функции, с заданной cap у слайса
	return &FizzBuzz{
		ResultSlice: make([]Result, 0, numb),
	}
}

func CalcFizzBuzz(f *FizzBuzz, numb int) *FizzBuzz { //функция формирования последовательности
	for i := 1; i < numb+1; i++ {
		var result Result
		switch {
		case i%15 == 0:
			result = Result{Number: i, Output: "FizzBuzz", Type: "fizzbuzz"}
		case i%3 == 0:
			result = Result{Number: i, Output: "Fizz", Type: "fizz"}
		case i%5 == 0:
			result = Result{Number: i, Output: "Buzz", Type: "buzz"}
		default:
			result = Result{Number: i, Output: strconv.Itoa(i), Type: "number"}

		}
		f.ResultSlice = append(f.ResultSlice, result)
	}
	return f
}

func PrintFizzBuzz(f *FizzBuzz) { //функция печати
	if len(f.ResultSlice) != 0 {
		for _, value := range f.ResultSlice {
			fmt.Printf("number:%v, output:%s, type: %s\n", value.Number, value.Output, value.Type)
		}
	}
}

func PrintFizzBuzzNumb(f *FizzBuzz) { //функция печати
	if len(f.ResultSlice) != 0 {
		numbOfNumb, numbOfFizz, numbOfBuzz, numbOfFizzBuzz := 0, 0, 0, 0
		for _, value := range f.ResultSlice {
			switch {
			case value.Type == "fizzbuzz":
				numbOfFizzBuzz++
			case value.Type == "fizz":
				numbOfFizz++
			case value.Type == "buzz":
				numbOfBuzz++
			case value.Type == "number":
				numbOfNumb++
			}
		}
		fmt.Printf("Количество fizzbuzz: %v, Количество fizz: %v, Количество buzz: %v, Количество чисел: %v, ", numbOfFizzBuzz, numbOfFizz, numbOfBuzz, numbOfNumb)
	}
}
