package main

import (
	"bytes"
	"go-fizzbuzz-report/fizzbuzz"
	"go-fizzbuzz-report/output"
	"reflect"
	"testing"
)

func TestGenerateFizzBuzz(t *testing.T) {
	example := fizzbuzz.FizzBuzz{FizzBuzzResultSlice: []fizzbuzz.FizzBuzzItem{
		{Number: 1, Output: "1", Type: "number"},
		{Number: 2, Output: "2", Type: "number"},
		{Number: 3, Output: "Fizz", Type: "fizz"},
		{Number: 4, Output: "4", Type: "number"},
		{Number: 5, Output: "Buzz", Type: "buzz"},
	}}

	testing := fizzbuzz.FizzBuzz{FizzBuzzResultSlice: make([]fizzbuzz.FizzBuzzItem, 0, 5)}
	newTesting, _ := fizzbuzz.GenerateFizzBuzz(&testing, 5)
	if !reflect.DeepEqual(example, newTesting) {
		t.Errorf("допущена ошибка в создании цепочки\nПолучившийся массив: %v", testing)
	}
}

// с*№ с интернета тест. не совсем дошел, но пусть так
func TestPrintFizzBuzz(t *testing.T) {
	oldStdout := output.Stdout
	defer func() { output.Stdout = oldStdout }()
	testBuf := new(bytes.Buffer) //я так понял подменяем консоль нашим буффером
	output.Stdout = testBuf
	testData := [][3]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
	}
	err := output.PrintFizzBuzz(testData)
	if err != nil {
		t.Fatalf("Ошибка не ожидалась, но получена: %v", err)
	}
	expected := "1      2      3\n4      5      6\n"

	if testBuf.String() != expected {
		t.Errorf("\nОжидалось: %q\nПолучено:  %q", expected, testBuf.String())
	}
}

func TestPrintFizzBuzzNumb(t *testing.T) {
	oldStdout := output.Stdout
	defer func() { output.Stdout = oldStdout }()
	testBuf := new(bytes.Buffer)
	output.Stdout = testBuf
	testData := [4]int{1, 1, 1, 1}
	err := output.PrintFizzBuzzNumb(testData)
	if err != nil {
		t.Fatalf("Ошибка не ожидалась, но получена: %v", err)
	}
	expected := "Количество FizzBuzz - 1, Количество Fizz - 1, Количество Buzz - 1, Количество числовых значений - 1 \n"
	if testBuf.String() != expected {
		t.Errorf("\nОжидалось: %q\nПолучено:  %q", expected, testBuf.String())
	}

}
