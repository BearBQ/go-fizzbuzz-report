package fizzbuzz

import (
	"reflect"
	"testing"
)

func TestCalcFizzBuzz(t *testing.T) {
	example := FizzBuzz{FizzBuzzResultSlice: []Result{
		{Number: 1, Output: "1", Type: "number"},
		{Number: 2, Output: "2", Type: "number"},
		{Number: 3, Output: "Fizz", Type: "fizz"},
		{Number: 4, Output: "4", Type: "number"},
		{Number: 5, Output: "Buzz", Type: "buzz"},
	}}

	testing := FizzBuzz{FizzBuzzResultSlice: make([]Result, 0, 5)}
	CalcFizzBuzz(&testing, 5)
	if !reflect.DeepEqual(example, testing) {
		t.Errorf("допущена ошибка в создании цепочки")
	}
}
