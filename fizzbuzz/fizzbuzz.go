package fizzbuzz

import (
	"fmt"
	"go-fizzbuzz-report/input"
	"go-fizzbuzz-report/output"
	"strconv"
)

// Result Структура
type FizzBuzzItem struct {
	Number int    // само число
	Output string // что вывести: "Fizz", "Buzz", "FizzBuzz" или чисор
	Type   string // метка типа: "fizz", "buzz", "fizzbuzz", "number"
}

// FizzBuzz массив структур
type FizzBuzz struct {
	FizzBuzzResultSlice []FizzBuzzItem
}

// NewFizzBuzz создание новой физбаз заданной емкости
func NewFizzBuzz(numb int) (FizzBuzz, error) {
	if numb <= 0 {
		return FizzBuzz{}, fmt.Errorf("число не может быть отрицательным либо равным нулю")
	}
	return FizzBuzz{
		FizzBuzzResultSlice: make([]FizzBuzzItem, 0, numb),
	}, nil
}

// GenerateFizzBuzz возвращает новую структуру FizzBuzz и ошибку
func GenerateFizzBuzz(f *FizzBuzz, number int) (FizzBuzz, error) { //функция формирования последовательности

	fResult := FizzBuzz{
		FizzBuzzResultSlice: make([]FizzBuzzItem, 0, number),
	}
	for i := 1; i < number+1; i++ {
		var result FizzBuzzItem
		switch {
		case i%15 == 0:
			result = FizzBuzzItem{Number: i, Output: "FizzBuzz", Type: "fizzbuzz"}
		case i%3 == 0:
			result = FizzBuzzItem{Number: i, Output: "Fizz", Type: "fizz"}
		case i%5 == 0:
			result = FizzBuzzItem{Number: i, Output: "Buzz", Type: "buzz"}
		default:
			result = FizzBuzzItem{Number: i, Output: strconv.Itoa(i), Type: "number"}
		}
		fResult.FizzBuzzResultSlice = append(fResult.FizzBuzzResultSlice, result)
	}
	return fResult, nil
}

// FizzBuzz подсчитывает количество элементов каждого типа
func countFizzBuzz(f FizzBuzz) ([4]int, error) {
	var result [4]int

	if len(f.FizzBuzzResultSlice) == 0 {
		return result, fmt.Errorf("на вход поступил пустой слайс")
	}
	for _, value := range f.FizzBuzzResultSlice {
		switch {
		case value.Type == "fizzbuzz":
			result[0]++
		case value.Type == "fizz":
			result[1]++
		case value.Type == "buzz":
			result[2]++
		case value.Type == "number":
			result[3]++
		}
	}
	return result, nil
}

// GetPrintNames
func GetPrintNames(f FizzBuzz) ([][3]string, error) {
	lenSlice := len(f.FizzBuzzResultSlice)
	var valueEachElement [3]string
	valueFizzBuz := make([][3]string, 0, lenSlice)
	if lenSlice == 0 {
		return valueFizzBuz, fmt.Errorf("данные отсутствуют")
	}
	for _, value := range f.FizzBuzzResultSlice {
		valueEachElement[0] = strconv.Itoa(value.Number)
		valueEachElement[1] = value.Output
		valueEachElement[2] = value.Type
		valueFizzBuz = append(valueFizzBuz, valueEachElement)
	}
	return valueFizzBuz, nil
}

// RunfizzBuzz - основная логика программы
func RunFizzBuzz() error {
	number, err := input.InputNumber()
	if err != nil {
		return err
	}

	//Создадим пустой срез заданной capacity
	fizBuzzEmpty, err := NewFizzBuzz(number)
	if err != nil {
		return err
	}
	fizBuzzFinal, err := GenerateFizzBuzz(&fizBuzzEmpty, number) //создаем последовательность
	if err != nil {
		return err
	}

	var typeCounts [4]int //считаем типы
	typeCounts, err = countFizzBuzz(fizBuzzFinal)
	if err != nil {
		return err
	}

	var printValue [][3]string
	printValue, err = GetPrintNames(fizBuzzFinal)
	if err != nil {
		return err
	}

	err = output.PrintFizzBuzz(printValue)
	if err != nil {
		return err
	}

	err = output.PrintFizzBuzzNumb(typeCounts)
	if err != nil {
		return err
	}
	return nil
}
